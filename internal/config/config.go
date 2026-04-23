package config

import "os"

type Config struct {
	AppPort        string
	DatabaseURL    string
	JWTSecret      string
	AWSRegion      string
	S3Bucket       string
	S3Endpoint     string
	S3AccessKey    string
	S3SecretKey    string
	UsePathStyleS3 bool
}

func Load() Config {
	port := getEnv("APP_PORT", "8080")
	if port[0] != ':' {
		port = ":" + port
	}

	return Config{
		AppPort:        port,
		DatabaseURL:    getEnv("DATABASE_URL", "postgres://postgres:postgres@postgres:5432/saas_db?sslmode=disable"),
		JWTSecret:      getEnv("JWT_SECRET", "change-me"),
		AWSRegion:      getEnv("AWS_REGION", "us-east-1"),
		S3Bucket:       getEnv("S3_BUCKET", "saas-demo-bucket"),
		S3Endpoint:     os.Getenv("S3_ENDPOINT"),
		S3AccessKey:    os.Getenv("S3_ACCESS_KEY"),
		S3SecretKey:    os.Getenv("S3_SECRET_KEY"),
		UsePathStyleS3: getEnv("S3_USE_PATH_STYLE", "true") == "true",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
