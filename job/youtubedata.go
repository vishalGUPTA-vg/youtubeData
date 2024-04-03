package job

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	config "youtubedata/config/configs"
	"youtubedata/models"
	"youtubedata/services"

	"time"

	"go.uber.org/zap"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func YoutubeJob() {
	flag.Parse()
	cnf := config.Get()

	keys := cnf.YoutubeApiKeys
	keyIndex := 0
	initalTime := "2024-04-03T00:00:00Z"
	for i := 0; i <= 2; i++ {
		var ctx context.Context
		service, err := youtube.NewService(ctx, option.WithAPIKey(keys[keyIndex]))
		if err != nil {
			log.Fatalf("Error creating new YouTube client: %v", err)
		}

		// Make the API call to YouTube.
		call := service.Search.List([]string{"snippet"}).
			Q("football").
			MaxResults(50).
			Order("date").
			PublishedAfter(initalTime).
			Type("video")

		response, err := call.Do()
		if err != nil {
			fmt.Println("error not able process", zap.Any("err", err))
			if response.HTTPStatusCode == http.StatusForbidden {
				fmt.Println("Error making API call:", err)
				if keyIndex < len(keys)-1 {
					keyIndex++ // Switch to the next API key if the current one expires
					continue
				}

				//break the loop if all api key have expired
				if keyIndex == len(keys)-1 {
					fmt.Println("all api keys have expired ", err)
					break
				}
			}
			return
		}
		fmt.Println(*response)
		if len(response.Items) > 0 {
			services := services.NewYouttubeServices()
			err = services.Upsert(maptomodels(response))
			if err != nil {
				fmt.Println("error not able process")
			}
		}
		initalTime = time.Now().UTC().Format(time.RFC3339)

		time.Sleep(10 * time.Second)
	}

}

func maptomodels(req *youtube.SearchListResponse) []*models.YouTubeData {
	var response []*models.YouTubeData
	for _, val := range req.Items {
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", val.Snippet.PublishedAt)
		if err != nil {
			fmt.Println("error not able process ", err)
		}
		res := &models.YouTubeData{
			IdKind:                 val.Id.Kind,
			IdVideoId:              val.Id.VideoId,
			PublishedAt:            parsedTime,
			ChannelId:              val.Id.ChannelId,
			Title:                  val.Snippet.Title,
			Description:            val.Snippet.Description,
			DefaultThumbnailUrl:    val.Snippet.Thumbnails.Default.Url,
			DefaultThumbnailWidth:  val.Snippet.Thumbnails.Default.Width,
			DefaultThumbnailHeight: val.Snippet.Thumbnails.Default.Height,
			MediumThumbnailUrl:     val.Snippet.Thumbnails.Medium.Url,
			MediumThumbnailWidth:   val.Snippet.Thumbnails.Medium.Width,
			MediumThumbnailHeight:  val.Snippet.Thumbnails.Medium.Height,
			HighThumbnailUrl:       val.Snippet.Thumbnails.High.Url,
			HighThumbnailWidth:     val.Snippet.Thumbnails.High.Width,
			HighThumbnailHeight:    val.Snippet.Thumbnails.High.Height,
		}
		response = append(response, res)
	}

	return response
}
