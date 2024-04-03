package daos


import (
   "fmt"
   db "youtubedata/config"
   "youtubedata/models"


   "go.uber.org/zap"
)


type IYoutubeData interface {
   Upsert(req ...*models.YouTubeData) error
   Get(req *models.YouTubeData) (*models.YouTubeData, error)
   GetserachList(q string) ([]*models.YouTubeData, error)
   GetList(page int) ([]*models.YouTubeData, error)
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
   TableName = "youtubedata"
)


func (y *YoutubeData) Upsert(req ...*models.YouTubeData) error {
   err := y.db.DB.Debug().Table(TableName).Save(req).Error
   if err != nil {
       fmt.Println("not able upsert data", zap.Error(err))
       return err
   }
   return nil
}


func (y *YoutubeData) Get(req *models.YouTubeData) (*models.YouTubeData, error) {
   var res *models.YouTubeData
   err := y.db.DB.Debug().Table(TableName).Find(&res).Error
   if err != nil {
       fmt.Println("not able upsert data", zap.Error(err))
       return nil, err
   }
   return res, nil
}


func (y *YoutubeData) GetserachList(q string) ([]*models.YouTubeData, error) {
   var res []*models.YouTubeData
  q= "%"+q+"%"
   err := y.db.DB.Debug().Table(TableName).Or(`description ilike  ? `, q).Or(`title ilike ?`, q).Find(&res).Limit(10).Error
   if err != nil {
       fmt.Println("not able upsert data", zap.Error(err))
       return nil, err
   }
   return res, nil
}


func (y *YoutubeData) GetList(page int) ([]*models.YouTubeData, error) {
   var res []*models.YouTubeData
   offset := (page - 1) * 10
   err := y.db.DB.Debug().Table(TableName).Find(&res).Order("published_at DESC").Offset(offset).Limit(10).Error
   if err != nil {
       fmt.Println("not able upsert data", zap.Error(err))
       return nil, err
   }
   return res, nil
}


func (y *YoutubeData) GetListCount() (*int64, error) {
   var res int64
   err := y.db.DB.Debug().Table(TableName).Count(&res).Error
   if err != nil {
       fmt.Println("not able upsert data", zap.Error(err))
       return nil, err
   }
   return &res, nil
}





