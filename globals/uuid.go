package globals

import "github.com/satori/go.uuid"

var Nil = uuid.Nil

func UUID() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
