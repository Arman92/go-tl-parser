package generator

import (
	"bytes"
	"fmt"

	"github.com/Arman92/go-tl-parser/tlparser"
)

func GenerateCode(schema *tlparser.TlSchema, packageName, outputDir string) []byte {
	buf := bytes.NewBufferString("")

	buf.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))
	buf.WriteString(`
	
	import (
		"encoding/json"
		"fmt"
		"strconv"
		"strings"
	)
	

	type tdCommon struct {
		Type string ` + `json:"@type"` + `
		Extra string ` + `json:"@extra"` + `
	}
	`)

	return buf.Bytes()
}
