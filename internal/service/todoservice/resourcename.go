package todoservice

import (
	"fmt"
	"strings"
)

const TodosCollectionID = "todos"

type TodoResourceName struct {
	ID string
}

func (r *TodoResourceName) String() string {
	return fmt.Sprintf("%s/%s", TodosCollectionID, r.ID)
}

func (r *TodoResourceName) UnmarshalString(s string) error {
	parts := strings.Split(s, "/")
	if len(parts) != 2 {
		return fmt.Errorf("unmarshal todo resource name '%s': invalid URI", s)
	}
	if parts[0] != TodosCollectionID {
		return fmt.Errorf("unmarshal todo resource name '%s': invalid collection", s)
	}
	r.ID = parts[1]
	return nil
}
