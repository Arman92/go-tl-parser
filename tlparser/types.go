package tlparser

type TlSchema struct {
	Enums      []*EnumInfo
	Interfaces []*InterfaceInfo
	Classes    []*ClassInfo
	Functions  []*FunctionInfo
}

// ClassInfo holds info of a Class in .tl file
type ClassInfo struct {
	Name        string     `json:"name"`
	Properties  []Property `json:"properties"`
	Description string     `json:"description"`
	RootName    string     `json:"rootName"`
}

// FunctionInfo holds info of a function in .tl file
type FunctionInfo struct {
	Name          string     `json:"name"`
	Properties    []Property `json:"properties"`
	Description   string     `json:"description"`
	ReturnType    string     `json:"return_type"`
	IsSynchronous bool       `json:"is_synchronous"`
}

// Property holds info about properties of a class (or function)
type Property struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// InterfaceInfo equals to abstract base classes in .tl file
type InterfaceInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EnumInfoItem struct {
	OriginalType string `json:"original_type"`
	GolangType   string `json:"golang_type"`
}

// EnumInfo ...
type EnumInfo struct {
	EnumType string         `json:"enumType"`
	Items    []EnumInfoItem `json:"items"`
}
