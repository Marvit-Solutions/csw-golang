package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) Regency(req request.LocationRequest) ([]*response.LocationResponse, error) {
	reqUrl, err := http.NewRequest("GET", req.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := http.DefaultClient.Do(reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var regencies *response.RegencyResponse
	err = json.Unmarshal(body, &regencies)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	results := []*response.LocationResponse{}
	for _, item := range regencies.Value {
		results = append(results, &response.LocationResponse{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	return results, nil
}
