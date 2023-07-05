package main

import (
	"context"
	_ "embed"
	"fmt"
	"log"

	"github.com/open-policy-agent/opa/rego"
)

//go:embed policy.rego
var policy string

type person struct {
	age  int
	name string
}

func main() {
	query := rego.Query("data.age_authorization") // corresponding with "package age_authorization" in policy.tego
	module := rego.Module("policy.rego", policy)
	authorizer, err := rego.New(query, module).PrepareForEval(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	kira := person{age: 22, name: "Kira"}
	crystal := person{age: 17, name: "Crystal"}

	checkIfAllowedDrinking(kira, authorizer)
	checkIfAllowedDrinking(crystal, authorizer)
}

func checkIfAllowedDrinking(p person, authorizer rego.PreparedEvalQuery) {
	input := map[string]interface{}{"age": p.age} // "age" will be referenced inside policy.rego
	result, err := authorizer.Eval(context.Background(), rego.EvalInput(input))
	if err != nil {
		log.Fatal(err)
	}

	who := fmt.Sprintf("%s,%d", p.name, p.age)
	log.Printf("[%-10s]: %+v\n", who, result[0].Expressions[0])
}
