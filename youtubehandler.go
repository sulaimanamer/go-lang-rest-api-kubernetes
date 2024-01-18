package main

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type YoutubeStats struct {
	Subscribers    int    `json:"subscribers"`
	ChannelName    string `json:"channelName"`
	MinutesWatched int    `json:"minutesWatched"`
	Views          int    `json:"views"`
}

func getChannelStats(k string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		//w.Write([]byte("response!"))
		yt := YoutubeStats{
			Subscribers:    5,
			ChannelName:    "my channel",
			MinutesWatched: 50,
			Views:          100,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(yt); err != nil {
			panic(err)
		}

	}
}
