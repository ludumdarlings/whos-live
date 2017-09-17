package main

import (
	"fmt"
	"io/ioutil"
	// "reflect"
	// "regexp"
	// "strconv"
	// "strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type Stringer interface {
	String() string
}

func main() {
	data, err := ioutil.ReadFile("client_secret.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet("1ALqjmtsdFTibM1ZTIBPgOKPSs6Brtu_IATVWlA_Vi7o")
	checkError(err)
	sheet, err := spreadsheet.SheetByIndex(0)
	checkError(err)
	streamers_dirty := sheet.Columns[3]
	removeBlanks(streamers_dirty)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func removeBlanks(cells []spreadsheet.Cell) {
	for n := 0; n <= 1000; n++ {
		cell := cells[n].Value
		if cell != "" {
			fmt.Println(cell)
		}
	}
}
