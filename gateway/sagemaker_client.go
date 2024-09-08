package gateway

import (
	"context"
	"encoding/json"
	"go-lambda-sagemaker/types"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerruntime"
)

func InvokeSageMaker(ctx context.Context, input types.SageMakerInput) (map[string]interface{}, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}

	svc := sagemakerruntime.NewFromConfig(cfg)

	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	response, err := svc.InvokeEndpoint(ctx, &sagemakerruntime.InvokeEndpointInput{
		EndpointName: aws.String("price-prediction-endpoint-modal"),
		ContentType:  aws.String("application/json"),
		Body:         body,
	})

	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response.Body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
