package main

import (
	"../util"
	"github.com/prataprc/goparsec"
)

var openparan = parsec.Token(`\(`, "OPENPARAN")
var closeparan = parsec.Token(`\)`, "CLOSEPARAN")
var addop = parsec.Token(`\+`, "ADD")
var multop = parsec.Token(`\*`, "MULT")

func getPart1Parser() parsec.Parser {
	var expr parsec.Parser
	anyOp := parsec.OrdChoice(util.One2one, multop, addop)
	groupExpr := parsec.And(util.ExprNode, openparan, &expr, closeparan)
	value := parsec.OrdChoice(util.ExprValueNode, parsec.Int(), groupExpr)
	expr = parsec.And(
		func(ns []parsec.ParsecNode) parsec.ParsecNode {
			if len(ns) > 0 {
				val := ns[0].(int)
				for _, x := range ns[1].([]parsec.ParsecNode) {
					y := x.([]parsec.ParsecNode)
					n := y[1].(int)
					switch y[0].(*parsec.Terminal).Name {
					case "ADD":
						val += n
					case "MULT":
						val *= n
					}
				}
				return val
			}
			return nil
		},
		value,
		parsec.Kleene(nil, parsec.And(util.Many2many, anyOp, value)))

	return expr
}

func getPart2Parser() parsec.Parser {
	var expr parsec.Parser
	groupExpr := parsec.And(util.ExprNode, openparan, &expr, closeparan)
	term := parsec.OrdChoice(util.ExprValueNode, parsec.Int(), groupExpr)

	factor := parsec.And(
		func(ns []parsec.ParsecNode) parsec.ParsecNode {
			if len(ns) > 0 {
				val := ns[0].(int)
				for _, x := range ns[1].([]parsec.ParsecNode) {
					y := x.([]parsec.ParsecNode)
					n := y[1].(int)
					val += n
				}
				return val
			}
			return nil
		},
		term, parsec.Kleene(nil, parsec.And(util.Many2many, addop, term)))

	expr = parsec.And(
		func(ns []parsec.ParsecNode) parsec.ParsecNode {
			if len(ns) > 0 {
				val := ns[0].(int)
				for _, x := range ns[1].([]parsec.ParsecNode) {
					y := x.([]parsec.ParsecNode)
					n := y[1].(int)
					val *= n
				}
				return val
			}
			return nil
		},
		factor,
		parsec.Kleene(nil, parsec.And(util.Many2many, multop, factor)))

	return expr
}

func main() {

	parser1 := getPart1Parser()
	parser2 := getPart2Parser()

	count := 0
	count2 := 0
	for _, expression := range util.ReadInput("day18/input.txt") {

		s := parsec.NewScanner([]byte(expression))
		result, _ := parser1(s)
		count += result.(int)

		result2, _ := parser2(s)
		count2 += result2.(int)
	}
	println(count)
	println(count2)
}
