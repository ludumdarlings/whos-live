package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

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
	streamerlist := removeBlanks(streamers_dirty)
	twitchlist := pickUrls(streamerlist, "twitch.tv/")
	mixerlist := pickUrls(streamerlist, "mixer.com")
	twitchlist = getUsernames(twitchlist, "twitch.tv/")
	mixerlist = getUsernames(mixerlist, "mixer.com/")
	fmt.Println()
	fmt.Println(twitchlist)
	fmt.Println()
	fmt.Println(mixerlist)
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func getUsernames(list []string, matchurl string) []string {
	a := []string{}
	total := len(list)
	for n := 0; n < total; n++ {
		url := strings.SplitAfter(list[n], matchurl)
		newurl := strings.Split(url[1], "/")
		a = append(a, newurl[0])
	}
	return a
}

func removeBlanks(cells []spreadsheet.Cell) []string {
	a := []string{}
	for n := 0; n <= 1000; n++ {
		cell := cells[n].Value
		if cell != "" {
			a = append(a, cell)
		}
	}
	return a
}

func pickUrls(list []string, matchurl string) []string {
	a := []string{}
	total := len(list)
	matchurl = strings.ToLower(matchurl)
	for n := 0; n < total; n++ {
		url := strings.ToLower(list[n])
		if strings.Contains(url, matchurl) {
			a = append(a, url)
		}
	}
	return a
}
