package google

import (
	"context"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

type ImageScanner struct {
	client *vision.ImageAnnotatorClient
	context context.Context
}

func NewClient() (*ImageScanner, error) {
	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
    if err != nil {
            log.Fatalf("Failed to create client: %v", err)
    }

	return &ImageScanner{
		client: client,
		context: ctx,
	}, nil
}

func (imageScanner *ImageScanner) DetectDocumentText(file *os.File) (string, error) {
	image, err := vision.NewImageFromReader(file)
    if err != nil {
        log.Fatalf("Failed to create image: %v", err)
    }
	text, err := imageScanner.client.DetectDocumentText(imageScanner.context, image, nil)
	if err != nil {
        log.Fatalf("Failed to detect text: %v", err)
    }

	return text.Text, nil
}

func (imageScanner *ImageScanner) Close() {
	imageScanner.client.Close()
}
