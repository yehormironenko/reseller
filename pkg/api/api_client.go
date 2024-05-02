package api

import (
	"context"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/common/log"

	fastshot "github.com/opus-domini/fast-shot"

	"reseller/pkg/model"
)

type ResellerApiClient interface {
	GetBookByParams(ctx context.Context, bookname, author, genre string) (model.Books, error)
}

type resellerApiClient struct {
	BaseUrl    string
	HttpClient fastshot.Client
}

func NewResellerApiClient(baseUrl string, httpClient fastshot.Client) ResellerApiClient {
	return &resellerApiClient{
		BaseUrl:    baseUrl,
		HttpClient: httpClient,
	}
}

func (r resellerApiClient) GetBookByParams(ctx context.Context, bookname, author, genre string) (model.Books, error) {

	log.Info("GetBookByParams request with params ", bookname, author, genre)

	response, err := r.HttpClient.GET(r.BaseUrl).Context().
		Set(ctx).Query().
		SetParams(map[string]string{"book": bookname, "author": author, "genre": genre}).
		Send()

	if err != nil {
		return model.Books{}, err
	}

	var responseModel model.Books

	if err = jsoniter.NewDecoder(response.RawResponse.Body).Decode(&responseModel); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response")
	}

	return responseModel, nil
}
