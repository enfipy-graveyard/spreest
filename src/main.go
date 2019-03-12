package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/enfipy/spreest/src/helpers"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

func main() {
	b, err := ioutil.ReadFile("creds/credentials.json")
	helpers.PanicOnError(err)

	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	helpers.PanicOnError(err)
	client := helpers.GetClient("creds/token.json", config)

	srv, err := sheets.New(client)
	helpers.PanicOnError(err)

	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	readRange := "A1:C4"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	helpers.PanicOnError(err)

	if len(resp.Values) == 0 {
		log.Println("No data found.")
		return
	}

	for _, row := range resp.Values {
		log.Printf("%+v\n", row)
	}
}
