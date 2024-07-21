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

	minioStorage, err := storage.NewMinioStorage(ctx, "localhost:9000", "minioadmin", "minioadmin", "mybucket", false)
	if err != nil {
		panic(err)
	}
	defer minioStorage.Close(ctx)

	// Example file upload
	file, err := os.Open("./image.jpg")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	if err := minioStorage.Put(ctx, file, "image.jpg"); err != nil {
		log.Fatalf("Failed to upload file to MinIO: %v", err)
	}
	fmt.Println("File uploaded to MinIO successfully")

	// Example file link generation
	link, err := minioStorage.GetLink(ctx, "image.jpg", 1*time.Minute)
	if err != nil {
		log.Fatalf("Failed to get file link: %v", err)
	}
	fmt.Printf("Download link: %s\n", link)

	// Example file delete
	if err := minioStorage.Destroy(ctx, "image.jpg"); err != nil {
		log.Fatalf("Failed to delete file: %v", err)
	}
	fmt.Println("File deleted successfully")
}
