package utils

import (
	"github.com/google/uuid"
	"net/http"
)

func ParseUUID(str string) uuid.UUID {
	id, err := uuid.Parse(str)
	if err != nil {
		CatchError(NewHttpError(http.StatusBadRequest, "Invalid UUID", nil))
	}

	return id
}
