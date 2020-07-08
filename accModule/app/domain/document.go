package domain

import guid "github.com/satori/go.uuid"
type document struct {
	uuid guid
}

func NewDocument() *document {
	return &document{}
}

func Document(*document) string {
	return "111"
}