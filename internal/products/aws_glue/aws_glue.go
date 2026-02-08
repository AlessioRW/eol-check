package aws_glue

import (
	"eol-checker/internal/products"
	aws_glue_methods "eol-checker/internal/products/aws_glue/methods"
)

type Product struct{}

func (g Product) GetMethods() map[string]products.Method {
	return map[string]products.Method{
		"cle": aws_glue_methods.CLI{},
	}
}

func (p Product) GetEndpoint() string {
	return "golang"
}
