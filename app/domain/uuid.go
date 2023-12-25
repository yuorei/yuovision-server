package domain

import "github.com/google/uuid"

func NewUUID() string {
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return uuidObj.String()
}
