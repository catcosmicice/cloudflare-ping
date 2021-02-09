package main

import (
	"github.com/go-chi/chi"
	"github.com/sparrc/go-ping"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pinger, err := ping.NewPinger("discord.com")
		if err != nil {
			panic(err)
		}
		pinger.Count = 1
		pinger.SetPrivileged(true)
		pinger.Run()
		stats := pinger.Statistics()

		w.Write([]byte(strconv.Itoa(int(stats.AvgRtt.Milliseconds()))))
	})

	log.Fatal(http.ListenAndServe(":4242", r))
}
