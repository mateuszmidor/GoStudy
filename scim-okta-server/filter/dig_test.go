package filter_test

import (
	"testing"
)

func Test_DigAttribute(t *testing.T) {
	// type StreetDetails struct {
	// 	Name   string
	// 	Number int
	// }

	// type Address struct {
	// 	Street   StreetDetails
	// 	PostCode string
	// }
	// type Person struct {
	// 	Addresses []Address
	// }

	// type args struct {
	// 	resource Person
	// 	path     string
	// 	value    interface{}
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want bool
	// }{
	// 	{
	// 		name: "check street",
	// 		args: args{
	// 			resource: Person{Addresses: []Address{{Street: StreetDetails{Name: "klonowa"}}}},
	// 			path:     "Addresses.0.Street.Name",
	// 			value:    "klonowa",
	// 		},
	// 		want: true,
	// 	},
	// 	{
	// 		name: "check number",
	// 		args: args{
	// 			resource: Person{Addresses: []Address{{Street: StreetDetails{}}, {Street: StreetDetails{Number: 55}}}},
	// 			path:     "Addresses.1.Street.Number",
	// 			value:    55,
	// 		},
	// 		want: true,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := filter.DigAttribute(tt.args.resource, tt.args.path); got != tt.want {
	// 			t.Errorf("hasNestedFieldValue() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
