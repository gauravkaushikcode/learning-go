package main

import "fmt"

type magazine struct {
	publication
}

func (m magazine) String() string {
	return fmt.Sprintf("magazine name %s", m.name)
}

func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}
}
