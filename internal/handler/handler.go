package handler

import (
	"eol-checker/internal/api"
	"eol-checker/internal/config"
	"eol-checker/internal/output"
	"eol-checker/internal/products"
	"eol-checker/internal/products/golang"
)

func Run(config *config.EolConfig) error {
	eolResults := []*api.EolCheck{}
	for _, product := range config.Config {
		p := switchProduct(product.Product)

		checkFunc := p.GetMethods()[product.Method]
		version, err := checkFunc.Run(product.Id, []any{product.Path})
		if err != nil {
			return err
		}

		vData, err := api.GetEolData(product.Id, product.Product, version)
		if err != nil {
			return err
		}

		eolResults = append(eolResults, vData)

	}

	err := output.WriteOut(eolResults)
	if err != nil {
		return err
	}

	return nil
}

func switchProduct(product string) products.Product {
	switch product {
	case "golang":
		return golang.Product{}
	default:
		return nil
	}
}
