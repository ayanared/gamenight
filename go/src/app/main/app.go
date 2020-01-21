package main

import (
	"net/http"
	"model"
	"time"
	"encoding/json"
	"fmt"
	"strings"
)

type gameDetail struct {
	Game *model.Game
	GamePlayers []*model.GamePlayer
	Players []*model.Player
}

func detail(w http.ResponseWriter, r *http.Request) {
	//this should only be called with a get request
	if strings.ToLower(r.Method) != "get" {
		http.Error(w, fmt.Errorf("Unsupported Method").Error(), http.StatusMethodNotAllowed)
		return
	}

	gd := &gameDetail{
		Players: []*model.Player{
			{
				Name: "Tony Stark",
				Email: "ironman@mcu.fake",
			},
			{
				Name: "T'Challa",
				Email: "blackpanther@mcu.fake",
			},
			{
				Name: "Natasha Romanoff",
				Email: "blackwidow@mcu.fake",
			},
		},

		Game: &model.Game{
			ID: "1",
			Time: time.Date(2050, time.December, 31, 7, 0, 0, 0, time.Local),
			Name: "Brass Birmingham",
			BGGLink: "https://boardgamegeek.com/boardgame/224517/brass-birmingham",
			Location: "Avengers Mansion",
			MinPlayers: 2,
			MaxPlayers: 4,
		},

		GamePlayers: []*model.GamePlayer{
			{
				PlayerEmail: "ironman@mcu.fake",
				GameID: "1",
			},
			{
				PlayerEmail: "blackpanther@mcu.fake",
				GameID: "1",
			},
		},
	} //end of gd

	fmt.Printf("gd\n%+v\n\n", gd)
	
	gdBytes, err := json.Marshal(gd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}

	fmt.Printf("%s\n", gdBytes)

	w.Header()["Content-Type"] = []string{"application/json"}

	_, err = w.Write(gdBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    return
	}

}

func main() {
 http.HandleFunc("/", detail)
 err := http.ListenAndServe(":8080", nil)
 if err != nil {
	 panic(err)
 }
}
