package model

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
