-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "youtubedata" (
   "id_kind" TEXT,
   "id_video_id" TEXT,
   "published_at" TIMESTAMP,
   "channel_id" TEXT,
   "title" TEXT,
   "description" TEXT,
   "default_thumbnail_url" TEXT,
   "default_thumbnail_width" INT,
   "default_thumbnail_height" INT,
   "medium_thumbnail_url" TEXT,
   "medium_thumbnail_width" INT,
   "medium_thumbnail_height" INT,
   "high_thumbnail_url" TEXT,
   "high_thumbnail_width" INT,
   "high_thumbnail_height" INT,
   "channel_title" TEXT,
   "live_broadcast_content" TEXT,
   "publish_time" TIMESTAMP,
   CONSTRAINT youtube_data PRIMARY KEY (id_video_id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "youtubedata";
-- +goose StatementEnd



