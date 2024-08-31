package hello

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/hello", new(Hello))
}

type Hello struct{}

func (h *Hello) HelloWorld() string {
	return "Hello World!"
}
