package resourcefilter

import (
	"fmt"

	"cloud.google.com/go/spanner/spansql"
	"github.com/golang/protobuf/proto"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/types"
)

type Filter struct {
	resourceName string
	env          cel.Env
	ast          cel.Ast
	program      cel.Program
}

func (r *Filter) Test(resource proto.Message) bool {
	out, _, _ := r.program.Eval(map[string]interface{}{r.resourceName: resource})
	return out == types.True
}

func (r *Filter) SpannerSQL() (spansql.BoolExpr, error) {
	spanExpr, err := exprToSpannerSQL(r.ast.Expr())
	if err != nil {
		return nil, fmt.Errorf("resource filter to spanner SQL: %w", err)
	}
	spanBoolExpr, ok := spanExpr.(spansql.BoolExpr)
	if !ok {
		return nil, fmt.Errorf("resource filter to spanner SQL: non-bool expression")
	}
	return spanBoolExpr, nil
}
