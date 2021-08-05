package generator

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Arman92/go-tl-parser/tlparser"
	"github.com/asaskevich/govalidator"
)

func GenerateClasses(schema *tlparser.TlSchema, packageName, outputDir string) error {

	for _, class := range schema.Classes {
		buf := bytes.NewBufferString("\n")
		filePath := filepath.Join(outputDir, firstLower(replaceKeyWords(class.RootName))+".go")

		// Only add go package and imports if file does not exist already.
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			buf.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))

			buf.WriteString(`
		
			import (
				"encoding/json"
				"fmt"
			)
		
			`)
		}

		structName := firstUpper((class.Name))
		structName = replaceKeyWords(structName)
		structNameCamel := firstLower(structName)

		hasInterfaceProps := false
		propsStr := ""
		propsStrWithoutInterfaceOnes := ""
		assignStr := fmt.Sprintf("%s.tdCommon = tempObj.tdCommon\n", structNameCamel)
		assignInterfacePropsStr := ""

		for i, prop := range class.Properties {
			propName := govalidator.UnderscoreToCamelCase(prop.Name)
			propName = replaceKeyWords(propName)

			dataType, isPrimitive := convertDataType(prop.Type)
			propsStrItem := ""
			if isPrimitive || checkIsInterface(dataType, schema.Interfaces) {
				propsStrItem += fmt.Sprintf("%s %s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
			} else {
				propsStrItem += fmt.Sprintf("%s *%s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
			}
			if i < len(class.Properties)-1 {
				propsStrItem += "\n"
			}

			propsStr += propsStrItem
			if !checkIsInterface(prop.Type, schema.Interfaces) {
				propsStrWithoutInterfaceOnes += propsStrItem
				assignStr += fmt.Sprintf("%s.%s = tempObj.%s\n", structNameCamel, propName, propName)
			} else {
				hasInterfaceProps = true
				assignInterfacePropsStr += fmt.Sprintf(`
					field%s, _  := 	unmarshal%s(objMap["%s"])
					%s.%s = field%s
					`,
					propName, dataType, prop.Name,
					structNameCamel, propName, propName)
			}
		}

		buf.WriteString(fmt.Sprintf("// %s %s \ntype %s struct {\n"+
			"tdCommon\n"+
			"%s\n"+
			"}\n\n", structName, class.Description, structName, propsStr))

		buf.WriteString(fmt.Sprintf("// MessageType return the string telegram-type of %s \nfunc (%s *%s) MessageType() string {\n return \"%s\" }\n\n",
			structName, structNameCamel, structName, class.Name))

		paramsStr := ""
		paramsDesc := ""
		assingsStr := ""
		for i, param := range class.Properties {
			propName := govalidator.UnderscoreToCamelCase(param.Name)
			propName = replaceKeyWords(propName)
			dataType, isPrimitive := convertDataType(param.Type)
			paramName := convertToArgumentName(param.Name)

			if isPrimitive || checkIsInterface(dataType, schema.Interfaces) {
				paramsStr += paramName + " " + dataType

			} else { // if is not a primitive, use pointers
				paramsStr += paramName + " *" + dataType
			}

			if i < len(class.Properties)-1 {
				paramsStr += ", "
			}
			paramsDesc += "\n// @param " + paramName + " " + param.Description

			assingsStr += fmt.Sprintf("%s : %s,\n", propName, paramName)
		}

		buf.WriteString(fmt.Sprintf(`
		// New%s creates a new %s
		// %s
		func New%s(%s) *%s {
			%sTemp := %s {
				tdCommon: tdCommon {Type: "%s"},
				%s
			}

			return &%sTemp
		}
		`, structName, structName, paramsDesc,
			structName, paramsStr, structName, structNameCamel,
			structName, class.Name, assingsStr, structNameCamel))

		if hasInterfaceProps {
			buf.WriteString(fmt.Sprintf(`
		// UnmarshalJSON unmarshal to json
		func (%s *%s) UnmarshalJSON(b []byte) error {
			var objMap map[string]*json.RawMessage
			err := json.Unmarshal(b, &objMap)
			if err != nil {
				return err
			}
			tempObj := struct {
				tdCommon
				%s
			}{}
			err = json.Unmarshal(b, &tempObj)
			if err != nil {
				return err
			}

			%s

			%s	
			
			return nil
		}
		`, structNameCamel, structName, propsStrWithoutInterfaceOnes,
				assignStr, assignInterfacePropsStr))
		}

		if checkIsInterface(class.RootName, schema.Interfaces) {
			rootName := replaceKeyWords(class.RootName)
			buf.WriteString(fmt.Sprintf(`
				// Get%sEnum return the enum type of this object 
				func (%s *%s) Get%sEnum() %sEnum {
					 return %s 
				}

				`,
				rootName,
				firstLower(structName),
				structName, rootName, rootName,
				structName+"Type"))
		}

		err := writeToFileAndFormat(buf.Bytes(), filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
