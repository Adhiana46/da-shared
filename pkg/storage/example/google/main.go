package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/adhiana46/da-shared/pkg/storage"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancel()

	gcs, err := storage.NewGoogleStorage(ctx, "bridgtl-slt-d-sales-bucket", storage.GoogleStorageCredential{
		ProjectId:    "your-project-id",
		PrivateKeyId: "your-private-key-id",
		PrivateKey:   "your-private-key",
		ClientId:     "your-client-id",
		ClientEmail:  "your-client-email",
	})
	if err != nil {
		panic(err)
	}
	defer gcs.Close(ctx)

	// Example file upload
	file, err := os.Open("./image.jpg")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	if err := gcs.Put(ctx, file, "image.jpg"); err != nil {
		log.Fatalf("Failed to upload file to Google Cloud Storage: %v", err)
	}
	fmt.Println("File uploaded to Google Cloud Storage successfully")

	// Example file link generation
	link, err := gcs.GetLink(ctx, "image.jpg", 1*time.Minute)
	if err != nil {
		log.Fatalf("Failed to get file link: %v", err)
	}
	fmt.Printf("Download link: %s\n", link)

	// Example file delete
	if err := gcs.Destroy(ctx, "image.jpg"); err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	fmt.Println("File deleted successfully")

	fmt.Println("GCS", gcs)
}
