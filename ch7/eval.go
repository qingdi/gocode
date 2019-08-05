package ch7

import (
	"fmt"
	"testing"
)

type Var string
type literal float64
type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

type unary struct {
	op rune
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn  string
	arg []Expr
}

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"-F", Env{"F": 100}, "-100"},
	}

	var preExpr string
	for _, test := range tests {
		if test.expr != preExpr {
			fmt.Printf("\n%s\n", test.expr)
			preExpr = test.expr
		}
		//io.Writer()
	}
}
