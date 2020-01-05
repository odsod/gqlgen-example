package resourcefilter

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/golang/protobuf/proto"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	resource, err := NewResource("user", &userv1beta1.User{})
	require.NoError(t, err)
	filter, err := resource.CompileFilter(`user.display_name in ["User 1", "User 2"] && !user.deleted`)
	require.NoError(t, err)
	require.True(t, filter.Test(&userv1beta1.User{
		DisplayName: "User 1",
	}))
	require.False(t, filter.Test(&userv1beta1.User{
		DisplayName: "User 1",
		Deleted:     true,
	}))
	spanSQL, err := filter.SpannerSQL()
	require.NoError(t, err)
	spew.Dump(spanSQL)
	require.Equal(t, `((display_name = "User 1" OR display_name = "User 2") AND NOT deleted)`, spanSQL.SQL())
}

func Test2(t *testing.T) {
	env, err := cel.NewEnv(
		cel.Types(&userv1beta1.User{}),
		cel.Declarations(
			decls.NewIdent("user", decls.NewObjectType(proto.MessageName(&userv1beta1.User{})), nil),
		),
	)
	require.NoError(t, err)
	ast, issues := env.Parse(`user.display_name in ["User 1", "User 2"]`)
	require.Nil(t, issues)
	ast, issues = env.Check(ast)
	require.Nil(t, issues)
	fmt.Println(proto.MarshalTextString(ast.Expr()))
}
