package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/odsod/gqlgen-example/internal/gen/dataloader"
	userv1beta1 "github.com/odsod/gqlgen-example/internal/gen/proto/go/odsod/user/v1beta1"
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

func (m *Dataloader) FetchUsers(ctx context.Context, names []string) ([]*userv1beta1.User, []error) {
	m.Logger.Debug("fetch users", zap.Strings("ids", names))
	users := make([]*userv1beta1.User, len(names))
	errs := make([]error, len(names))
	response, err := m.UserServiceClient.BatchGetUsers(ctx, &userv1beta1.BatchGetUsersRequest{
		Names: names,
	})
	if err != nil {
		return nil, []error{err}
	}
	nameToIndexMap := makeIndexMap(names)
	for _, user := range response.FoundUsers {
		users[nameToIndexMap[user.Name]] = user
	}
	for _, missingName := range response.MissingNames {
		errs[nameToIndexMap[missingName]] = fmt.Errorf("not found: %s", missingName)
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
