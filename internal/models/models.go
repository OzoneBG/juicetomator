package models

// DownloadInfo contains all required information to download videos.
type DownloadInfo struct {
	Username string   `json:"username" binding:"required"`
	Songs    []string `json:"songs" binding:"required"`
}
