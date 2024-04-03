package models


import "time"


type YouTubeData struct {
   IdKind                 string    `json:"id_kind"`
   IdVideoId              string    `json:"id_video_id"  gorm:"primaryKey"`
   PublishedAt            time.Time `json:"published_at"`
   ChannelId              string    `json:"channel_id"`
   Title                  string    `json:"title"`
   Description            string    `json:"description"`
   DefaultThumbnailUrl    string    `json:"default_thumbnail_url"`
   DefaultThumbnailWidth  int64     `json:"default_thumbnail_width"`
   DefaultThumbnailHeight int64     `json:"default_thumbnail_height"`
   MediumThumbnailUrl     string    `json:"medium_thumbnail_url"`
   MediumThumbnailWidth   int64     `json:"medium_thumbnail_width"`
   MediumThumbnailHeight  int64     `json:"medium_thumbnail_height"`
   HighThumbnailUrl       string    `json:"high_thumbnail_url"`
   HighThumbnailWidth     int64     `json:"high_thumbnail_width"`
   HighThumbnailHeight    int64     `json:"high_thumbnail_height"`
   PublishTime            time.Time `json:"publish_time"`
}





