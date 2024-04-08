package location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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

	var provinces *response.RegencyResponse
	err = json.Unmarshal(body, &provinces)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	results := []*response.LocationResponse{}
	for _, item := range provinces.Value {
		cleanedID := helper.CleanNumericString(item.ID)

		provinceID, err := strconv.Atoi(cleanedID)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int: %w", err)
		}

		results = append(results, &response.LocationResponse{
			ID:   provinceID,
			Name: item.Name,
		})
	}

	return results, nil
}
