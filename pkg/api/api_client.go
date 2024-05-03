package api

import (
	"context"
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/prometheus/common/log"

	fastshot "github.com/opus-domini/fast-shot"

	"github.com/yehormironenko/reseller/pkg/model"
)

type ResellerApiClient interface {
	GetBookByParams(ctx context.Context, bookname, author, genre string) (model.Books, string, error)
}

type resellerApiClient struct {
	HttpClient fastshot.ClientHttpMethods
}

func NewResellerApiClient(baseUrl string) ResellerApiClient {
	return &resellerApiClient{
		HttpClient: fastshot.DefaultClient(baseUrl),
	}
}

func (r resellerApiClient) GetBookByParams(ctx context.Context, bookname, author, genre string) (model.Books, string, error) {

	log.Info("GetBookByParams request with params ", bookname, author, genre)

	response, err := r.HttpClient.GET("search").Context().
		Set(ctx).Query().
		SetParams(map[string]string{"book": bookname, "author": author, "genre": genre}).
		Send()

	if err != nil {
		return model.Books{}, "", err
	}

	var responseModel model.Books

	if err = jsoniter.NewDecoder(response.RawResponse.Body).Decode(&responseModel); err != nil {
		return nil, response.Status(), fmt.Errorf("failed to unmarshal response")
	}

	return responseModel, response.Status(), nil
}
