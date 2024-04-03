package services

import (
	"fmt"
	"youtubedata/config/configs"
	"youtubedata/daos"
	"youtubedata/dtos"
	"youtubedata/models"
)

type IYouttubeServices interface {
	Upsert(req []*models.YouTubeData) error
	GetList(pg int) (*dtos.Response, error)
	GetSearchList(q string) ([]*models.YouTubeData, error)
}

var cnf *configs.Config

type YouttubeServices struct {
	db daos.IYoutubeData
}

func NewYouttubeServices() IYouttubeServices {
	cnf=configs.Get()
	return &YouttubeServices{
		db: daos.NewYoutubeData(),
	}
}

func (y *YouttubeServices) Upsert(req []*models.YouTubeData) error {

	err := y.db.Upsert(req...)
	if err != nil {
		fmt.Println("not able to uppsert ")
		return err
	}
	return nil
}

func (y *YouttubeServices) GetList(pg int) (*dtos.Response, error) {

	data, err := y.db.GetList(pg)
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

func (y *YouttubeServices) GetSearchList(q string) ([]*models.YouTubeData, error) {

	data, err := y.db.GetserachList(q)
	if err != nil {
		fmt.Println("not able to uppsert")
		return nil, err
	}
	return data, err
}
