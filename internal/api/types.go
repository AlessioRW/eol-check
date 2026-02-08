package api

type ClientInterface interface {
	Query(product string) ([]byte, error)
}

type Version struct {
	VersionNum   string
	IsLTS        bool
	IsMaintained bool
	IsEol        bool
	EolDate      string
}

type EolCheck struct {
	CheckId   string
	Product   string
	Current   *Version
	NextAlive *Version
	Latest    *Version
}

// API response for /v1/products/:product

type ProductsResponse struct {
	Result ProductsResponseResult `json:"result"`
}

type ProductsResponseResult struct {
	Releases []ProductsResponsRelease `json:"releases"`
}

type ProductsResponsRelease struct {
	Name         string `json:"name"`
	IsLts        bool   `json:"isLts"`
	IsEol        bool   `json:"isEol"`
	EolFrom      string `json:"eolFrom"`
	IsMaintained bool   `json:"isMaintained"`
}

func (r ProductsResponsRelease) toVersionType() *Version {

	if r.EolFrom == "" {
		r.EolFrom = "N/A"
	}

	return &Version{
		VersionNum:   r.Name,
		IsLTS:        r.IsLts,
		IsMaintained: r.IsMaintained,
		IsEol:        r.IsEol,
		EolDate:      r.EolFrom,
	}
}
