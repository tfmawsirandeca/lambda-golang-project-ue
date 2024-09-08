package handler

import (
	"context"
	"encoding/json"
	"go-lambda-sagemaker/service"
	"go-lambda-sagemaker/types"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var bodyReq types.RequestBody
	// Parse the request body
	err := json.Unmarshal([]byte(request.Body), &bodyReq)
	if err != nil {
		log.Printf("error parsing request body: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `{"error": "Invalid request body"}`,
		}, err
	}

	// Call the ProcessItems function to get the response
	apiResponse, err := service.ProcessItems(ctx, bodyReq)
	if err != nil {
		log.Printf("error processing items: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Internal Server Error"}`,
		}, err
	}

	// Marshal the ApiResponse to JSON
	body, err := json.Marshal(apiResponse)
	if err != nil {
		log.Printf("error marshalling response: %v", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       `{"error": "Internal Server Error"}`,
		}, err
	}

	// Return a properly formatted API Gateway response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(body),
	}, nil
}
