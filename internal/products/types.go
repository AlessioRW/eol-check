package products

type Product interface {
	GetMethods() map[string]Method
	GetEndpoint() string
}

type Method interface {
	Run(id string, args []any) (string, error)
}
