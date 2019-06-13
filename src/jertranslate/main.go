// Sample translate-quickstart translates "Hello, world!" into Russian.
package main

import (
//	"context"
//	"fmt"
	"log"
	"mydb"

//	"cloud.google.com/go/translate"
//	"golang.org/x/text/language"
)

func main() {
	mydb.SetDSN("jerusalem:jerusalem@tcp(127.0.0.1:3306)/jerusalem?charset=utf8")


	languages, err := mydb.GetLanguages()
	if err != nil {
		log.Fatalf("Failed to get languages: %v", err)
	}
	log.Print(languages)

	/*
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to translate.
	text := "Hello, world!"
	// Sets the target language.
	target, err := language.Parse("ru")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)


	 */
}
