// Package person
package person

type Person struct {
	firstName string
	lastName  string
}

// NewPerson
func NewPerson(firstName string, LastName string) *Person {
	return &Person{firstName: firstName, lastName: LastName}
}

// GetFirstName
func (p *Person) GetFirstName() string {
	return p.firstName
}


func (p *Person) GetLastName() string {
	return p.lastName
}
