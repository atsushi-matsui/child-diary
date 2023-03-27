// Sample vision-quickstart uses the Google Cloud Vision API to label an image.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/atsushi-mastui/child-diary/domain"
	myGoogle "github.com/atsushi-mastui/child-diary/google"
	"github.com/joho/godotenv"
)

func main() {
	// read env file
	loadEnv()

	// get arg
	fileName := getArg()
	if fileName == "" {
		log.Fatalf("Not found arg")
	}

	client, err := myGoogle.NewClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	imageFile, err := domain.NewFile(string(os.Getenv("IMAGE_DIR")), fileName, "png")
	if err != nil {
		log.Fatalf("Opening file: %v", err)
	}
	defer imageFile.Close()

	text, err := client.DetectDocumentText(imageFile.File)
	if err != nil {
		log.Fatalf("Failed to detect text: %v", err)
	}
	fmt.Println("Texts:")
	fmt.Println(text)

	blog, err := domain.NewBlog(fileName, fileName)
	if err != nil {
		log.Fatalf("Failed to new blog: %v", err)
	}
	content := blog.ToText() + text

	diaryFile, err := domain.NewFile(string(os.Getenv("DIARY_DIR")), fileName, "md")
	if err != nil {
		log.Fatalf("Opening file: %v", err)
	}
	defer diaryFile.Close()
	diaryFile.Write(content)
}

func getArg() string{
	arg := flag.String("argument", "", "argument from Makefile")
	flag.Parse()

	return *arg
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
}
