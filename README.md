目录结构
```go
helper: 提供了管理用户token的帮助方法
im: 提供了素材上传和给用户发送消息的接口
oauth： oauth2.0相关接口
resouce： 视频和图片相关接口
service：抖音相关服务
user: 用户接口
util: 工具方法
```


例子：
```go

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/owen-gxz/douyin-sdk/im"
	"github.com/owen-gxz/douyin-sdk/oauth"
	"github.com/owen-gxz/douyin-sdk/resource"
	"github.com/owen-gxz/douyin-sdk/service"
	"github.com/owen-gxz/douyin-sdk/user"
	"github.com/rs/xid"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	openid    = "openid"
	token     = "token"
	itemID    = "itemID"
	commentID = "commentID"
	userID    = "userID"
)

func main() {
	scopes := []string{oauth.GetAllScope()}
	ac := oauth.Config{
		ClientKey:    "ClientKey",
		ClientSecret: "ClientSecret",
		RedirectURL:  "https://xxx.xx.xxx",
		Scopes:       strings.Join(scopes, ","),
		Endpoint: struct {
			AuthURL         string
			TokenURL        string
			RefreshTokenURL string
			ClientTokenURL  string
		}{
			AuthURL:         "https://open.douyin.com/platform/oauth/connect",
			TokenURL:        "https://open.douyin.com/oauth/access_token",
			RefreshTokenURL: "https://open.douyin.com/oauth/refresh_token",
			ClientTokenURL:  "https://open.douyin.com/oauth/client_token",
		},
	}
	se := service.NewService(&ac)
	r := gin.Default()
	g:= r.Group("douyin")
	g.GET("/oauth", func(c *gin.Context) {
		state := xid.New().String()
		url := ac.AuthCodeURL(state)
		c.SetCookie("oauth_state", state, 120, "/", "xxxx.xxx.xx", false, true)
		c.Redirect(http.StatusFound, url)
	})
	g.GET("/callback", func(c *gin.Context) {
		state := c.Query("state")
		cookieState, err := c.Cookie("oauth_state")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if cookieState != state {
			fmt.Println("state error:", cookieState, state)
			return
		}
		code := c.Query("code")
		userToken, err := ac.Token(code)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(userToken)
	})
	g.GET("/video/comment/reply", func(c *gin.Context) {
		resource.VideoCommentReply(token, openid, resource.ComentReq{
			CommentID: commentID,
			ItemID:    itemID,
			Content:   "我也是这么认为",
		})

	})
	g.GET("/sendmessage", func(c *gin.Context) {
		im.SendMessage(token, openid, im.MessageReq{
			ToUserID:    userID,
			MessageType: "text",
			Content:     "douyin hello",
		})

	})
	g.GET("/userinfo", func(c *gin.Context) {
		uinfo, err := user.GetUserInfo(token, openid)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, uinfo)
	})
	g.GET("/following", func(c *gin.Context) {
		fans, err := user.GetFollowing(token, openid, 0, 20)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, fans)
	})
	g.GET("/fans", func(c *gin.Context) {
		fans, err := user.GetFans(token, openid, 0, 20)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, fans)
	})
	g.GET("/video", func(c *gin.Context) {
		data, err := ioutil.ReadFile("./1.mp4")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resource.Video2DouYin(token, openid, data, "test", nil)
	})
	g.GET("/videos", func(c *gin.Context) {
		vs, err := resource.GetVideoList(token, openid, 0, 20)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, vs)
	})
	g.GET("/video/comments", func(c *gin.Context) {
		vs, err := resource.GetVideoCommentList(token, openid, 0, 20, itemID)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(http.StatusOK, vs)
	})
	g.GET("/image", func(c *gin.Context) {
		data, err := ioutil.ReadFile("./1.jpeg")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		resource.Image2DouYin(token, openid, data, "test", nil)
	})
	g.POST("/webhook", func(c *gin.Context) {
        //授权webhook自动处理
		se.CallBack(c.Request, c.Writer)
	})
    
    // 增加处理的webhook类型
	se.AddHandle(service.ReceiveMsgEvent, func(req service.HookReq) []byte {
		data := req.ReceiveMsgEvent()
		fmt.Println(data.Content.Text)
		return nil
	})

	r.Run(":8000")
}

```
