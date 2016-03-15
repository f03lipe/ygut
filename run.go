package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const S3_DATA_PATH = "data.json"

// See github.com/aws/aws-sdk-go
func S3UploadBody(path string, body []byte) error {
	ss := session.New()
	conf := aws.Config{
		Region: aws.String("us-east-1"),
		//Credentials: ec2rolecreds.NewCredentials(nil, 5*time.Minute),
	}
	svc := s3.New(ss, &conf)

	params := &s3.PutObjectInput{
		Bucket: aws.String("ypncks"),
		Key:    aws.String(path),
		Body:   bytes.NewReader(body),
		ACL:    aws.String("public-read"),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalf("Failed to upload object to %s.\n", path)
	}
	fmt.Println(resp)
	return nil
}

func main() {
	fmt.Println("Starting run.go")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		meals := GetMeals()
		b, err := json.MarshalIndent(meals, "", "    ")
		if err != nil {
			log.Fatal("Failed to marshal")
		}
		w.Write(b)
	})

	log.Printf("Listening on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
