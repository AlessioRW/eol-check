package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

const API_URL = "https://endoflife.date/api/v1"

func DoEolCheck(product string, version string) (*EolCheck, error) {

	res, err := http.Get(fmt.Sprintf("%v/products/%v", API_URL, product))
	if err != nil {
		slog.Error("error getting eol info", "error", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("error reading response body", "error", err)
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		slog.Error("non-200 status code", "error", err, "body", string(body))
		return nil, err
	}

	eolData := &ProductsResponse{}
	err = json.Unmarshal(body, eolData)
	if err != nil {
		slog.Error("error unmarshalling api response", "error", err)
	}

	EolInfo := &EolCheck{}

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
