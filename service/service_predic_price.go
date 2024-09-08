package service

import (
	"context"
	"fmt"
	"go-lambda-sagemaker/gateway"
	"go-lambda-sagemaker/types"
	"log"
	"strconv"
	"sync"
	"time"
)

var ingredientsModelo = map[string]string{
	"TOMATE":       "tomate",
	"AJO":          "ajo",
	"ARROZ":        "rice",
	"PEPINO":       "pepino",
	"CALAMAR":      "squid",
	"CAMARON":      "shrimp",
	"CEBOLLA":      "cebolla",
	"PATATAS":      "patatas",
	"ACEITE_OLIVA": "aceite_oliva",
	"PAN":          "bread",
	"PIMIENTO":     "pimiento",
	"CARNE":        "beef",
	"HUEVO":        "huevo",
	"JAMON":        "jamon",
}

var quantityModelo = map[string]int{
	"TOMATE":       1,
	"AJO":          1,
	"ARROZ":        1,
	"PEPINO":       1,
	"CALAMAR":      1,
	"CAMARON":      1,
	"CEBOLLA":      1,
	"PATATAS":      1,
	"ACEITE_OLIVA": 1,
	"PAN":          1,
	"PIMIENTO":     1,
	"CARNE":        1,
	"HUEVO":        1,
	"JAMON":        1,
	"SAL":          1,
}

var measurementModelo = map[string]string{
	"TOMATE":       "kg",
	"AJO":          "kg",
	"ARROZ":        "kg",
	"PEPINO":       "kg",
	"CALAMAR":      "kg",
	"CAMARON":      "kg",
	"CEBOLLA":      "kg",
	"PATATAS":      "kg",
	"ACEITE_OLIVA": "L",
	"PAN":          "UNIDAD",
	"PIMIENTO":     "kg",
	"CARNE":        "kg",
	"HUEVO":        "kg",
	"JAMON":        "kg",
	"SAL":          "kg",
}

var hasModelol = map[string]bool{
	"TOMATE":       true,
	"AJO":          true,
	"ARROZ":        true,
	"PEPINO":       true,
	"CALAMAR":      true,
	"CAMARON":      true,
	"CEBOLLA":      true,
	"PATATAS":      true,
	"ACEITE_OLIVA": true,
	"PAN":          true,
	"PIMIENTO":     true,
	"CARNE":        true,
	"HUEVO":        true,
	"JAMON":        true,
	"SAL":          false,
}

func ProcessItems(ctx context.Context, request types.RequestBody) (types.ApiResponse, error) {
	var itemsFinal []types.ItemProduct
	//url := "https://9n2ncmueqe.execute-api.us-east-1.amazonaws.com/Prod/webscraping"
	//url := "https://7c75-95-18-16-108.ngrok-free.app/webscraping"
	// items, err := gateway.Post(url, make([]byte, 0))
	// if err != nil {
	// 	return types.ApiResponse{}, err
	// }

	currentTime := time.Now()
	// Format the date to "YYYY-MM-DD"
	formattedDate := currentTime.Format("2006-01-02")
	fmt.Print(formattedDate)
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]types.Result, 0)

	newItemSal := types.ItemProduct{
		Name:        "SAL",
		Price:       "0.64",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemAjo := types.ItemProduct{
		Name:        "AJO",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemArroz := types.ItemProduct{
		Name:        "ARROZ",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemCebolla := types.ItemProduct{
		Name:        "CEBOLLA",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemCarne := types.ItemProduct{
		Name:        "CARNE",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemHuevo := types.ItemProduct{
		Name:        "HUEVO",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemPan := types.ItemProduct{
		Name:        "PAN",
		Price:       "",
		Description: "1",
		UnitMeasure: "UNIDAD",
	}

	newItemPatata := types.ItemProduct{
		Name:        "PATATAS",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemJamon := types.ItemProduct{
		Name:        "JAMON",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemPepino := types.ItemProduct{
		Name:        "PEPINO",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemPimiento := types.ItemProduct{
		Name:        "PIMIENTO",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemTomate := types.ItemProduct{
		Name:        "TOMATE",
		Price:       "",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemAceite := types.ItemProduct{
		Name:        "ACEITE_OLIVA",
		Price:       "730.0",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemCalamar := types.ItemProduct{
		Name:        "CALAMAR",
		Price:       "22.90",
		Description: "1",
		UnitMeasure: "kg",
	}

	newItemCamaron := types.ItemProduct{
		Name:        "CAMARON",
		Price:       "11.99",
		Description: "1",
		UnitMeasure: "kg",
	}

	// Append items to the slice
	itemsFinal = append(itemsFinal, newItemSal, newItemAjo, newItemArroz, newItemCebolla, newItemCarne,
		newItemHuevo, newItemPan, newItemPatata, newItemJamon, newItemPepino, newItemPimiento, newItemTomate,
		newItemAceite, newItemCalamar, newItemCamaron)

	for _, item := range itemsFinal {
		wg.Add(1)
		go func(item types.ItemProduct) {
			defer wg.Done()
			// Convert the result map to Result type
			var result types.Result

			if hasModelol[item.Name] {
				inputData := types.SageMakerInput{
					Ingredient:   ingredientsModelo[item.Name],
					DateForecast: request.Date,
					Data: map[string][]string{
						"DATE":  {formattedDate},
						"PRICE": {item.Price},
					},
				}
				fmt.Print("IRENE")
				fmt.Print(inputData)
				resultMap, err := gateway.InvokeSageMaker(ctx, inputData)
				if err != nil {
					fmt.Println("error sage maker")
					fmt.Println(item.Name)
					fmt.Println(err)
					log.Printf("failed to invoke SageMaker endpoint for %s: %v", item.Name, err)
					return
				}

				if priceForecast, ok := resultMap["price_forecasted"].(float64); ok {
					result = types.Result{
						PriceForecasted: priceForecast, Ingredient: item.Name,
						Quantity:    quantityModelo[item.Name],
						Measurement: measurementModelo[item.Name],
					}
				}
			}

			if !hasModelol[item.Name] {
				// Convert the string to a float64
				floatValue, err := strconv.ParseFloat(item.Price, 64)
				if err != nil {
					fmt.Println("Error converting string to float:", err)
					return
				}
				result = types.Result{
					PriceForecasted: floatValue, Ingredient: item.Name,
					Quantity:    quantityModelo[item.Name],
					Measurement: measurementModelo[item.Name],
				}
			}

			mu.Lock()
			results = append(results, result)
			mu.Unlock()

		}(item)
	}

	wg.Wait()

	// Return ApiResponse with results
	return types.ApiResponse{Results: results}, nil
}
