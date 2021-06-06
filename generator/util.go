package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"

	"github.com/Arman92/go-tl-parser/tlparser"
	"github.com/asaskevich/govalidator"
)

func writeToFile(input []byte, filePath string) error {

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(input)
	writer.Flush()

	return err
}

func writeToFileAndFormat(input []byte, filePath string) error {
	// write the file
	err := writeToFile(input, filePath)
	if err != nil {
		return err
	}

	// format files
	cmd := exec.Command("gofmt", "-w", filePath)
	cmd.Run()
	cmd = exec.Command("goimports", "-w", filePath)
	cmd.Run()

	return nil
}

func firstUpper(str string) string {
	for i, r := range str {
		return string(unicode.ToUpper(r)) + str[i+1:]
	}

	return str
}

func firstLower(str string) string {
	for i, r := range str {
		return string(unicode.ToLower(r)) + str[i+1:]
	}

	return str
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

func appendPackageName(buffer *bytes.Buffer, packageName string) {

	buffer.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))

}

func convertDataType(input string) (string, bool) {
	propType := ""
	isPrimitiveType := true

	if strings.HasPrefix(input, "vector") {
		input = "[]" + input[len("vector<"):len(input)-1]
		isPrimitiveType = true
	}
	if strings.HasPrefix(input, "[]vector") {
		input = "[][]" + input[len("[]vector<"):len(input)-1]

	}

	switch {
	case strings.Contains(input, "string") || strings.Contains(input, "int32") ||
		strings.Contains(input, "int64"):
		propType = strings.Replace(input, "int64", "JSONInt64", 1)
		break
	case strings.Contains(input, "Bool"):
		propType = strings.Replace(input, "Bool", "bool", 1)
		break
	case strings.Contains(input, "double"):
		propType = strings.Replace(input, "double", "float64", 1)
		break
	case strings.Contains(input, "int53"):
		propType = strings.Replace(input, "int53", "int64", 1)
		break
	case strings.Contains(input, "bytes"):
		propType = strings.Replace(input, "bytes", "[]byte", 1)
		break
	default:
		if strings.HasPrefix(input, "[][]") {
			propType = "[][]" + strings.ToUpper(input[len("[][]"):len("[][]")+1]) + input[len("[][]")+1:]
		} else if strings.HasPrefix(input, "[]") {
			propType = "[]" + strings.ToUpper(input[len("[]"):len("[]")+1]) + input[len("[]")+1:]
		} else {
			propType = strings.ToUpper(input[:1]) + input[1:]
			isPrimitiveType = false
		}
	}

	propType = replaceKeyWords(propType)

	return propType, isPrimitiveType
}

func checkIsInterface(input string, interfaces []*tlparser.InterfaceInfo) bool {
	for _, interfaceInfo := range interfaces {
		if interfaceInfo.Name == input || replaceKeyWords(interfaceInfo.Name) == input {
			return true
		}
	}

	return false
}

func convertToArgumentName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = replaceKeyWords(paramName)
	paramName = firstLower(paramName)
	paramName = strings.Replace(paramName, "type", "typeParam", 1)
	paramName = strings.Replace(paramName, "json", "jsonString", 1)

	return paramName
}
