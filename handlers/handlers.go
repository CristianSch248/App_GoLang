package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/CristianSch248/App_GoLang/model"

	"github.com/BurntSushi/toml"

	"context"
	"io"

	language "cloud.google.com/go/language/apiv2"
	"cloud.google.com/go/language/apiv2/languagepb"
)

type Vagalume struct {
	Key struct {
		key string `toml:"key"`
	} `toml:"vagalume"`
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

	var corpoResposta model.Resposta

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

func AnalizeMusic() {
	// analyzeSentiment(musica)
}

func analyzeSentiment(w io.Writer, text string) error {
	ctx := context.Background()

	// Initialize client.
	client, err := language.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	resp, err := client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})

	if err != nil {
		return fmt.Errorf("AnalyzeSentiment: %w", err)
	}
	fmt.Fprintf(w, "Response: %q\n", resp)

	return nil
}
