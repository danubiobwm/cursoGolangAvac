package usecase

import (
	"bytes" // <-- Novo import necessário
	"encoding/json"
	"errors"
	"net/http"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/mock/repository"
)

func BuscarDados(client repository.HTTPClient, url string) (map[string]any, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code inesperado")
	}

	var response map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.New("json invalido")
	}

	return response, nil
}

func CriarDados(client repository.HTTPClient, url string, data map[string]any) (map[string]any, error) {
	// 1. Converte o map para JSON ([]byte)
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// 2. Passa o JSON convertido como um io.Reader para o client.Post
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated { // Muitas APIs retornam 201 Created no POST
		return nil, errors.New("status code inesperado")
	}

	var response map[string]any

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.New("json invalido")
	}

	return response, nil
}
