package resolver

import (
	"log"

	"github.com/odsod/gqlgen-getting-started/internal/graph"
	"github.com/odsod/gqlgen-getting-started/internal/model"
)

type Root struct {
	todos []*model.Todo
}

func (r *Root) Mutation() graph.MutationResolver {
	log.Printf("root: mutation")
	return &mutationResolver{r}
}

func (r *Root) Query() graph.QueryResolver {
	log.Printf("root: query")
	return &query{r}
}

func (r *Root) Todo() graph.TodoResolver {
	log.Printf("root: todo")
	return &todo{r}
}
