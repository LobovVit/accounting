package domain

import uuid "github.com/google/uuid"

type document struct {
	uuid uuid.UUID //

}

func NewDocument() *document {
	return &document{}
}

func Document(*document) string {
	return "111"
}
