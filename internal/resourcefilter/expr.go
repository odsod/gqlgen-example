package resourcefilter

import (
	"github.com/google/cel-go/common/operators"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

func Or(lhs *expr.Expr, rhs *expr.Expr) *expr.Expr {
	return &expr.Expr{
		ExprKind: &expr.Expr_CallExpr{
			CallExpr: &expr.Expr_Call{
				Function: operators.LogicalOr,
				Args:     []*expr.Expr{lhs, rhs},
			},
		},
	}
}

func Ident(name string) *expr.Expr {
	return &expr.Expr{
		ExprKind: &expr.Expr_IdentExpr{
			IdentExpr: &expr.Expr_Ident{
				Name: name,
			},
		},
	}
}

func Select(operand *expr.Expr, field string) *expr.Expr {
	return &expr.Expr{
		ExprKind: &expr.Expr_SelectExpr{
			SelectExpr: &expr.Expr_Select{
				Operand: operand,
				Field:   field,
			},
		},
	}
}

func In(lhs *expr.Expr, listExpr *expr.Expr_ListExpr) *expr.Expr {
	return &expr.Expr{
		ExprKind: &expr.Expr_CallExpr{
			CallExpr: &expr.Expr_Call{
				Function: operators.In,
				Args: []*expr.Expr{
					lhs,
					{ExprKind: listExpr},
				},
			},
		},
	}
}

func Strings(ss ...string) *expr.Expr_ListExpr {
	listExpr := &expr.Expr_ListExpr{
		ListExpr: &expr.Expr_CreateList{},
	}
	for _, s := range ss {
		listExpr.ListExpr.Elements = append(listExpr.ListExpr.Elements, String(s))
	}
	return listExpr
}

func String(s string) *expr.Expr {
	return &expr.Expr{
		ExprKind: &expr.Expr_ConstExpr{
			ConstExpr: &expr.Constant{
				ConstantKind: &expr.Constant_StringValue{
					StringValue: s,
				},
			},
		},
	}
}
