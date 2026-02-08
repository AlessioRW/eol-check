package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type Client struct {
	URL string
}

// TODO: api calls for a product can be cached in case called again
func GetData(apiClient Client, id string, product string, productEndpoint string, version string) (EolCheck, error) {
	EolInfo := EolCheck{
		CheckId: id,
		Product: product,
	}

	eolData := &ProductsResponse{}

	body, err := apiClient.Query(productEndpoint)
	if err != nil {
		return EolInfo, err
	}

	err = json.Unmarshal(body, eolData)
	if err != nil {
		slog.Error("error unmarshalling api response", "error", err)
		return EolInfo, err
	}

	for i, versionData := range eolData.Result.Releases {
		if i == 0 {
			EolInfo.Latest = versionData.toVersionType()
		}
		if versionData.Name == string(version) {
			EolInfo.Current = versionData.toVersionType()
		} else if !versionData.IsEol {
			EolInfo.NextAlive = versionData.toVersionType()
		}
	}

	return EolInfo, nil
}

func (caller *Client) Query(productEndpoint string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%v/products/%v", caller.URL, productEndpoint))
	if err != nil {
		slog.Error("error calling eol API info", "error", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("error reading response body from eol API", "error", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		slog.Error("non-200 status code calling eol API", "error", err, "body", string(body))
		return nil, err
	}

	return body, nil
}
