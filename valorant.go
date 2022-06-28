package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ValorantData struct {
	Player ValorantPlayer
}

func getJson(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (s *ValorantData) Fetch(address string) error {

	var (
		playerJson []byte
		err        error
	)

	if playerJson, err = getJson("https://api.henrikdev.xyz/valorant/v2/mmr/eu/" + address); err != nil {
		return err
	}

	if err := json.Unmarshal(playerJson, &s.Player); err != nil {
		return err
	}

	if err := verifyValorantData(s); err != nil {
		return err
	}

	return nil
}

func verifyValorantData(data *ValorantData) error {

	ErrInvalidValorantData := "Invalid Valorant Data"

	if data.Player.Data.Name == "" {
		return errors.New(ErrInvalidValorantData)
	}

	if data.Player.Data.Tag == "" {
		return errors.New(ErrInvalidValorantData)
	}

	if data.Player.Data.CurrentData.Elo == 0 {
		return errors.New(ErrInvalidValorantData)
	}

	return nil
}

type ValorantPlayer struct {
	Status int `json:"status"`
	Data   struct {
		Name        string `json:"name"`
		Tag         string `json:"tag"`
		Puuid       string `json:"puuid"`
		CurrentData struct {
			Currenttier          int    `json:"currenttier"`
			Currenttierpatched   string `json:"currenttierpatched"`
			RankingInTier        int    `json:"ranking_in_tier"`
			MmrChangeToLastGame  int    `json:"mmr_change_to_last_game"`
			Elo                  int    `json:"elo"`
			GamesNeededForRating int    `json:"games_needed_for_rating"`
			Old                  bool   `json:"old"`
		} `json:"current_data"`
		BySeason struct {
			E5A3 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a3"`
			E5A2 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a2"`
			E5A1 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e5a1"`
			E4A3 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a3"`
			E4A2 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a2"`
			E4A1 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e4a1"`
			E3A3 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a3"`
			E3A2 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a2"`
			E3A1 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e3a1"`
			E2A3 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a3"`
			E2A2 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a2"`
			E2A1 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e2a1"`
			E1A3 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a3"`
			E1A2 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a2"`
			E1A1 struct {
				Wins             int    `json:"wins"`
				NumberOfGames    int    `json:"number_of_games"`
				FinalRank        int    `json:"final_rank"`
				FinalRankPatched string `json:"final_rank_patched"`
				ActRankWins      []struct {
					PatchedTier string `json:"patched_tier"`
					Tier        int    `json:"tier"`
				} `json:"act_rank_wins"`
				Old bool `json:"old"`
			} `json:"e1a1"`
		} `json:"by_season"`
	} `json:"data"`
}
