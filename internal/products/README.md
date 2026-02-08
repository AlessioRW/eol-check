# Package - Products

Products are designed in a way that every check for all products work universally, this is to isolate the version checking logic which is expected to be expanded upon from the rest of the code.

## Design

This implementation relies on a `Product` and `Method` interface per product. 

#### **Product** 
Represents a product availble to query for end-of-life data on endoflife.date. 

`GetMethods()` - Gets the available methods of getting the verion of a product

```go
type Product interface {
	GetMethods() map[string]Method 
}
```


#### **Method**

Represents one implementation of getting a products used version. 

`Run` - Logic to get the product version
 - ID - Id of the check, specified in the config file.
 - Args - Array of arguments to method. Type casted in function.
```go
type Method interface {
	Run(id string, args []any) (string, error)
}
```

