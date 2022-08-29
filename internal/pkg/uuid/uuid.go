package uuid

import "github.com/google/uuid"

func New() string {
	id, err := uuid.NewUUID()
	if err != nil {
		id = uuid.New()
	}

	return id.String()
}
