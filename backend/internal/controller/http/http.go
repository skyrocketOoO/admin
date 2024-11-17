package http

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type HttpServer struct{}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Stream(c *gin.Context) {
	videoPath := filepath.Join("data", "astar.mp4")

	file, err := os.Open(videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to open video file"})
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get file info"})
		return
	}

	// Serve content with Range support
	c.Writer.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(c.Writer, c.Request, stat.Name(), stat.ModTime(), file)
}
