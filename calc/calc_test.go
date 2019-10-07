package main

import (
	"testing"
)

type testPairCalc struct {
	expr string
	value int
}

type testPairBackspace struct {
	expr string
	exprBackspace string
}

type testPairEval struct {
	expr []string
	value int
}

var testsCalc = []testPairCalc{
	{ "45-45/5+8", 44 },
	{ "-100/5", -20 },
	{ "100-11*9", 1 },
	{ "(100-99)*9", 9},
	{ "-3", -3},
	{ "1000/(100-99)*4", 4000},
}

var testsBackspace = []testPairBackspace{
	{ "45-45/5+8", "45 - 45 / 5 + 8" },
	{ "-100/5", "-100 / 5" },
	{ "100-11*9", "100 - 11 * 9" },
	{ "(100-99)*9", "( 100 - 99 ) * 9"},
	{ "-3", "-3" },
	{ "1000/(100-99)*4", "1000 / ( 100 - 99 ) * 4"},
}

var testsEval = []testPairEval{
	{ []string{"45", "45", "5", "/", "-", "8", "+"}, 44 },
	{ []string{"-100", "5", "/"}, -20 },
	{ []string{"100", "11", "9", "*", "-"}, 1 },
	{ []string{"100", "99", "-", "9", "*"}, 9},
	{ []string{"-3"}, -3},
	{ []string{"1000", "100", "99", "-", "/", "4", "*"}, 4000},
}

func TestCalc (t *testing.T) {
	for _, pair := range testsCalc {
		v := Calc(pair.expr)
		if v != pair.value {
			t.Error(
				"For", pair.expr,
				"expected", pair.value,
				"got", v,
			)
		}
	}
}

func TestToBackspace(t *testing.T) {
	for _, pair := range testsBackspace {
		v := Backspace(pair.expr)
		if v != pair.exprBackspace {
			t.Error(
				"For", pair.expr,
				"expected", pair.exprBackspace,
				"got", v,
				)
		}
	}
}

func TestEval(t *testing.T) {
	for _, pair := range testsEval {
		v := Eval(pair.expr)
		if v != pair.value {
			t.Error(
				"For", pair.expr,
				"expected", pair.value,
				"got", v,
			)
		}
	}
}