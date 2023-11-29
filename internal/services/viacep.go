package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/guhstanley/go-viacep/internal/models"
)

const (
	baseUrl string = "http://viacep.com.br/ws/"
)

type ViaCepService struct {
	BaseUrl string `json:"url"`
}

func NewViaCepService() (*ViaCepService, error) {
	return &ViaCepService{
		BaseUrl: baseUrl,
	}, nil
}

func (v *ViaCepService) GetJsonURL(cep string) string {
	return v.BaseUrl + cep + "/json/"
}

func (v *ViaCepService) Execute(url string) (*models.ViaCep, error) {
	viacep := &models.ViaCep{}

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	err = json.Unmarshal(res, viacep)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return viacep, nil
}
