package models

//Serve frontend docs show and trial
//Help LLMs suggest matching structs
//

type FieldMeta struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Tag     string `json:"tag"`
	Comment string `json:"comment"`
}

type MethodMeta struct {
	Name string `json:"name"`
	Doc  string `json:"doc"`
}

type StructMeta struct {
	Name    string       `json:"name"`
	Doc     string       `json:"doc"`
	Fields  []FieldMeta  `json:"fields"`
	Methods []MethodMeta `json:"methods"`
}
