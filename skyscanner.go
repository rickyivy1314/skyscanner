package skyscanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Skyscanner implements the FlightFeeProvider interface
type SkyscannerProvider struct {
	apiKey string
}

func NewSkyscannerProvider(apiKey string) *SkyscannerProvider {
	return &SkyscannerProvider{apiKey: apiKey}
}

// GetFlightFees fetches flight fees from the Skyscanner API
func (p *SkyscannerProvider) GetFlightFees(from, to string) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf("https://partners.api.skyscanner.net/apiservices/browseroutes/v1.0/US/USD/en-US/%s/%s/2023-12-25?apiKey=%s", from, to, p.apiKey)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return result, nil
}
