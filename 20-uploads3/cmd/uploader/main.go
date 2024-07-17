package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("your-access-key-id", "your-secret-key", ""),
	})
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-example-bucket-lschenkel"
	wg = sync.WaitGroup{}
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
	uploadControl := make(chan struct{}, 100)
	fileUploadError := make(chan string, 10)

	go func() {
		for {
			select {
			case filename := <-fileUploadError:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(filename, uploadControl, fileUploadError)
			}
		}
	}()

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}

		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, fileUploadError)
	}

	wg.Wait()
}

func uploadFile(filename string, uploadControl <-chan struct{}, fileUploadError chan<- string) {
	defer wg.Done()
	fullFileName := fmt.Sprintf("./tmp/%s", filename)
	fmt.Printf("Uploading file %s to bucket %s\n", fullFileName, s3Bucket)
	f, err := os.Open(fullFileName)
	if err != nil {
		fmt.Printf("Error opening file %s: %s\n", fullFileName, err)
		<-uploadControl
		fileUploadError <- filename
		return
	}

	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", fullFileName)
		<-uploadControl
		fileUploadError <- filename
		return
	}

	fmt.Printf("File %s uploaded successfully\n", fullFileName)
	<-uploadControl
}
