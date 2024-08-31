package main

import (
	"fmt"

	"go.k6.io/k6/js/modules"
)

type MyModule struct{}

func (m *MyModule) HelloWorld() {
	fmt.Println("Hello, k6 World!")
}

func init() {
	modules.Register("k6/x/mymodule", &MyModule{})
}
