package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"webtemp/shortvideoanalysis/db"
	"webtemp/shortvideoanalysis/models"
	"webtemp/shortvideoanalysis/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
}

// 微视链接前缀
var weseePrefix = "https://h5.weishi.qq.com/weishi/feed/"

// 微视链接正则，判断是否合法，还要通过这个正则提取feedID
var regWesee = regexp.MustCompile(`(?s-m)^https://h5\.weishi\.qq\.com/weishi/feed/[\d\w]+`)

// 微视提取真实地址正则
var regWeseeRealURL = regexp.MustCompile(`http://v\.weishi\.qq\.com/v\.weishi\.qq\.com/.+?mp4`)

// ParseShortVideo ...
func ParseShortVideo(c *gin.Context) {
	c.HTML(http.StatusOK, "shortvideo/parseshortvideo.html", gin.H{})
}

// ParseShortVideoByURL ...
func ParseShortVideoByURL(c *gin.Context) {
	parseURL := c.PostForm("parse_url")
	if strings.Contains(parseURL, "h5.weishi.qq.com") {
		c.JSON(http.StatusOK, gin.H{"real_url": handleURLWesee(parseURL)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"real_url": "解析失败！未分析过该短视频网站"})
}

// 处理微视短视频链接
func handleURLWesee(parseURL string) string {
	res := regWesee.FindString(parseURL)
	if res == "" {
		fmt.Println("微视分享链接不对！")
		return "微视分享链接不对！"
	}
	feedID := strings.Replace(res, weseePrefix, "", -1)

	// 查询mongodb，看有没有对应feedID的realUrl，有就返回realUrl给客户端，没有就解析然后插入数据库（下次解析直接从数据库查询到，提高效率）并返回realUrl给客户端
	// collection: short_video_url
	var shortVideoRes models.ShortVideoURL
	db.DbManager.FindOneShortVideo(bson.M{"feed_id": feedID}, &shortVideoRes)
	if shortVideoRes != (models.ShortVideoURL{}) {
		// 不为nil，则查询到，直接返回real_url_lossless字段
		return shortVideoRes.RealURLLossless
	}
	/*********************  解析网页获取真实地址并插入数据库  **********************/
	return analysisWeseeURL(feedID)
}

// 请求分析微视短视频链接
func analysisWeseeURL(feedID string) string {
	// 需要解析的微视短视频链接
	shortVideoURL := weseePrefix + feedID + "/?from=pc"
	// 请求返回body数据
	body := utils.Get(shortVideoURL)
	// 正则搜索真实地址，并返回
	realURL := string(regWeseeRealURL.Find(body))
	// 判断正则find结果是否为空，为空则没有该微视短视频，feedID错误
	if realURL == "" {
		return "该微视视频不存在！请检查短视频链接"
	}
	// 真实地址中的f0替换成f30得到无损压缩的视频地址
	realURLLossless := strings.Replace(realURL, ".f0.", ".f30.", 1)
	// 定义model数据
	shortVideoData := models.ShortVideoURL{
		FeedID:          feedID,
		RealURL:         realURL,
		RealURLLossless: realURLLossless,
	}
	res, err := db.DbManager.InsertOneShortVideo(shortVideoData)
	if err != nil {
		fmt.Println("db.DbManager.InsertOneShortVideo error:", err.Error())
		return "内部错误！"
	}
	fmt.Println("微视解析数据插入成功，ObjectID:", res)
	// 插入数据库中
	return realURLLossless
}
