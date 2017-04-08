package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	// START OMIT
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	log.Print("page, err := driver.NewPage()") // OMIT
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
	}
	log.Print("page.Navigate(https://www.amazon.co.jp)") // OMIT
	if err := page.Navigate("https://www.amazon.co.jp"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	log.Print("page.FindByID(twotabsearchtextbox).Fill(golang)") // OMIT
	if err := page.FindByID("twotabsearchtextbox").Fill("golang"); err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	log.Print("page.FindByID(twotabsearchtextbox).Submit()") // OMIT
	if err := page.FindByID("twotabsearchtextbox").Submit(); err != nil {
		log.Fatalf("Failed to submit: %v", err)
	}
	log.Print("page.Screenshot(golang.png)") // OMIT
	if err := page.Screenshot("golang.png"); err != nil {
		log.Fatalf("Failed to screenshot: %v", err)
	}
	log.Print("Successed.") // OMIT
	// END OMIT
}
