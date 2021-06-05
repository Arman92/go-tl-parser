package tlparser

import (
	"bufio"
	"io"
	"log"
	"strings"
)

func ParseInputSchema(reader io.Reader) (*TlSchema, error) {
	rd := bufio.NewReader(reader)
	var hitFunctions = false

	schema := &TlSchema{
		Enums:      []*EnumInfo{},
		Classes:    []*ClassInfo{},
		Interfaces: []*InterfaceInfo{},
	}

	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read file line error: %v", err)
			return nil, err
		}

		switch {
		case strings.HasPrefix(line, "//@description"):
			break
		case strings.HasPrefix(line, "//@class "):
			schema.Classes = append(schema.Classes, parseTlClass(line))
		}

		if strings.Contains(line, "---functions---") {
			hitFunctions = true
			continue
		}
	}
}

func parseTlClass(line string) *ClassInfo {
	class := &ClassInfo{
		Name:        "",
		Description: "",
	}

	classLineParts := strings.Split(line, "@")

	_, class.Name = parseProperty(classLineParts[1])
	_, class.Description = parseProperty(classLineParts[2])

	return class
}
