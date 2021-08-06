package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Arman92/go-tl-parser/tlparser"
)

func getFilePath(outputDir, returnType string) string {
	return filepath.Join(outputDir, firstLower(replaceKeyWords(returnType))+".go")
}

func GenerateMethods(schema *tlparser.TlSchema, basePackageUri, typePackageName, packageName, outputDir string) error {
	// first remove all files previously created (note: we don't want to remove the whole directory
	// as there is other non-autogenerated work in the methods directory
	for _, function := range schema.Functions {
		filePath := getFilePath(outputDir, function.ReturnType)
		os.Remove(filePath)
	}

	for _, function := range schema.Functions {
		buf := bytes.NewBufferString("\n")

		filePath := getFilePath(outputDir, function.ReturnType)
		// Only add go package and imports if file does not exist already.
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			buf.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))

			buf.WriteString(fmt.Sprintf(`
		
			import (
				"encoding/json"
				"fmt"
				"strconv"
				"strings"
				"%s/%s"
			)
		
			`, basePackageUri, typePackageName))
		}

		methodName := firstUpper(function.Name)
		methodName = replaceKeyWords(methodName)
		returnType := strings.ToUpper(function.ReturnType[:1]) + function.ReturnType[1:]
		returnType = replaceKeyWords(returnType)
		returnTypeCamel := strings.ToLower(returnType[:1]) + returnType[1:]
		returnIsInterface := checkIsInterface(returnType, schema.Interfaces)

		asterike := "*" + typePackageName + "."
		ampersign := "&"
		if returnIsInterface {
			asterike = ""
			ampersign = ""
			returnType = typePackageName + "." + returnType
		}

		paramsStr := ""
		paramsDesc := ""
		for i, param := range function.Properties {
			paramName := convertToArgumentName(param.Name)
			dataType, isPrimitive := convertDataType(param.Type)
			isPackageType := isPackageType(dataType)
			isInterface := checkIsInterface(dataType, schema.Interfaces)

			if isPrimitive && !isPackageType {
				paramsStr += paramName + " " + dataType
			} else if isInterface || isPackageType {
				if strings.HasPrefix(dataType, "[]") {
					paramsStr += paramName + " []" + typePackageName + "." + dataType[len("[]"):]
				} else if strings.HasPrefix(dataType, "[][]") {
					paramsStr += paramName + " [][]" + typePackageName + "." + dataType[len("[][]"):]
				} else {
					if isInterface {
						paramsStr += paramName + " " + typePackageName + "." + dataType
					} else {
						paramsStr += paramName + " *" + typePackageName + "." + dataType
					}
				}
			} else {
				paramsStr += paramName + " *" + typePackageName + "." + dataType
			}

			if i < len(function.Properties)-1 {
				paramsStr += ", "
			}
			paramsDesc += "\n// @param " + paramName + " " + param.Description
		}

		buf.WriteString(fmt.Sprintf(`
		// %s %s %s
		func (client *Client) %s(%s) (%s%s, error)`, methodName, function.Description, paramsDesc, methodName,
			paramsStr, asterike, returnType))

		paramsStr = ""
		for i, param := range function.Properties {
			paramName := convertToArgumentName(param.Name)

			paramsStr += fmt.Sprintf("\"%s\":   %s,", param.Name, paramName)
			if i < len(function.Properties)-1 {
				paramsStr += "\n"
			}
		}

		illStr := `fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])`
		if strings.Contains(paramsStr, returnTypeCamel) {
			returnTypeCamel = returnTypeCamel + "Dummy"
		}
		if returnIsInterface {
			enumType := returnType[len(typePackageName+"."):] + "Enum"
			casesStr := ""

			for _, enumInfo := range schema.Enums {
				if enumInfo.EnumType == enumType {
					for _, item := range enumInfo.Items {
						casesStr += fmt.Sprintf(`
						case %s:
							var %s %s
							err = json.Unmarshal(result.Raw, &%s)
							return &%s, err
							`, typePackageName+"."+item.GolangType+"Type", returnTypeCamel, typePackageName+"."+item.GolangType, returnTypeCamel,
							returnTypeCamel)
					}
					break
				}
			}

			buf.WriteString(fmt.Sprintf(` {
			result, err := client.SendAndCatch(%s.UpdateData{
				"@type":       "%s",
				%s
			})

			if err != nil {
				return nil, err
			}

			if result.Data["@type"].(string) == "error" {
				return nil, %s
			}

			switch %s(result.Data["@type"].(string)) {
				%s
			default:
				return nil, fmt.Errorf("Invalid type")
			}
			}
			
			`, typePackageName, function.Name, paramsStr, illStr,
				typePackageName+"."+enumType, casesStr))

		} else {
			buf.WriteString(fmt.Sprintf(` {
			result, err := client.SendAndCatch(%s.UpdateData{
				"@type":       "%s",
				%s
			})

			if err != nil {
				return nil, err
			}

			if result.Data["@type"].(string) == "error" {
				return nil, %s
			}

			var %s %s.%s
			err = json.Unmarshal(result.Raw, &%s)
			return %s%s, err

			}
			
			`, typePackageName, function.Name, paramsStr, illStr, returnTypeCamel,
				typePackageName, returnType, returnTypeCamel, ampersign, returnTypeCamel))
		}

		err := writeToFileAndFormat(buf.Bytes(), filePath)
		if err != nil {
			return err
		}

		cmd := exec.Command("goimports", "-e", "-s", filePath)
		cmd.Run()
	}

	return nil
}
