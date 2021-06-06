package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Arman92/go-tl-parser/tlparser"
)

func GenerateFunctions(schema *tlparser.TlSchema, packageName, outputDir string) error {
	for _, function := range schema.Functions {
		buf := bytes.NewBufferString("\n")
		filePath := filepath.Join(outputDir, firstLower(replaceKeyWords(function.ReturnType))+".go")

		// Only add go package and imports if file does not exist already.
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			buf.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))

			buf.WriteString(`
		
			import (
				"encoding/json"
				"fmt"
				"strconv"
				"strings"
			)
		
			`)
		}

		methodName := firstUpper(function.Name)
		methodName = replaceKeyWords(methodName)
		returnType := strings.ToUpper(function.ReturnType[:1]) + function.ReturnType[1:]
		returnType = replaceKeyWords(returnType)
		returnTypeCamel := strings.ToLower(returnType[:1]) + returnType[1:]
		returnIsInterface := checkIsInterface(returnType, schema.Interfaces)

		asterike := "*"
		ampersign := "&"
		if returnIsInterface {
			asterike = ""
			ampersign = ""
		}

		paramsStr := ""
		paramsDesc := ""
		for i, param := range function.Properties {
			paramName := convertToArgumentName(param.Name)
			dataType, isPrimitive := convertDataType(param.Type)
			if isPrimitive || checkIsInterface(dataType, schema.Interfaces) {
				paramsStr += paramName + " " + dataType

			} else {
				paramsStr += paramName + " *" + dataType
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
			enumType := returnType + "Enum"
			casesStr := ""

			for _, enumInfoe := range schema.Enums {
				if enumInfoe.EnumType == enumType {
					for _, item := range enumInfoe.Items {
						casesStr += fmt.Sprintf(`
						case %s:
							var %s %s
							err = json.Unmarshal(result.Raw, &%s)
							return &%s, err
							`, item+"Type", returnTypeCamel, item, returnTypeCamel,
							returnTypeCamel)
					}
					break
				}
			}

			buf.WriteString(fmt.Sprintf(` {
			result, err := client.SendAndCatch(UpdateData{
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
			
			`, function.Name, paramsStr, illStr,
				enumType, casesStr))

		} else {
			buf.WriteString(fmt.Sprintf(` {
			result, err := client.SendAndCatch(UpdateData{
				"@type":       "%s",
				%s
			})

			if err != nil {
				return nil, err
			}

			if result.Data["@type"].(string) == "error" {
				return nil, %s
			}

			var %s %s
			err = json.Unmarshal(result.Raw, &%s)
			return %s%s, err

			}
			
			`, function.Name, paramsStr, illStr, returnTypeCamel,
				returnType, returnTypeCamel, ampersign, returnTypeCamel))
		}

		err := writeToFileAndFormat(buf.Bytes(), filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
