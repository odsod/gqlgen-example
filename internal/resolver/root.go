package resolver

import (
	"github.com/odsod/gqlgen-example/internal/graph"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
)

type Root struct {
	Storage          *storage.InMemory
	Logger           *zap.Logger
	MutationResolver *Mutation
	QueryResolver    *Query
	TodoResolver     *Todo
}

func (r *Root) Mutation() graph.MutationResolver {
	r.Logger.Debug("root: mutation")
	return r.MutationResolver
}

func (r *Root) Query() graph.QueryResolver {
	r.Logger.Debug("root: query")
	return r.QueryResolver
}

func (r *Root) Todo() graph.TodoResolver {
	r.Logger.Debug("root: todo")
	return r.TodoResolver
}
