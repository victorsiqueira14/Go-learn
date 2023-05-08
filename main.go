package main

import (
	"encoding/json"
	"io"
	"net/http"
	"fmt"
)

type ViaCep struct {
	Cep         string `json:"cep"` // o que tiver no json, vai ser armazenado na variavel cep
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

// com o asterisco eu estou passando o endereço de memoria da variavel e alterando o valor dela
// sem isso ela so esta copiando o valor e duplicando
func (v *ViaCep) SetCep(cep string) {
	v.Cep = cep
	fmt.Println(v.Cep)
}

//method, pois eu tenho a struct anexada a ela
func (v ViaCep) EnderecoCompleto() string {
	return fmt.Sprintf("%s, %s, %s, %s, %s, %s", v.Cep, v.Logradouro, v.Complemento, v.Bairro, v.Localidade, v.Uf)
}

func main() {
	cep := "01001000"
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	//variavel é do tipo viaCep
	var data ViaCep
	err = json.Unmarshal(res, &data) // o & é para passar o endereço de memoria, onde vai ser armazenado o valor
	if err != nil {
		panic(err)
	}
	fmt.Println(data.EnderecoCompleto())
	data.SetCep("32234324")
	fmt.Println(data.Cep)
}
