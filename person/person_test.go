package person_test

import (
	"reflect"
	"testing"

	"github.com/taisa831/sandbox-go-impl/person"
)

func TestNewPerson(t *testing.T) {
	type args struct {
		firstName string
		LastName  string
	}
	tests := []struct {
		name string
		args args
		want *person.Person
	}{
		{
			"NewPerson",
			args{"Taro", "Yamada"},
			&person.Person{firstName: "Taro", lastName: "Yamada"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := person.NewPerson(tt.args.firstName, tt.args.LastName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetFirstName(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"GetFirstName",
			fields{FirstName:"Taro", LastName:"Yamada"},
			"Taro",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				firstName: tt.fields.FirstName,
				lastName:  tt.fields.LastName,
			}
			if got := p.GetFirstName(); got != tt.want {
				t.Errorf("GetFirstName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPerson_GetLastName(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"GetLastName",
			fields{FirstName:"Taro", LastName:"Yamada"},
			"Yamada",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Person{
				firstName: tt.fields.FirstName,
				lastName:  tt.fields.LastName,
			}
			if got := p.GetLastName(); got != tt.want {
				t.Errorf("GetLastName() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Example() {
//	person := NewPerson("Taro", "Yamada")
//	fmt.Println(person.GetFirstName())
//	// output: Taro
//}
//
//func ExampleNewPerson() {
//	person := NewPerson("Taro", "Yamada")
//	fmt.Println(person.GetFirstName())
//	// output: Taro
//}
//
//func ExamplePerson_GetFirstName() {
//	person := NewPerson("Taro", "Yamada")
//	fmt.Println(person.GetFirstName())
//	// output: Taro
//}
//
//func ExamplePerson_GetLastName() {
//	person := NewPerson("Taro", "Yamada")
//	fmt.Println(person.GetLastName())
//	// output: Yamada
//}