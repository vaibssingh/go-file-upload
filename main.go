package main

import (
	"context"
	"log"
	"fmt"
	"os"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

//GetEnvWithKey : get env value
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
}

var awsS3Client *s3.Client

func main() {
	LoadEnv()
	configS3()

	http.HandleFunc("/upload", handlerUpload)
	
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func configS3() {
	AccessKeyID := GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey := GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	MyRegion := GetEnvWithKey("AWS_REGION")

	creds := credentials.NewStaticCredentialsProvider(AccessKeyID, SecretAccessKey, "") 

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(MyRegion))

	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	awsS3Client = s3.NewFromConfig(cfg)
}

func showError(w http.ResponseWriter, r *http.Request, status int, message string) {
	http.Error(w, message, status)
}