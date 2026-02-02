package golang

import (
	"eol-checker/internal/products"
	golang_methods "eol-checker/internal/products/golang/methods"
)

type Product struct {
}

func (g Product) GetMethods() map[string]products.Method {
	return map[string]products.Method{
		"file":    golang_methods.File{},
		"command": golang_methods.Command{},
	}
}

func (p Product) GetEndpoint() string {
	return "golang"
}
