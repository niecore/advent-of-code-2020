package util

import (
	parsec "github.com/prataprc/goparsec"
	"strconv"
)

func One2one(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns[0]
}

func Many2many(ns []parsec.ParsecNode) parsec.ParsecNode {
	if ns == nil || len(ns) == 0 {
		return nil
	}
	return ns
}

func ExprValueNode(ns []parsec.ParsecNode) parsec.ParsecNode {
	if len(ns) == 0 {
		return nil
	} else if term, ok := ns[0].(*parsec.Terminal); ok {
		val, _ := strconv.Atoi(term.Value)
		return val
	}
	return ns[0]
}

func ExprNode(ns []parsec.ParsecNode) parsec.ParsecNode {
	if len(ns) == 0 {
		return nil
	}
	return ns[1]
}
