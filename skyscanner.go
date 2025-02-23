package skyscanner

import "errors"

// Skyscanner implements the FlightFeeProvider interface
type Skyscanner struct {
	apiKey string
}

func NewSkyscannerProvider(apiKey string) *Skyscanner {
	// 使用 apiKey 初始化 Skyscanner
	return &Skyscanner{apiKey: apiKey}
}

func (p *Skyscanner) GetFlightFees(from, to string) (map[string]interface{}, error) {
	// 实现获取费用的逻辑
	if from == "" || to == "" {
		return nil, errors.New("invalid input")
	}
	return map[string]interface{}{
		"provider": "Skyscanner",
		"fee":      200.0, // 示例返回值
	}, nil
}
