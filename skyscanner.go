package skyscanner

import "errors"

// Skyscanner implements the FlightFeeProvider interface
type Skyscanner struct {
	apiKey string
}

func NewSkyscannerProvider(apiKey string) *Skyscanner {
	return &Skyscanner{apiKey: apiKey}
}

func (p *Skyscanner) GetFlightFees(from, to string) (map[string]interface{}, error) {
	if from == "" || to == "" {
		return nil, errors.New("invalid input")
	}
	return map[string]interface{}{
		"provider": "Skyscanner",
		"fee":      200.0,
	}, nil
}
