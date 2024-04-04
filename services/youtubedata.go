package services

import (
	"fmt"
	"youtubedata/config/configs"
	"youtubedata/daos"
	"youtubedata/dtos"
	"youtubedata/models"
)

// IYouttubeServices defines the interface for YouTube data services
type IYouttubeServices interface {
	Upsert(req []*models.YouTubeData) error
	GetList(pg int) (*dtos.Response, error)
	GetSearchList(q string) ([]*models.YouTubeData, error)
}

var cnf *configs.Config

// YouttubeServices implements IYouttubeServices interface
type YouttubeServices struct {
	db daos.IYoutubeData
}

// NewYouttubeServices creates a new instance of YouttubeServices
func NewYouttubeServices() IYouttubeServices {
	cnf=configs.Get()
	return &YouttubeServices{
		db: daos.NewYoutubeData(),
	}
}

// Upsert inserts or updates YouTube data
func (y *YouttubeServices) Upsert(req []*models.YouTubeData) error {

	err := y.db.Upsert(req...)
	if err != nil {
		fmt.Println("not able to uppsert ")
		return err
	}
	return nil
}

// GetList retrieves a paginated list of YouTube data
func (y *YouttubeServices) GetList(pg int) (*dtos.Response, error) {

	data, err := y.db.GetPeginatedList(pg)
	if err != nil {
		fmt.Println("not able to uppsert ")
		return nil, err
	}

	total, err := y.db.GetListCount()
	if err != nil {
		fmt.Println("not able to uppsert ")
		return nil, err
	}
	return ListResponse(data, *total, int64(cnf.ItemsPerPage)), err
}


// GetSearchList retrieves a list of YouTube data based on search query
func (y *YouttubeServices) GetSearchList(q string) ([]*models.YouTubeData, error) {

	data, err := y.db.GetserachList(q)
	if err != nil {
		fmt.Println("not able to uppsert")
		return nil, err
	}
	return data, err
}
