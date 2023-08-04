package filter

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/scim2/filter-parser/v2"
)

type comparePredicate func(a, b interface{}) bool
type logicalPredicate func(a, b bool) bool

func getComparePredicate(op filter.CompareOperator) (comparePredicate, error) {
	switch op {
	case filter.PR:
		return opPresent, nil
	case filter.EQ:
		return opEqual, nil
	case filter.NE:
		return opNotEqual, nil
	case filter.CO:
		return opContains, nil
		// // SW is an abbreviation for 'starts with'.
		// SW CompareOperator = "sw"
		// // EW an abbreviation for 'ends with'.
		// EW CompareOperator = "ew"
		// // GT is an abbreviation for 'greater than'.
		// GT CompareOperator = "gt"
		// // LT is an abbreviation for 'less than'.
		// LT CompareOperator = "lt"
		// // GE is an abbreviation for 'greater or equal than'.
		// GE CompareOperator = "ge"
		// // LE is an abbreviation for 'less or equal than'.
		// LE CompareOperator = "le"
	default:
		return nil, fmt.Errorf("unknown comparison operator: %s", op)
	}
}

func getLogicalPredicate(op filter.LogicalOperator) (logicalPredicate, error) {
	switch op {
	case filter.AND:
		return opAnd, nil
	case filter.OR:
		return opOr, nil
	default:
		return nil, fmt.Errorf("unknown logical operator: %s", op)
	}
}

func opPresent(a, b interface{}) bool {
	return a != nil
}

func opEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func opNotEqual(a, b interface{}) bool {
	return !reflect.DeepEqual(a, b)
}

func opContains(a, b interface{}) bool {
	aStr, ok := a.(string)
	if !ok {
		return false
	}
	bStr, ok := b.(string)
	if !ok {
		return false
	}
	return strings.Contains(aStr, bStr)
}

func opAnd(a, b bool) bool {
	return a && b
}

func opOr(a, b bool) bool {
	return a || b
}
