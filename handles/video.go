package handles

import (
	"net/http"
	"regexp"

	"videocrawler"

	"github.com/gin-gonic/gin"
)

type VideoHandle struct {
}

func InitVideoHandle(e *gin.Engine) {
	v := VideoHandle{}

	// Setup Routes
	e.POST("/video/get", v.get)
}

func (v *VideoHandle) get(c *gin.Context) {
	videoUrl := c.DefaultPostForm("url", "")
	var title string
	var m3 map[string]string
	var flv = make([]string, 0)
	var domainLimit = false
	if videoUrl != "" {
		videoInfo, err := videocrawler.GetVideoInfo(videoUrl)
		if err == nil {
			title = videoInfo.Title
			m3 = videoInfo.M3
			flv = videoInfo.Flv
			domainLimit = videoInfo.DomainLimit
		}
	}
	m3Download := false
	r, _ := regexp.Compile(".+youku.+")
	for _, v := range m3 {
		if v != "" {
			if r.MatchString(v) {
				m3Download = true
			}
		}
		break
	}
	c.JSON(http.StatusOK, gin.H{
		"title":        title,
		"m3":           m3,
		"flv":          flv,
		"domain_limit": domainLimit,
		"m3_download":  m3Download,
	})
}
