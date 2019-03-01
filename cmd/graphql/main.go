package main

import (
	"github.com/issho-ni/issho/internal/app/graphql"
	_ "github.com/issho-ni/issho/internal/pkg/issho"
)

func main() {
	graphql.StartServer()
}
