package main

import (
	"testing"
)

func Test_hasNestedFieldValue(t *testing.T) {
	type Street map[string]interface{} // e.g. name, number

	type Address struct {
		Street   Street
		PostCode string
	}

	type Person struct {
		Addresses []Address
	}

	type args struct {
		resource Person
		path     string
		value    interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check street",
			args: args{
				resource: Person{Addresses: []Address{{Street: Street{"Name": "klonowa"}}}},
				path:     "Addresses.0.Street.Name",
				value:    "klonowa",
			},
			want: true,
		},
		{
			name: "check number",
			args: args{
				resource: Person{Addresses: []Address{{Street: Street{}}, {Street: Street{"Number": 55}}}},
				path:     "Addresses.1.Street.Number",
				value:    55,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FieldAtPathMatches(tt.args.resource, tt.args.path, tt.args.value); got != tt.want {
				t.Errorf("hasNestedFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
