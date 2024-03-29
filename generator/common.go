package generator

import (
	"bytes"
	"fmt"
	"path/filepath"
)

func generateCommonFiles(packageName, outputDir string) error {
	buf := bytes.NewBufferString("")

	buf.WriteString(fmt.Sprintf("%s\n\npackage %s\n\n", defaultHeader, packageName))
	buf.WriteString(`
	
	import (
		"strconv"
		"strings"
	)

	type tdCommon struct {
		Type string ` + "`json:\"@type\"`" + `
		Extra string ` + "`json:\"@extra\"`" + `
	}

	// TdMessage is the interface for all messages send and received to/from tdlib
	type TdMessage interface{
		MessageType() string
	}

	// RequestError represents an error returned from tdlib.
	type RequestError struct {
		Code int
		Message string
	}

	func (re RequestError) Error() string {
		return "error! code: " + strconv.FormatInt(int64(re.Code), 10) + " msg: " + re.Message
	}

	// JSONInt64 alias for int64, in order to deal with json big number problem
	type JSONInt64 int64

	
	// UpdateData alias for use in UpdateMsg
	type UpdateData map[string]interface{}

	// UpdateMsg is used to unmarshal received json strings into
	type UpdateMsg struct {
		Data UpdateData
		Raw  []byte
	}
	
	// MarshalJSON marshals to json
	func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
		intStr := strconv.FormatInt(int64(*jsonInt), 10)
		return []byte(intStr), nil
	}
	
	// UnmarshalJSON unmarshals from json
	func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
		intStr := string(b)
		intStr = strings.Replace(intStr, "\"", "", 2)
		jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
		if err != nil {
			return err
		}
		*jsonInt = JSONInt64(jsonBigInt)
		return nil
	}

`)

	commonFilePath := filepath.Join(outputDir, commonFileName)

	return writeToFileAndFormat(buf.Bytes(), commonFilePath)
}
