package domain

import "github.com/google/uuid"

type transaction struct {
	TrUuid      uuid.UUID
	TrEventUuid uuid.UUID
}
