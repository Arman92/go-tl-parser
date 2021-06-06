package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Arman92/go-tl-parser/tlparser"
)

func GenerateInterfaceAndEnums(schema *tlparser.TlSchema, packageName, outputDir string) error {

	for _, interfaceInfo := range schema.Interfaces {
		buf := bytes.NewBufferString("")
		appendPackageName(buf, packageName)
		buf.WriteString(`
		
		import (
			"encoding/json"
			"fmt"
			"strconv"
			"strings"
		)
	
		`)

		interfaceInfo.Name = replaceKeyWords(interfaceInfo.Name)
		typesCases := ""

		buf.WriteString(fmt.Sprintf("// %s %s \ntype %s interface {\nGet%sEnum() %sEnum\n}\n\n",
			interfaceInfo.Name, interfaceInfo.Description, interfaceInfo.Name, interfaceInfo.Name, interfaceInfo.Name))

		for _, enum := range schema.Enums {

			// Only include related enums to this interface (abstract class)
			if enum.EnumType == interfaceInfo.Name+"Enum" {

				// Generate enum and it's types
				buf.WriteString(fmt.Sprintf(`
			// %s Alias for abstract %s 'Sub-Classes', used as constant-enum here
			type %s string
			`,
					enum.EnumType,
					strings.TrimSuffix(enum.EnumType, "Enum"),
					enum.EnumType))

				constStr := ""
				for _, item := range enum.Items {
					// Generate enum const items
					constStr += fmt.Sprintf(`%sType %s = "%s"%s`, item, enum.EnumType, firstLower(item), "\n")

					typeName := item
					typeNameCamel := firstLower(typeName)
					typesCases += fmt.Sprintf(`case %s:
						var %s %s
						err := json.Unmarshal(*rawMsg, &%s)
						return &%s, err
						
						`,
						item+"Type", typeNameCamel, typeName,
						typeNameCamel, typeNameCamel)

				}

				buf.WriteString(fmt.Sprintf(`
			// %s enums
			const (
				%s
			)`, strings.TrimSuffix(enum.EnumType, "Enum"), constStr))

				// We won't process other enums as it's unnecessary
				break
			}

			// Generate Unmarshal functions

		}

		buf.WriteString(fmt.Sprintf(`
		func unmarshal%s(rawMsg *json.RawMessage) (%s, error){

			if rawMsg == nil {
				return nil, nil
			}
			var objMap map[string]interface{}
			err := json.Unmarshal(*rawMsg, &objMap)
			if err != nil {
				return nil, err
			}

			switch %sEnum(objMap["@type"].(string)) {
				%s
			default:
				return nil, fmt.Errorf("Error UnMarshaling, unknown type:" +  objMap["@type"].(string))
			}
		}
		`, interfaceInfo.Name, interfaceInfo.Name, interfaceInfo.Name,
			typesCases))

		filePath := filepath.Join(outputDir, firstLower(replaceKeyWords(interfaceInfo.Name))+".go")
		err := writeToFileAndFormat(buf.Bytes(), filePath)
		if err != nil {
			return err
		}
	}

	return nil
}
