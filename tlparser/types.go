package tlparser

type TlSchema struct {
	Enums      []*EnumInfo
	Classes    []*ClassInfo
	Interfaces []*InterfaceInfo
}

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

// EnumInfo ...
type EnumInfo struct {
	EnumType string   `json:"enumType"`
	Items    []string `json:"description"`
}
