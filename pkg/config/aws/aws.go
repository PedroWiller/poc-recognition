package aws

import (
	"context"
	"log"

	"poc-recognition/pkg/config/env"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

type RekognitionService struct {
	client *rekognition.Client
	ctx    context.Context
}

func LoadConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(env.AWS_REGION))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return cfg
}

func NewSession() (*RekognitionService, error) {
	cfg := LoadConfig()
	return &RekognitionService{
		client: rekognition.NewFromConfig(cfg),
		ctx:    context.Background(),
	}, nil
}
