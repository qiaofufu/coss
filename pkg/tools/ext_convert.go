package tools

import "path/filepath"

var contentTypeMap = map[string]string{
	".txt":  "text/plain",
	".html": "text/html",
	".css":  "text/css",
	".js":   "application/javascript",
	".json": "application/json",
	".xml":  "application/xml",
	".pdf":  "application/pdf",
	".zip":  "application/zip",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".ppt":  "application/vnd.ms-powerpoint",
	".pptx": "application/vnd.openxmlformats-officedocument.presentationml.presentation",
	".jpeg": "image/jpeg",
	".jpg":  "image/jpeg",
	".png":  "image/png",
	".gif":  "image/gif",
	".bmp":  "image/bmp",
	".ico":  "image/x-icon",
	".svg":  "image/svg+xml",
	".mp3":  "audio/mpeg",
	".ogg":  "audio/ogg",
	".wav":  "audio/wav",
	".mp4":  "video/mp4",
	".webm": "video/webm",
	".csv":  "text/csv",
	".tsv":  "text/tab-separated-values",
	".yaml": "application/x-yaml",
	".yml":  "application/x-yaml",
	".exe":  "application/octet-stream",
	".dll":  "application/octet-stream",
	".dat":  "application/octet-stream",
	".bin":  "application/octet-stream",
	".sql":  "application/sql",
	".log":  "text/plain",
}

func GetContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	contentType, exists := contentTypeMap[ext]
	if !exists {
		return "application/octet-stream" // default to binary if content type is unknown
	}
	return contentType
}
