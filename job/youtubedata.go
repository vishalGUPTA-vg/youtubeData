package job

import (
	"context"
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
	// Get configuration
	cnf := config.Get()
	// Get YouTube API keys from configuration
	keys := cnf.YoutubeApiKeys
	// Initialize index for accessing API keys
	keyIndex := 0
	// Set initial time for API call
	initalTime := "2024-04-03T00:00:00Z"
	for i := 0; i <= 2; i++ {
		// Initialize context
		var ctx context.Context
		// Create a new YouTube service client using the current API key
		service, err := youtube.NewService(ctx, option.WithAPIKey(keys[keyIndex]))
		if err != nil {
			log.Fatalf("Error creating new YouTube client: %v", err)
		}

		// Make the API call to search for football videos published after the initial time
		call := service.Search.List([]string{"snippet"}).
			Q("football").
			MaxResults(50).
			Order("date").
			PublishedAfter(initalTime).
			Type("video")

			// Execute the API call
		response, err := call.Do()
		if err != nil {

			// Check if the error is due to an API key issue
			if response.HTTPStatusCode == http.StatusForbidden {
				log.Println("Error making API call:", err)
				if keyIndex < len(keys)-1 {
					// Switch to the next API key if the current one expires
					keyIndex++
					continue
				}

				//break the loop if all api key have expired
				if keyIndex == len(keys)-1 {
					log.Println("all api keys have expired ", err)
					break
				}
			}
			// Handle API call errors
			log.Fatal("error not able connect to youtube", zap.Any("err", err))

			return
		}
		// If API call is successful and there are search results
		if len(response.Items) > 0 {
			// Convert API response to models and upsert into the database
			services := services.NewYouttubeServices()
			err = services.Upsert(maptomodels(response))
			if err != nil {
				log.Println("error unable to upsert")
			}
		}

		// Update initial time to current time for the next API call
		initalTime = time.Now().UTC().Format(time.RFC3339)

		// Pause execution for 10 seconds before next iteration
		time.Sleep(10 * time.Second)
	}

}

// maptomodels converts API response to models
func maptomodels(req *youtube.SearchListResponse) []*models.YouTubeData {
	var response []*models.YouTubeData
	for _, val := range req.Items {
		// Parse published time
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z", val.Snippet.PublishedAt)
		if err != nil {
			log.Println("error not able process ", err)
		}
		// Map API response to models
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
