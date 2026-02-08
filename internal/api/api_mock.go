package api

type MockAPIClient struct {
}

func (api *MockAPIClient) Query(id string, product string, version string) (*EolCheck, error) {
	return nil, nil
}
