package userservice

import (
	"fmt"
	"strings"
)

const UsersCollectionID = "users"

type UserResourceName struct {
	ID string
}

func (r *UserResourceName) String() string {
	return fmt.Sprintf("%s/%s", UsersCollectionID, r.ID)
}

func (r *UserResourceName) UnmarshalString(s string) error {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return fmt.Errorf("unmarshal user resource name '%s': invalid URI", s)
	}
	if parts[0] != UsersCollectionID {
		return fmt.Errorf("unmarshal user resource name '%s': invalid collection", s)
	}
	r.ID = parts[1]
	return nil
}
