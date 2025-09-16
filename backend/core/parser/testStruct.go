package parser

// @Doc(desc="MyTestStruct is a sample struct for testing", author="Jane Doe")
// MyTestStruct is a sample struct for testing
type MyTestStruct struct {
	ID    int     `json:"id"`   //ID is the unique identifier
	Name  string  `json:"name"` //Name is the name of the entity
	Other *string //Other is an optional field
	//Items is a list of integers(only can above field)
	Items []int
}
