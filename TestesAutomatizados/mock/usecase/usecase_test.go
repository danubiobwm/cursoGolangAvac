package usecase

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/danubiobwm/cursoGolangAvac/cursoGolangAvac/TestesAutomatizados/mock/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBuscarDados(t *testing.T) {
	t.Parallel()

	testes := []struct {
		nome string
		url  string

		mockResponse *http.Response
		mockErr      error

		responseEsperado map[string]any
		erroEsperado     error
	}{
		{
			nome:         "Erro na requisição",
			url:          "http://example.com/data",
			mockErr:      errors.New("erro de conexão"),
			erroEsperado: errors.New("erro de conexão"),
		},
		{
			nome: "Status não OK",
			url:  "http://example.com/data",
			mockResponse: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "500 Internal Server Error",
				Body:       io.NopCloser(bytes.NewBufferString("")),
			},
			erroEsperado: errors.New("status code inesperado"),
		},
		{
			nome: "Erro ao parsear JSON",
			url:  "http://example.com/data",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("<html><body>Erro</body></html>")),
			},
			erroEsperado: errors.New("json invalido"),
		},
		{
			nome: "Sucesso",
			url:  "http://example.com/data",
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"status":"success","message":"Dados recebidos com sucesso"}`)),
			},
			responseEsperado: map[string]any{
				"status":  "success",
				"message": "Dados recebidos com sucesso",
			},
		},
	}

	for _, teste := range testes {
		t.Run(teste.nome, func(t *testing.T) {
			t.Parallel()
			mockClient := new(mocks.HTTPClient)

			mockClient.On("Get", teste.url).Return(teste.mockResponse, teste.mockErr)

			// Executa o teste
			result, err := BuscarDados(mockClient, teste.url)

			// Valida o comportamento esperado
			assert.Equal(t, teste.responseEsperado, result)
			assert.Equal(t, teste.erroEsperado, err)

			// Verifica se o mock foi chamado corretamente
			mockClient.AssertExpectations(t)
		})
	}
}

func TestCriarDados(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		url  string
		data map[string]any

		mockResponse *http.Response
		mockErr      error

		wantResponse map[string]any
		wantErr      error
	}{
		{
			name:    "Erro na requisição",
			url:     "http://example.com/data",
			data:    map[string]any{"key": "value"},
			mockErr: errors.New("erro de conexão"),
			wantErr: errors.New("erro de conexão"),
		},
		{
			name: "Status não OK",
			url:  "http://example.com/data",
			data: map[string]any{"key": "value"},
			mockResponse: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Status:     "500 Internal Server Error",
				Body:       io.NopCloser(bytes.NewBufferString("")),
			},
			wantErr: errors.New("status code inesperado"),
		},
		{
			name: "Erro ao parsear JSON",
			url:  "http://example.com/data",
			data: map[string]any{"key": "value"},
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString("<html><body>Erro</body></html>")),
			},
			wantErr: errors.New("json invalido"),
		},
		{
			name: "Sucesso",
			url:  "http://example.com/data",
			data: map[string]any{"key": "value"},
			mockResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"status":"success","message":"Dados enviados com sucesso"}`)),
			},
			wantResponse: map[string]any{
				"status":  "success",
				"message": "Dados enviados com sucesso",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Configura o mock para cada cenário
			mockClient := new(mocks.HTTPClient)

			mockClient.On("Post", tt.url, "application/json", mock.Anything).Return(tt.mockResponse, tt.mockErr)

			// Executa o teste
			result, err := CriarDados(mockClient, tt.url, tt.data)

			// Valida o comportamento esperado
			assert.Equal(t, tt.wantResponse, result)
			assert.Equal(t, tt.wantErr, err)

			// Verifica se o mock foi chamado corretamente
			mockClient.AssertExpectations(t)
		})
	}
}
