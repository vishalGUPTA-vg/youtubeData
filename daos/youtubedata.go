package daos

import (
	"log"
	db "youtubedata/config"
	"youtubedata/models"

	"go.uber.org/zap"
)

type IYoutubeData interface {
	Upsert(req ...*models.YouTubeData) error
	Get(req *models.YouTubeData) (*models.YouTubeData, error)
	GetserachList(q string) ([]*models.YouTubeData, error)
	GetPeginatedList(page int) ([]*models.YouTubeData, error)
	GetListCount() (*int64, error)
}

type YoutubeData struct {
	db *db.DBConn
}

func NewYoutubeData() IYoutubeData {
	dburl := db.New()
	return &YoutubeData{
		db: dburl,
	}
}

const (
   // tableName defines the name of the database table
	tableName = "youtubedata"
)

// Upsert inserts or updates YouTube data
func (y *YoutubeData) Upsert(req ...*models.YouTubeData) error {
	err := y.db.DB.Debug().Table(tableName).Save(req).Error
	if err != nil {
		log.Println("not able upsert data", zap.Error(err))
		return err
	}
	return nil
}

// Get retrieves YouTube data by ID
func (y *YoutubeData) Get(req *models.YouTubeData) (*models.YouTubeData, error) {
	var res *models.YouTubeData
	err := y.db.DB.Debug().Table(tableName).Find(&res).Error
	if err != nil {
		log.Println("not able upsert data", zap.Error(err))
		return nil, err
	}
	return res, nil
}

// GetserachList retrieves YouTube data based on search query
func (y *YoutubeData) GetserachList(q string) ([]*models.YouTubeData, error) {
	var res []*models.YouTubeData
	q = "%" + q + "%"
	err := y.db.DB.Debug().Table(tableName).Or(`description ilike  ? `, q).Or(`title ilike ?`, q).Find(&res).Limit(10).Error
	if err != nil {
		log.Println("not able upsert data", zap.Error(err))
		return nil, err
	}
	return res, nil
}

// GetList retrieves a paginated list of YouTube data
func (y *YoutubeData) GetPeginatedList(page int) ([]*models.YouTubeData, error) {
	var res []*models.YouTubeData
	offset := (page - 1) * 10
	err := y.db.DB.Debug().Table(tableName).Find(&res).Order("published_at DESC").Offset(offset).Limit(10).Error
	if err != nil {
		log.Println("not able upsert data", zap.Error(err))
		return nil, err
	}
	return res, nil
}

// GetListCount retrieves the total count of YouTube data
func (y *YoutubeData) GetListCount() (*int64, error) {
	var res int64
	err := y.db.DB.Debug().Table(tableName).Count(&res).Error
	if err != nil {
		log.Println("not able upsert data", zap.Error(err))
		return nil, err
	}
	return &res, nil
}
