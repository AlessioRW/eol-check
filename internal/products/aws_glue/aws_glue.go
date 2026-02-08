package aws_glue

import (
	"eol-checker/internal/products"
	aws_glue_methods "eol-checker/internal/products/aws_glue/methods"
)

type Product struct{}

func (g Product) GetMethods() map[string]products.Method {
	return map[string]products.Method{
		"cli": aws_glue_methods.CLI{},
		"sdk": aws_glue_methods.SDK{},
	}
}

func (p Product) GetEndpoint() string {
	return "amazon-glue"
}
