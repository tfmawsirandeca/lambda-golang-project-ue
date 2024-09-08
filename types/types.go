package types

type InputItem struct {
	Name  string `json:"name"`
	Price string `json:"price"`
	Model string `json:"model"`
}

type SageMakerInput struct {
	Ingredient   string              `json:"ingredient"`
	DateForecast string              `json:"date_forescast"`
	Data         map[string][]string `json:"data"`
}

type SageMakerResponse struct {
	// Define the structure of the expected response
}

// Result represents each item in the results array.
type Result struct {
	PriceForecasted float64 `json:"price_forecasted"`
	Ingredient      string  `json:"ingredient"`
	Quantity        int     `json:"quantity"`
	Measurement     string  `json:"measurement"`
}

// ApiResponse represents the structure of the entire JSON response.
type ApiResponse struct {
	Results []Result `json:"results"`
}

type RequestBody struct {
	Date string `json:"date"`
}
