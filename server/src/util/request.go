package util

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"net/http"
	"src/util/config"
	"time"
)

func HttpGet(url string, c *gin.Context) (string, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Timeout:   5 * time.Second, //默认5秒超时时间
		Transport: tr,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	/*if config.Conf.Jaeger.Open{
		tracer,_:=c.Get("Tracer")
		parent
	}*/
}
