package main

// Importing dependencies
import (
	// "log"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"os"
	"time"
)

var clientID string

func init() {
	clientID = os.Getenv("CLIENT_ID")
	/*
		streamers := [64]string{
			"cypheroftyr",
			"ind1fference",
			"nymgamer_saffista",
			"savepointsirens",
			"jpayano",
			"spawnonme",
			"videogamesandthebible",
			"spawnonme2",
			"xFats",
			"beyondthegameplay",
			"jaxnos",
			"RaqibMarvelous",
			"ugrgaming",
			"rophoenixprofile",
			"reefjackson",
			"madamstarr",
			"deepboyce",
			"dessiblissdashboard",
			"simplyundrea",
			"blackblizzard66",
			"dimples216profile",
			"majorlinux",
			"ahmedz",
			"sharkyshood",
			"gynesis",
			"divinejustiniav ",
			"tonymizugaming",
			"Bunneh3000 ",
			"dmgphenom",
			"roninonthestickz",
			"involuntarysass",
			"namikuro",
			"mikeauxlprofile",
			"angelic_sweety8D",
			"xdianamoonx",
			"angelic_sweety8D",
			"bboyspaz",
			"TheRealEWord",
			"yumi_mads",
			"daylightful",
			"CrownedxRoyalty",
			"data_dave",
			"blackoni",
			"fortheyuri",
			"kuromon20",
			"alishabikes",
			"dognolia",
			"jdubeous",
			"brownandnerdy",
			"blackmageblake",
			"Kodzo_Gaming",
			"majorlinux",
			"BreezyCarter24 ",
			"ellesria",
			"imthemyth",
			"Optimus_blu",
			"wimsy113",
			"tjx2045",
			"auronblade60",
			"killerrue",
			"supercuddles",
			"8bitoffun",
			"idaikoruberu",
			"ishadaha",
		}
	*/
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_login=BobRoss", nil)
	fmt.Println(err)
	req.Header.Add("Client-ID", clientID)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var dat map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["data"])
}
