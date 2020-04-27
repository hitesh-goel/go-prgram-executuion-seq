package main

import (
	"fmt"

	"github.com/hitesh-goel/version"
)

func init() {
	fmt.Println("app/fetch-version.go ==> init()")
}

func fetchVersion() string {
	fmt.Println("app/fetch-version.go ==> fetchVersion()")
	return version.Version
}
