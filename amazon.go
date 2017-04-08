package main

import (
	"github.com/sclevine/agouti"
	"log"
)

func main() {
	driver := agouti.PhantomJS()
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page: %v", err)
	}
	if err := page.Navigate("https://www.amazon.co.jp"); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	if err := page.FindByID("twotabsearchtextbox").Fill("golang"); err != nil {
		log.Fatalf("Failed to search: %v", err)
	}
	if err := page.FindByID("twotabsearchtextbox").Submit(); err != nil {
		log.Fatalf("Failed to submit: %v", err)
	}
	page.Screenshot("golang.png")
	log.Print("Successed.")
}
