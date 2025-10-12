package config

import (
	"errors"
	"os"
)

var (
	AWS_ACCESS_KEY_ID     string
	AWS_SECRET_ACCESS_KEY string
	AWS_REGION            string
	S3_BUCKET             string
	MONGO_COLLECTION      string
	MONGODB_NAME          string
	MONGO_URI             string
	LOG_LEVEL             string
	DB_POSTGRES_URI       string
	REDIS_ADDR            string
	REDIS_PASSWORD        string
	REDIS_DB              int
)

func Load() error {
	AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
	if AWS_ACCESS_KEY_ID == "" {
		return errors.New("missing AWS_ACCESS_KEY_ID environment variable")
	}

	AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
	if AWS_SECRET_ACCESS_KEY == "" {
		return errors.New("missing AWS_SECRET_ACCESS_KEY environment variable")
	}

	AWS_REGION = os.Getenv("AWS_REGION")
	if AWS_REGION == "" {
		return errors.New("missing AWS_REGION environment variable")
	}

	S3_BUCKET = os.Getenv("S3_BUCKET")
	if S3_BUCKET == "" {
		return errors.New("missing S3_BUCKET environment variable")
	}

	return nil
}
