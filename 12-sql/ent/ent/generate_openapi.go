//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entoas.NewExtension()

	if err != nil {
		log.Fatalf("failed creating extension: %v", err)
	}

	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))

	if err != nil {
		log.Fatalf("failed generating ent: %v", err)
	}
}
