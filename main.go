package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	port   = flag.String("port", "1091", "The port to listen on for HTTP requests.")
	listen = flag.String("listen", "localhost", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()

	log.Println("Starting valorant exporter on http://" + *listen + ":" + *port + " ...")

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/probe", func(w http.ResponseWriter, req *http.Request) {
		probeHandler(w, req)
	})

	log.Fatal(http.ListenAndServe(*listen+":"+*port, nil))
}

func probeHandler(w http.ResponseWriter, r *http.Request) {

	var (
		valorant_elo = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "valorant_elo",
				Help: "Current elo of player.",
			},
			[]string{"username", "tagline"},
		)
		valorant_tier = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "valorant_tier",
				Help: "Current tier of player.",
			},
			[]string{"username", "tagline"},
		)
		valorant_games = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "valorant_games",
				Help: "Number of games played.",
			},
			[]string{"username", "tagline"},
		)
		valorant_wins = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Name: "valorant_wins",
				Help: "Number of games won.",
			},
			[]string{"username", "tagline"},
		)
	)

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()
	r = r.WithContext(ctx)

	// get ?target= parameter from request
	target := r.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}

	// create registry containing metrics
	registry := prometheus.NewPedanticRegistry()

	// add metrics to registry
	registry.MustRegister(valorant_elo)
	registry.MustRegister(valorant_tier)
	registry.MustRegister(valorant_games)
	registry.MustRegister(valorant_wins)

	// get shelly data from target
	var data ValorantData
	if err := data.Fetch(target); err != nil {
		// TODO better error handling
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	// define labels used for all metrics
	var valorant_labels prometheus.Labels = prometheus.Labels{"username": data.Player.Data.Name, "tagline": data.Player.Data.Tag}

	// set metrics
	valorant_elo.With(prometheus.Labels(valorant_labels)).Set(float64(data.Player.Data.CurrentData.Elo))
	valorant_tier.With(prometheus.Labels(valorant_labels)).Set(float64(data.Player.Data.CurrentData.Currenttier))

	// set games played
	gamesEpisode1 := data.Player.Data.BySeason.E1A1.NumberOfGames + data.Player.Data.BySeason.E1A2.NumberOfGames + data.Player.Data.BySeason.E1A3.NumberOfGames
	gamesEpisode2 := data.Player.Data.BySeason.E2A1.NumberOfGames + data.Player.Data.BySeason.E2A2.NumberOfGames + data.Player.Data.BySeason.E2A3.NumberOfGames
	gamesEpisode3 := data.Player.Data.BySeason.E3A1.NumberOfGames + data.Player.Data.BySeason.E3A2.NumberOfGames + data.Player.Data.BySeason.E3A3.NumberOfGames
	gamesEpisode4 := data.Player.Data.BySeason.E4A1.NumberOfGames + data.Player.Data.BySeason.E4A2.NumberOfGames + data.Player.Data.BySeason.E4A3.NumberOfGames
	gamesEpisode5 := data.Player.Data.BySeason.E5A1.NumberOfGames + data.Player.Data.BySeason.E5A2.NumberOfGames + data.Player.Data.BySeason.E5A3.NumberOfGames
	games := gamesEpisode1 + gamesEpisode2 + gamesEpisode3 + gamesEpisode4 + gamesEpisode5
	valorant_games.With(prometheus.Labels(valorant_labels)).Set(float64(games))

	// set games won
	winsEpisode1 := data.Player.Data.BySeason.E1A1.Wins + data.Player.Data.BySeason.E1A2.Wins + data.Player.Data.BySeason.E1A3.Wins
	winsEpisode2 := data.Player.Data.BySeason.E2A1.Wins + data.Player.Data.BySeason.E2A2.Wins + data.Player.Data.BySeason.E2A3.Wins
	winsEpisode3 := data.Player.Data.BySeason.E3A1.Wins + data.Player.Data.BySeason.E3A2.Wins + data.Player.Data.BySeason.E3A3.Wins
	winsEpisode4 := data.Player.Data.BySeason.E4A1.Wins + data.Player.Data.BySeason.E4A2.Wins + data.Player.Data.BySeason.E4A3.Wins
	winsEpisode5 := data.Player.Data.BySeason.E5A1.Wins + data.Player.Data.BySeason.E5A2.Wins + data.Player.Data.BySeason.E5A3.Wins
	wins := winsEpisode1 + winsEpisode2 + winsEpisode3 + winsEpisode4 + winsEpisode5
	valorant_wins.With(prometheus.Labels(valorant_labels)).Set(float64(wins))

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)

}
