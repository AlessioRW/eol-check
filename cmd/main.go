package main

import (
	"eol-checker/internal/api"
	"eol-checker/internal/config"
	"eol-checker/internal/output"
	"eol-checker/internal/products/golang"
	"fmt"
	"log/slog"
)

type Checker interface {
	IsEol(c config.ProductConfig) (*api.EolCheck, error)
}

func getChecker(product string) Checker {
	switch product {
	case "golang":
		return golang.GolangChecker{}
	default:
		return nil
	}
}

func main() {
	config, err := config.ParseConfig()
	if err != nil {
		return
	}

	eolResults := []*api.EolCheck{}

	for _, product := range config.Config {
		c := getChecker(product.Product)
		if c == nil {
			slog.Error(fmt.Sprintf("error in check: %v, product %v not recognised", product.Id, product.Product))
			continue
		}
		eolData, err := c.IsEol(product)
		if err != nil {
			return
		}

		eolResults = append(eolResults, eolData)
	}

	err = output.WriteOut(eolResults)
	if err != nil {
		return
	}
}
