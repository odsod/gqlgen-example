package resourcefilter

import (
	"fmt"

	"cloud.google.com/go/spanner/spansql"
	"github.com/google/cel-go/common/operators"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

func exprToSpannerSQL(e *expr.Expr) (spansql.Expr, error) {
	switch e := e.ExprKind.(type) {
	case *expr.Expr_CallExpr:
		switch e.CallExpr.Function {
		case operators.In:
			return inExprToSpannerSQL(e.CallExpr.Args)
		case operators.LogicalAnd:
			return binaryLogicalExprToSpannerSQL(spansql.And, e.CallExpr.Args)
		case operators.LogicalOr:
			return binaryLogicalExprToSpannerSQL(spansql.Or, e.CallExpr.Args)
		case operators.LogicalNot:
			return notExprToSpannerSQL(e.CallExpr.Args)
		default:
			return nil, fmt.Errorf("unsupported function: %s", e.CallExpr.Function)
		}
	case *expr.Expr_IdentExpr:
		return spansql.ID(e.IdentExpr.Name), nil
	case *expr.Expr_SelectExpr:
		return spansql.ID(e.SelectExpr.Field), nil
	case *expr.Expr_ConstExpr:
		switch k := e.ConstExpr.ConstantKind.(type) {
		case *expr.Constant_StringValue:
			return spansql.StringLiteral(k.StringValue), nil
		default:
			return nil, fmt.Errorf("unsupported const expr: %v", k)
		}
	default:
		return nil, fmt.Errorf("unsupported expr: %v", e)
	}
}

func notExprToSpannerSQL(args []*expr.Expr) (spansql.BoolExpr, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("unexpected number of arguments to `!` expression: %d", len(args))
	}
	rhsExpr, err := exprToSpannerSQL(args[0])
	if err != nil {
		return nil, err
	}
	rhsBoolExpr, ok := rhsExpr.(spansql.BoolExpr)
	if !ok {
		return nil, fmt.Errorf("unexpected argument to `!`: not a bool expr")
	}
	spanExpr := spansql.LogicalOp{
		Op:  spansql.Not,
		RHS: rhsBoolExpr,
	}
	return spanExpr, nil
}

func binaryLogicalExprToSpannerSQL(op spansql.LogicalOperator, args []*expr.Expr) (spansql.BoolExpr, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("unexpected number of arguments to `&&` expression: %d", len(args))
	}
	lhsExpr, err := exprToSpannerSQL(args[0])
	if err != nil {
		return nil, err
	}
	rhsExpr, err := exprToSpannerSQL(args[1])
	if err != nil {
		return nil, err
	}
	lhsBoolExpr, ok := lhsExpr.(spansql.BoolExpr)
	if !ok {
		return nil, fmt.Errorf("unexpected arguments to `&&`: lhs not a bool expr")
	}
	rhsBoolExpr, ok := rhsExpr.(spansql.BoolExpr)
	if !ok {
		return nil, fmt.Errorf("unexpected arguments to `&&`: rhs not a bool expr")
	}
	boolExpr := spansql.LogicalOp{
		Op:  op,
		LHS: lhsBoolExpr,
		RHS: rhsBoolExpr,
	}
	parenExpr := spansql.Paren{
		Expr: boolExpr,
	}
	return parenExpr, nil
}

func inExprToSpannerSQL(args []*expr.Expr) (spansql.BoolExpr, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("unexpected number of arguments to `in` expression: %d", len(args))
	}
	needle, err := exprToSpannerSQL(args[0])
	if err != nil {
		return nil, err
	}
	list, ok := args[1].ExprKind.(*expr.Expr_ListExpr)
	if !ok {
		return nil, fmt.Errorf("arg of `in` expression must be list")
	}
	if len(list.ListExpr.Elements) == 0 {
		return nil, fmt.Errorf("`in` expression requires at least one element in list")
	}
	var boolExpr spansql.BoolExpr
	for i, element := range list.ListExpr.Elements {
		elementExpr, err := exprToSpannerSQL(element)
		if err != nil {
			return nil, err
		}
		eq := spansql.ComparisonOp{
			LHS: needle,
			Op:  spansql.Eq,
			RHS: elementExpr,
		}
		if i == 0 {
			boolExpr = eq
		} else {
			boolExpr = spansql.LogicalOp{
				LHS: boolExpr,
				Op:  spansql.Or,
				RHS: eq,
			}
		}
	}
	parenExpr := spansql.Paren{Expr: boolExpr}
	return parenExpr, nil
}
