package handler

import (
	"eol-checker/internal/api"
	"eol-checker/internal/config"
	"eol-checker/internal/output"
	"errors"
	"fmt"
	"log/slog"
)

const API_URL = "https://endoflife.date/api/v1"

func Run(config *config.EolConfig) error {

	apiClient := api.Client{
		URL: API_URL,
	}
	outputClient := output.Client{}

	eolResults := []api.EolCheck{}
	for _, product := range config.Config {
		p := switchProduct(product.Product)

		checkFunc, ok := p.GetMethods()[product.Method]
		if !ok {
			errMsg := fmt.Sprintf("method %v does not exist on prodcut %v", product.Method, product.Product)
			slog.Error(errMsg)
			return errors.New(errMsg)
		}

		version, err := checkFunc.Run(product.Id, product.Args)
		if err != nil {
			return err
		}

		vData, err := api.GetData(apiClient, product.Id, product.Product, p.GetEndpoint(), version)
		if err != nil {
			return err
		}

		eolResults = append(eolResults, vData)
	}

	err := output.Output(outputClient, eolResults)
	if err != nil {
		return err
	}

	return nil
}
