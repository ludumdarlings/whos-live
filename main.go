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

type StreamResponse struct {
	Data []Stream `json:"data"`
	Pagination map[string]interface{} `json:"pagination"`
}

type UserResponse struct {
	Data []User `json:"data"`
}

type Stream struct {
    UserId string `json:"user_id"` //105458682 
    GameId string `json:"game_id"`//488191 
    Title string `json:"title"`//Bob Ross Marathon! 
    ViewerCount int `json:"viewer_count"`
	Id string `json:"id"`
    CommunityIds []string `json:"community_ids"` //[275e4fd5-1b81-4024-a85a-bb84e77e760e] 
    Type string `json:"type"` //live 
    StartedAt string `json:"started_at"` //2017-09-16T01:02:03Z 
    Language string `json:"language"`//en 
    ThumbnailURL string `json:"thumbnail_url"` //https://static-cdn.jtvnw.net/previews-ttv/live_user_bobross-{width}x{height}.jpg id:26266286656
}

type User struct {
   Id string `json:"id"`
   Login string `json:"login"`
   DisplayName string `json:"display_name"`
   Type string `json:"type"`
   BroadcasterType string `json:"broadcaster_type"`
   Description string `json:"description"`
   ProfileImageURL string `json:"profile_image_url"`
   OfflineImageURL string `json:"offline_image_url"`
   ViewCount int `json:"view_count"`
   Email string `json:"email"`
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/streams?user_login=bobRoss", nil)
	fmt.Println(err)
	req.Header.Add("Client-ID", clientID)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var liveStreams StreamResponse

	if err := json.NewDecoder(resp.Body).Decode(&liveStreams); err != nil {
		panic(err)
	}
	fmt.Println(liveStreams.Data[0].Title)
}
