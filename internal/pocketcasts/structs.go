package pocketcasts

import "time"

type EpisodeList struct {
	Episodes []Episode `json:"episodes"`
}

type Episode struct {
	ID            int    `json:"id"`
	UUID          string `json:"uuid"`
	URL           string `json:"url"`
	PublishedAt   string `json:"published_at"`
	Duration      int    `json:"duration"`
	FileType      string `json:"file_type"`
	Title         string `json:"title"`
	Size          int    `json:"size"`
	PlayingStatus int    `json:"playing_status"`
	PlayedUpTo    int    `json:"played_up_to"`
	IsDeleted     bool   `json:"is_deleted"`
	Starred       bool   `json:"starred"`
	PodcastUUID   string `json:"podcast_uuid"`
}

type All struct {
	Podcasts []Podcast `json:"podcasts"`
	App      App       `json:"app"`
}

type Podcast struct {
	ID                int    `json:"id"`
	UUID              string `json:"uuid"`
	URL               string `json:"url"`
	Title             string `json:"title"`
	Description       string `json:"description"`
	ThumbnailURL      string `json:"thumbnail_url"`
	Author            string `json:"author"`
	EpisodesSortOrder int    `json:"episodes_sort_order"`
}

type App struct {
	UserVersionCode int     `json:"userVersionCode"`
	VersionCode     int     `json:"versionCode"`
	VersionName     float64 `json:"versionName"`
	VersionSummary  string  `json:"versionSummary"`
}

type UUID struct {
	ID string `json:"uuid"`
}

//{"uuid":"1ce1cdf0-d55e-0134-ebdd-4114446340cb","podcast_uuid":"39aba5d0-0428-012e-f9a0-00163e1b201c","playing_status":2,"duration":3614.422774,"played_up_to":2126.58121}

type Position struct {
	ID            string        `json:"uuid"`
	PodcastID     string        `json:"podcast_uuid"`
	PlayingStatus int           `json:"playing_status"`
	Duration      time.Duration `json:"duration"`
	Played        time.Duration `json:"played_up_to"`
}
