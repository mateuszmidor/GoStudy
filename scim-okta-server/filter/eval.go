package filter

import (
	"fmt"

	"github.com/scim2/filter-parser/v2"
)

// EvalExpression returns true if resource satisfies expression
func EvalExpression(resource interface{}, expression filter.Expression) (bool, error) {
	// no filter means test passed
	if expression == nil {
		return true, nil
	}

	// walk the expression
	switch v := expression.(type) {
	case *filter.AttributeExpression: // filter=userName eq "Bravo"
		predicate, err := getComparePredicate(v.Operator)
		if err != nil {
			return false, err
		}
		attribute := DigAttribute(resource, v.AttributePath.String())
		return predicate(attribute, v.CompareValue), nil

	case *filter.NotExpression: // filter=not(userName eq "Bravo")
		result, err := EvalExpression(resource, v.Expression)
		if err != nil {
			return false, err
		}
		return !result, err

	case *filter.LogicalExpression: // filter=userName eq "Johny" or userName eq "Bravo"
		predicate, err := getLogicalPredicate(v.Operator)
		if err != nil {
			return false, err
		}
		l, err := EvalExpression(resource, v.Left)
		if err != nil {
			return false, err
		}
		r, err := EvalExpression(resource, v.Right)
		if err != nil {
			return false, err
		}
		return predicate(l, r), nil

	case *filter.ValuePath: // filter=emails[value co "gmail"], for lists
		attributeList := DigAttribute(resource, v.AttributePath.String()).([]map[string]interface{})
		for _, attribute := range attributeList {
			result, err := EvalExpression(attribute, v.ValueFilter)
			if err != nil {
				return false, err
			}
			if result {
				return true, nil
			}
		}
		return false, nil
	}
	return false, fmt.Errorf("unknown expression type: %T", expression)
}
