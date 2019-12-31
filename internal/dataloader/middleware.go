package dataloader

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/odsod/gqlgen-example/internal/model"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
)

type (
	userLoaderKey struct{}
)

func UserLoaderFromContext(ctx context.Context) (*UserLoader, bool) {
	userLoader, ok := ctx.Value(userLoaderKey{}).(*UserLoader)
	return userLoader, ok
}

type Middleware struct {
	Storage *storage.InMemory
	Logger  *zap.Logger
}

func makeIndexMap(elements []string) map[string]int {
	result := make(map[string]int)
	for i, element := range elements {
		result[element] = i
	}
	return result
}

func (m *Middleware) FetchUsers(ctx context.Context, ids []string) ([]*model.User, []error) {
	m.Logger.Debug("fetch users", zap.Strings("ids", ids))
	users := make([]*model.User, len(ids))
	errs := make([]error, len(ids))
	foundUsers, missingIDs, err := m.Storage.BatchGetUsers(ctx, ids)
	if err != nil {
		for i := range errs {
			errs[i] = err
		}
		return nil, errs
	}
	idToIndexMap := makeIndexMap(ids)
	for _, user := range foundUsers {
		users[idToIndexMap[user.ID]] = user
	}
	for _, missingID := range missingIDs {
		errs[idToIndexMap[missingID]] = fmt.Errorf("not found: %s", missingID)
	}
	return users, errs
}

func (m *Middleware) ApplyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), userLoaderKey{}, NewUserLoader(UserLoaderConfig{
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(ids []string) (users []*model.User, errors []error) {
				return m.FetchUsers(r.Context(), ids)
			},
		})))
		next.ServeHTTP(w, r)
	})
}
