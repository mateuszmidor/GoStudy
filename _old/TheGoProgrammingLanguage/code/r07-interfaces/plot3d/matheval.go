package main

import (
	"fmt"
	"math"

	"github.com/apaxa-go/eval"
	"github.com/apaxa-go/helper/goh/constanth"
)

// MathEvalXY evaluates value of function(x, y)
type MathEvalXY struct {
	expr *eval.Expression
	args *eval.Args
}

// NewEvalXY is MathEvalXY construtor
func NewEvalXY(formula string) (MathEvalXY, error) {
	// empty formula is not supported
	if formula == "" {
		return MathEvalXY{}, fmt.Errorf("Empty formula")
	}

	// cast to formula to float64 so result is always float64
	float64Formula := fmt.Sprintf("float64(%s)", formula)

	// try parse formula
	expr, err := eval.ParseString(float64Formula, "")

	// register standard functions
	args := &eval.Args{
		"sin":   eval.MakeDataRegularInterface(math.Sin),
		"cos":   eval.MakeDataRegularInterface(math.Cos),
		"pow":   eval.MakeDataRegularInterface(math.Pow),
		"hypot": eval.MakeDataRegularInterface(math.Hypot),
	}
	return MathEvalXY{expr, args}, err
}

// Eval calculates f(x, y)
func (e MathEvalXY) Eval(x, y float64) (float64, error) {
	args := *e.args
	args["x"] = eval.MakeDataUntypedConst(constanth.MakeFloat64(x))
	args["y"] = eval.MakeDataUntypedConst(constanth.MakeFloat64(y))
	result, err := e.expr.EvalToInterface(args)
	if err != nil {
		return 0.0, err
	}
	return result.(float64), nil
}
