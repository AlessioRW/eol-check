package grafana

import (
	"eol-checker/internal/products"
	grafana_methods "eol-checker/internal/products/grafana/methods"
)

type Product struct {
}

func (g Product) GetMethods() map[string]products.Method {
	return map[string]products.Method{
		"aws_sdk": grafana_methods.AWSSDK{},
	}
}

func (p Product) GetEndpoint() string {
	return "grafana"
}
