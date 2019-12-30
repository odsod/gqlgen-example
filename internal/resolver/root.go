package resolver

import (
	"log"

	"github.com/odsod/gqlgen-getting-started/internal/graph"
	"github.com/odsod/gqlgen-getting-started/internal/storage"
)

type Root struct {
	Storage storage.InMemory
}

func (r *Root) Mutation() graph.MutationResolver {
	log.Printf("root: mutation")
	return &mutationResolver{root: r}
}

func (r *Root) Query() graph.QueryResolver {
	log.Printf("root: query")
	return &query{root: r}
}

func (r *Root) Todo() graph.TodoResolver {
	log.Printf("root: todo")
	return &todo{root: r}
}
