package handler

import (
	"eol-checker/internal/products"
	"eol-checker/internal/products/aws_glue"
	"eol-checker/internal/products/golang"
)

func switchProduct(product string) products.Product {
	switch product {
	case "golang":
		return golang.Product{}
	case "aws_glue":
		return aws_glue.Product{}
	default:
		return nil
	}
}
