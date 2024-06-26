package main

import "fmt"

func newPublication(pubType string, name string, pages int, pub string) (iPublication, error) {
	if pubType == "newspaper" {
		return createNewsPaper(name, pages, pub), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pages, pub), nil
	}
	return nil, fmt.Errorf("no Publication found")
}
