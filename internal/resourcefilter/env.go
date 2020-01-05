package resourcefilter

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

type Resource struct {
	resourceName string
	resource     proto.Message
	env          cel.Env
}

func NewResource(resourceName string, resource proto.Message) (*Resource, error) {
	env, err := cel.NewEnv(
		cel.Types(resource),
		cel.Declarations(
			decls.NewIdent(resourceName, decls.NewObjectType(proto.MessageName(resource)), nil),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("new entity filter: %w", err)
	}
	e := &Resource{
		resourceName: resourceName,
		resource:     resource,
		env:          env,
	}
	return e, nil
}

func (r *Resource) CompileFilter(filter string) (*Filter, error) {
	ast, issues := r.env.Parse(filter)
	if issues != nil {
		return nil, fmt.Errorf("compile filter: parse: %w", issues.Err())
	}
	ast, issues = r.env.Check(ast)
	if issues != nil {
		return nil, fmt.Errorf("compile filter: check: %w", issues.Err())
	}
	if ast.ResultType().GetPrimitive() != expr.Type_BOOL {
		return nil, fmt.Errorf("compile resource filter: non-bool result type: %v", ast.ResultType().GetPrimitive())
	}
	program, err := r.env.Program(ast)
	if err != nil {
		return nil, fmt.Errorf("compile filter: %w", issues.Err())
	}
	f := &Filter{
		resourceName: r.resourceName,
		env:          nil,
		ast:          ast,
		program:      program,
	}
	return f, nil
}
