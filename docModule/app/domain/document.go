package domain

type document struct {
	uuid string
}

func NewDocument(z string) *document {
	return &document{z}
}

func Document(d *document) string {
	return d.uuid
}
