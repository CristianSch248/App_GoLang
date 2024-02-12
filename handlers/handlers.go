package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

type Vagalume struct {
	Key struct {
		key string `toml:"key"`
	} `toml:"vagalume"`
}

/*
Resposta é uma estrutura que define a estrutura de dados da resposta
JSON que esperamos receber da API do Vagalume.

Ela possui campos correspondentes aos diferentes aspectos da resposta,
como tipo, informações do artista, informações das músicas, etc.
*/
type Resposta struct {
	Type string `json:"type"`
	Art  struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"art"`
	Mus []struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		Lang      int    `json:"lang"`
		Text      string `json:"text"`
		Translate []struct {
			ID   string `json:"id"`
			Lang int    `json:"lang"`
			URL  string `json:"url"`
			Text string `json:"text"`
		} `json:"translate"`
	} `json:"mus"`
	Badwords bool `json:"badwords"`
}

func GetMusicas(w http.ResponseWriter, r *http.Request) {

	var config Vagalume
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		fmt.Println("Erro ao carregar arquivo TOML:", err)
		return
	}
	os.Setenv("key", config.Key.key)
	apiKey := os.Getenv("key")

	artist := "ABBA"
	song := "Fernando"

	url := "https://api.vagalume.com.br/search.php?art=" + artist + "&mus=" + song + "&apikey=" + apiKey

	resposta, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer a requisição: %v", err)
		return
	}

	defer resposta.Body.Close()

	var corpoResposta Resposta

	err = json.NewDecoder(resposta.Body).Decode(&corpoResposta)
	if err != nil {
		fmt.Printf("Erro ao decodificar o JSON: %v", err)
		return
	}

	respostaJSON, err := json.Marshal(corpoResposta)
	if err != nil {
		fmt.Printf("Erro ao converter para JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respostaJSON)
}
