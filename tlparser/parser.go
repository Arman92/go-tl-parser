package tlparser

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ParseInputSchema(reader io.Reader) (*TlSchema, error) {
	scanner := bufio.NewScanner(reader)
	var hitFunctions = false

	schema := &TlSchema{
		Enums:      []*EnumInfo{},
		Classes:    []*ClassInfo{},
		Interfaces: []*InterfaceInfo{},
		Functions:  []*FunctionInfo{},
	}

	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "//@description"):
			if hitFunctions {
				schema.Functions = append(schema.Functions, parseFunction(line, scanner))
			} else {
				typeInfo := parseType(line, scanner)

				schema.Classes = append(schema.Classes, typeInfo)
				// Append to enum Items if this is sub-class of an abstract class.
				for _, enumInfo := range schema.Enums {
					if strings.TrimRight(enumInfo.EnumType, "Enum") == typeInfo.RootName {
						enumInfo.Items = append(enumInfo.Items, replaceKeyWords(strings.ToUpper(typeInfo.Name[0:1])+typeInfo.Name[1:]))

						break
					}
				}
				fmt.Print(typeInfo)
			}
		case strings.HasPrefix(line, "//@class "):
			interfaceInfo := parseTlAbstractClass(line)
			schema.Interfaces = append(schema.Interfaces, interfaceInfo)

			enumInfo := &EnumInfo{EnumType: replaceKeyWords(interfaceInfo.Name) + "Enum"}
			schema.Enums = append(schema.Enums, enumInfo)

		case line == "":

		case strings.Contains(line, "---functions---"):
			hitFunctions = true

		}

	}

	return schema, nil
}

func parseTlAbstractClass(line string) *InterfaceInfo {
	tlInterface := &InterfaceInfo{
		Name:        "",
		Description: "",
	}

	interfaceLineParts := strings.Split(line, "@")

	_, tlInterface.Name = parseProperty(interfaceLineParts[1])
	_, tlInterface.Description = parseProperty(interfaceLineParts[2])

	return tlInterface
}

func parseType(firstLine string, scanner *bufio.Scanner) *ClassInfo {
	name, description, class, properties, _ := parseEntity(firstLine, scanner)
	return &ClassInfo{
		Name:        name,
		Description: description,
		RootName:    class,
		Properties:  properties,
	}
}

func parseFunction(firstLine string, scanner *bufio.Scanner) *FunctionInfo {
	name, description, class, properties, isSynchronous := parseEntity(firstLine, scanner)
	return &FunctionInfo{
		Name:          name,
		Description:   description,
		ReturnType:    class,
		Properties:    properties,
		IsSynchronous: isSynchronous,
	}
}

func parseEntity(firstLine string, scanner *bufio.Scanner) (string, string, string, []Property, bool) {
	name := ""
	description := ""
	rootClass := ""
	properties := []Property{}

	propertiesLine := strings.TrimPrefix(firstLine, "//")

Loop:
	for scanner.Scan() {
		line := scanner.Text()

		switch {
		case strings.HasPrefix(line, "//@"):
			propertiesLine += " " + strings.TrimPrefix(line, "//")

		case strings.HasPrefix(line, "//-"):
			propertiesLine += " " + strings.TrimPrefix(line, "//-")

		default:
			bodyFields := strings.Fields(line)
			name = bodyFields[0]

			for _, rawProperty := range bodyFields[1 : len(bodyFields)-2] {
				propertyParts := strings.Split(rawProperty, ":")
				property := Property{
					Name: propertyParts[0],
					Type: propertyParts[1],
				}
				properties = append(properties, property)
			}
			rootClass = strings.TrimRight(bodyFields[len(bodyFields)-1], ";")
			break Loop
		}
	}

	rawProperties := strings.Split(propertiesLine, "@")
	for _, rawProperty := range rawProperties[1:] {
		name, value := parseProperty(rawProperty)
		switch {
		case name == "description":
			description = value
		default:
			name = strings.TrimPrefix(name, "param_")
			property := getProperty(properties, name)
			property.Description = value

		}
	}

	return name, description, rootClass, properties, strings.Contains(description, "Can be called synchronously")
}

func parseProperty(str string) (string, string) {
	strParts := strings.Fields(str)

	return strParts[0], strings.Join(strParts[1:], " ")
}

func getProperty(properties []Property, name string) *Property {
	for _, property := range properties {
		if property.Name == name {
			return &property
		}
	}

	return nil
}
