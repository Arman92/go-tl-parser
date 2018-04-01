package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/asaskevich/govalidator"
)

const (
	generatedPackageDefault = "tdlib"      // the package in which go types and methods will be placed into
	structsFileNameDefault  = "types.go"   // file name for structs
	methodsFileNameDefault  = "methods.go" // file name for methods
	generateDirDefault      = "./tdlib"
	modeDefault             = GenerateModeJSON
)

// GenerateMode indicates whether to create a json file of .tl file, or create .go files
type GenerateMode int

// GenerateModeJSON for generate json, GenerateModeGolang for generate golang
const (
	GenerateModeJSON GenerateMode = iota
	GenerateModeGolang
)

// ClassInfo holds info of a Class in .tl file
type ClassInfo struct {
	Name        string          `json:"name"`
	Properties  []ClassProperty `json:"properties"`
	Description string          `json:"description"`
	RootName    string          `json:"rootName"`
	IsFunction  bool            `json:"isFunction"`
}

// ClassProperty holds info about properties of a class (or function)
type ClassProperty struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// InterfaceInfo equals to abstract base classes in .tl file
type InterfaceInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func main() {

	var inputFile string
	var generatedPackage string
	var structsFileName string
	var methodsFileName string
	var generateDir string
	var generateMode GenerateMode

	flag.StringVar(&inputFile, "file", "./schema.tl", ".tl schema file")
	flag.StringVar(&generateDir, "dir", generateDirDefault, "Generate directory")
	flag.StringVar(&generatedPackage, "package", generatedPackageDefault, "Package in which generated files will be a part of")
	flag.StringVar(&structsFileName, "structs-file", structsFileNameDefault, "file name for structs")
	flag.StringVar(&methodsFileName, "methods-file", methodsFileNameDefault, "file name for methods")
	flag.IntVar((*int)(&generateMode), "mode", int(modeDefault), "Generate mod, indicates whether to create json file, or golang files")

	flag.Parse()

	f, err := os.OpenFile(inputFile, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer f.Close()

	entityDesc := ""
	var paramDescs map[string]string
	var params map[string]string
	var classInfoes []ClassInfo
	var interfaceInfoes []InterfaceInfo
	classInfoes = make([]ClassInfo, 0, 10)
	interfaceInfoes = make([]InterfaceInfo, 0, 1)

	paramDescs = make(map[string]string)
	params = make(map[string]string)
	hitFunctions := false

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return
		}
		if strings.Contains(line, "---functions---") {
			hitFunctions = true
			continue
		}

		if strings.HasPrefix(line, "//@class ") {
			line = line[len("//@class "):]
			interfaceName := line[:strings.Index(line, " ")]
			line = line[len(interfaceName):]
			line = line[len(" @description "):]
			entityDesc = line[:len(line)-1]
			interfaceInfo := InterfaceInfo{
				Name:        interfaceName,
				Description: entityDesc,
			}
			interfaceInfoes = append(interfaceInfoes, interfaceInfo)

		} else if strings.HasPrefix(line, "//@description ") { // Entity description
			line = line[len("//@description "):]
			indexOfFirstSign := strings.Index(line, "@")

			entityDesc = line[:len(line)-1]
			if indexOfFirstSign != -1 {
				entityDesc = line[:indexOfFirstSign]
			}

			if indexOfFirstSign != -1 { // there is some parameter description inline, parse them
				line = line[indexOfFirstSign+1:]
				rd2 := bufio.NewReader(strings.NewReader(line))
				for {
					paramName, _ := rd2.ReadString(' ')
					if paramName == "" {
						break
					}
					paramName = paramName[:len(paramName)-1]
					paramDesc, _ := rd2.ReadString('@')
					if paramDesc == "" {
						paramDesc, _ = rd2.ReadString('\n')

						paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
						break
					}

					paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
				}
			}
		} else if entityDesc != "" && strings.HasPrefix(line, "//@") {
			line = line[len("//@"):]
			rd2 := bufio.NewReader(strings.NewReader(line))
			for {
				paramName, _ := rd2.ReadString(' ')
				if paramName == "" {
					break
				}
				paramName = paramName[:len(paramName)-1]
				paramDesc, _ := rd2.ReadString('@')
				if paramDesc == "" {
					paramDesc, _ = rd2.ReadString('\n')

					paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
					break
				}

				paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
			}

		} else if entityDesc != "" && !strings.HasPrefix(line, "//") && len(line) > 2 {
			entityName := line[:strings.Index(line, " ")]

			line = line[len(entityName)+1:]
			for {
				if strings.Index(line, ":") == -1 {
					break
				}
				paramName := line[:strings.Index(line, ":")]
				line = line[len(paramName)+1:]
				paramType := line[:strings.Index(line, " ")]
				params[paramName] = paramType
				line = line[len(paramType)+1:]
			}

			rootName := line[len("= ") : len(line)-2]

			var classProps []ClassProperty
			classProps = make([]ClassProperty, 0, 0)
			for paramName, paramType := range params {
				classProp := ClassProperty{
					Name:        paramName,
					Type:        paramType,
					Description: paramDescs[paramName],
				}
				classProps = append(classProps, classProp)
			}
			classInfoe := ClassInfo{
				Name:        entityName,
				Description: entityDesc,
				RootName:    rootName,
				Properties:  classProps,
				IsFunction:  hitFunctions,
			}

			classInfoes = append(classInfoes, classInfoe)
			entityDesc = ""
			paramDescs = make(map[string]string)
			params = make(map[string]string)
		}
	}

	// Write results in a json file
	os.MkdirAll(generateDir, os.ModePerm)
	jsonFile, err := os.OpenFile(generateDir+"/schema.json", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalf("error openning file %v", err)
	}
	jsonBytes, _ := json.Marshal(classInfoes)
	w := bufio.NewWriter(jsonFile)
	w.Write(jsonBytes)

	gnrtdStructs := fmt.Sprintf("package %s\n\n", generatedPackage)
	gnrtdMethods := fmt.Sprintf("package %s\n\n", generatedPackage)
	gnrtdMethods += `
	
	import (
		"fmt"
		"github.com/mitchellh/mapstructure"
	)
	
	`

	for _, interfaceInfo := range interfaceInfoes {
		interfaceInfo.Name = replaceKeyWords(interfaceInfo.Name)

		gnrtdStructs += fmt.Sprintf("// %s %s \ntype %s interface {\nMessageType() string\n}\n\n",
			interfaceInfo.Name, interfaceInfo.Description, interfaceInfo.Name)
	}

	gnrtdStructs += "type tdCommon struct {\n" +
		"Type string `json:\"@type\"`\n" +
		"Extra string `json:\"extra\"`\n" +
		"}\n\n"

	for _, classInfoe := range classInfoes {
		if !classInfoe.IsFunction {
			structName := strings.ToUpper(classInfoe.Name[:1]) + classInfoe.Name[1:]
			structName = replaceKeyWords(structName)

			propsStr := ""
			for i, param := range classInfoe.Properties {
				propName := govalidator.UnderscoreToCamelCase(param.Name)
				propName = replaceKeyWords(propName)

				propsStr += fmt.Sprintf("%s %s `json:\"%s\"` // %s", propName, convertDataType(param.Type), param.Name, param.Description)
				if i < len(classInfoe.Properties)-1 {
					propsStr += "\n"
				}
			}
			gnrtdStructs += fmt.Sprintf("// %s %s \ntype %s struct {\n"+
				"tdCommon\n"+
				"%s\n"+
				"}\n\n", structName, classInfoe.Description, structName, propsStr)

			gnrtdStructs += fmt.Sprintf("// MessageType return the string telegram-type of %s \nfunc (%s *%s) MessageType() string {\n return \"%s\" }\n\n",
				structName, strings.ToLower(structName[0:1])+structName[1:], structName, classInfoe.Name)
		} else {
			methodName := strings.ToUpper(classInfoe.Name[:1]) + classInfoe.Name[1:]
			methodName = replaceKeyWords(methodName)
			returnType := strings.ToUpper(classInfoe.RootName[:1]) + classInfoe.RootName[1:]
			returnType = replaceKeyWords(returnType)
			returnTypeCamel := strings.ToLower(returnType[:1]) + returnType[1:]

			paramsStr := ""
			paramsDesc := ""
			for i, param := range classInfoe.Properties {
				paramName := convertToArgumentName(param.Name)
				paramsStr += paramName + " "
				paramsStr += convertDataType(param.Type)
				if i < len(classInfoe.Properties)-1 {
					paramsStr += ", "
				}
				paramsDesc += "\n// @param " + paramName + " " + param.Description
			}

			gnrtdMethods += fmt.Sprintf(`
				// %s %s %s
				func (client *Client) %s(%s) (*%s, error)`, methodName, classInfoe.Description, paramsDesc, methodName,
				paramsStr, returnType)

			paramsStr = ""
			for i, param := range classInfoe.Properties {
				paramName := convertToArgumentName(param.Name)

				paramsStr += fmt.Sprintf("\"%s\":   %s,", param.Name, paramName)
				if i < len(classInfoe.Properties)-1 {
					paramsStr += "\n"
				}
			}

			illStr := `fmt.Errorf("error! code: %d msg: %s", result["code"], result["message"])`
			if strings.Contains(paramsStr, returnTypeCamel) {
				returnTypeCamel = returnTypeCamel + "Dummy"
			}
			gnrtdMethods += fmt.Sprintf(` {
				result, err := client.SendAndCatch(UpdateMsg{
					"@type":       "%s",
					%s
				})

				if err != nil {
					return nil, err
				}

				if result["@type"].(string) == "error" {
					return nil, %s
				}

				var %s %s
				err = mapstructure.Decode(result, &%s)
				return &%s, err

				}
				
				`, classInfoe.Name, paramsStr, illStr, returnTypeCamel,
				returnType, returnTypeCamel, returnTypeCamel)
		}
	}

	os.Remove(generateDir + "/" + structsFileName)
	structsFile, err := os.OpenFile(generateDir+"/"+structsFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer structsFile.Close()
	if err != nil {
		log.Fatalf("error openning file %v", err)
	}
	wgo := bufio.NewWriter(structsFile)
	wgo.Write([]byte(gnrtdStructs))

	os.Remove(generateDir + "/" + methodsFileName)
	methodsFile, err := os.OpenFile(generateDir+"/"+methodsFileName, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer methodsFile.Close()

	if err != nil {
		log.Fatalf("error openning file %v", err)
	}
	wgo = bufio.NewWriter(methodsFile)
	wgo.Write([]byte(gnrtdMethods))

	// format files
	cmd := exec.Command("gofmt", "-w", generateDir+"/"+methodsFileName)
	cmd.Run()

	cmd = exec.Command("gofmt", "-w", generateDir+"/"+structsFileName)
	cmd.Run()
}

func replaceKeyWords(input string) string {
	input = strings.Replace(input, "Api", "API", -1)
	input = strings.Replace(input, "Url", "URL", -1)
	input = strings.Replace(input, "Id", "ID", -1)
	input = strings.Replace(input, "Ttl", "TTL", -1)
	input = strings.Replace(input, "Html", "HTML", -1)
	input = strings.Replace(input, "Uri", "URI", -1)
	input = strings.Replace(input, "Ip", "IP", -1)
	input = strings.Replace(input, "Udp", "UDP", -1)

	return input
}

func convertDataType(input string) string {
	propType := ""

	if strings.HasPrefix(input, "vector") {
		input = "[]" + input[len("vector<"):len(input)-1]
	}
	if strings.HasPrefix(input, "[]vector") {
		input = "[][]" + input[len("[]vector<"):len(input)-1]
	}
	if strings.Contains(input, "string") || strings.Contains(input, "int32") ||
		strings.Contains(input, "int64") {
		propType = input // do nothing
	} else if strings.Contains(input, "Bool") {
		propType = strings.Replace(input, "Bool", "bool", 1)
	} else if strings.Contains(input, "double") {
		propType = strings.Replace(input, "double", "float64", 1)
	} else if strings.Contains(input, "int53") {
		propType = strings.Replace(input, "int53", "int64", 1)
	} else if strings.Contains(input, "bytes") {
		propType = strings.Replace(input, "bytes", "[]byte", 1)
	} else {
		if strings.HasPrefix(input, "[][]") {
			propType = "[][]" + strings.ToUpper(input[len("[][]"):len("[][]")+1]) + input[len("[][]")+1:]
		} else if strings.HasPrefix(input, "[]") {
			propType = "[]" + strings.ToUpper(input[len("[]"):len("[]")+1]) + input[len("[]")+1:]
		} else {
			propType = strings.ToUpper(input[:1]) + input[1:]
		}
	}

	propType = replaceKeyWords(propType)

	return propType
}

func convertToArgumentName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = replaceKeyWords(paramName)
	paramName = strings.ToLower(paramName[0:1]) + paramName[1:]
	paramName = strings.Replace(paramName, "type", "typeParam", 1)

	return paramName
}
