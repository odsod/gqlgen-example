package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/odsod/gqlgen-example/internal/gen/dataloader"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
	"github.com/odsod/gqlgen-example/internal/storage"
	"go.uber.org/zap"
)

type (
	userLoaderKey struct{}
)

func UserLoaderFromContext(ctx context.Context) (*dataloader.UserLoader, bool) {
	userLoader, ok := ctx.Value(userLoaderKey{}).(*dataloader.UserLoader)
	return userLoader, ok
}

type Dataloader struct {
	Storage           *storage.InMemory
	UserServiceClient userv1beta1.UserServiceClient
	Logger            *zap.Logger
}

func makeIndexMap(elements []string) map[string]int {
	result := make(map[string]int)
	for i, element := range elements {
		result[element] = i
	}
	return result
}

func (m *Dataloader) FetchUsers(ctx context.Context, ids []string) ([]*userv1beta1.User, []error) {
	m.Logger.Debug("fetch users", zap.Strings("ids", ids))
	users := make([]*userv1beta1.User, len(ids))
	errs := make([]error, len(ids))
	response, err := m.UserServiceClient.BatchGetUsers(ctx, &userv1beta1.BatchGetUsersRequest{
		Ids: ids,
	})
	if err != nil {
		for i := range errs {
			errs[i] = err
		}
		return nil, errs
	}
	idToIndexMap := makeIndexMap(ids)
	for _, user := range response.FoundUsers {
		users[idToIndexMap[user.Id]] = user
	}
	for _, missingID := range response.MissingIds {
		errs[idToIndexMap[missingID]] = fmt.Errorf("not found: %s", missingID)
	}
	return users, errs
}

func (m *Dataloader) ApplyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), userLoaderKey{}, dataloader.NewUserLoader(dataloader.UserLoaderConfig{
			Wait:     1 * time.Millisecond,
			MaxBatch: 100,
			Fetch: func(ids []string) (users []*userv1beta1.User, errors []error) {
				return m.FetchUsers(r.Context(), ids)
			},
		})))
		next.ServeHTTP(w, r)
	})
}
