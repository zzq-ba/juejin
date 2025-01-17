package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	Response
	ArticleList []Article `json:"article_list,omitempty"`
}

type JumpResponse struct {
	Response
	ArticleResponse Article `json:"article,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")

	dbInit()
	defer db.Close()
	var userLogin []dbUser
	//查询登录用户信息
	db.Select(&userLogin, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where token=?", token)
	var articleList []Article
	//获取文章列表
	rows, _ := db.Query("select ID, AuthorID, PlayUrl, CoverUrl, FavoriteCount, CommentCount, IsFavorite, Title, Introduction from Article where ID>?", 0)
	//填充文章列表
	if rows != nil {
		for rows.Next() {
			var article dbArticle
			rows.Scan(&article.ID, &article.AuthorID, &article.Url, &article.FavoriteCount, &article.CommentCount, &article.IsFavorite, &article.Title, &article.PublishTime, &article.Text)
			//获取用户信息
			var users []dbUser
			db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", article.AuthorID)
			//若用户已登录，则判断是否已关注文章作者，否则默认未关注
			if userLogin != nil {
				var follow []dbFollower
				db.Select(&follow, "select IsFollow from FollowList where FollowerID=? and UserID=?", userLogin[0].ID, users[0].ID)
				if follow != nil {
					users[0].IsFollow = true
				} else {
					users[0].IsFollow = false
				}
			} else {
				users[0].IsFollow = false
			}
			var user = User{
				Id:            users[0].ID,
				Name:          users[0].Name,
				FollowCount:   users[0].FollowCount,
				FollowerCount: users[0].FollowerCount,
				IsFollow:      users[0].IsFollow,
			}
			articleList = append([]Article{
				{
					Id:     article.ID,
					Author: user,
					Url:    article.Url,
					//PublishTime:   article.PublishTime,
					FavoriteCount: article.FavoriteCount,
					CommentCount:  article.CommentCount,
					IsFavorite:    article.IsFavorite,
					Text:          "",
					Title:         article.Title,
					Introduction:  article.Introduction,
				},
			}, articleList...)
		}
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:    Response{StatusCode: 0, StatusMsg: ""},
		ArticleList: articleList,
	})
}

func Jump(c *gin.Context) {
	articleID := c.Query("id")

	dbInit()
	defer db.Close()
	//获取文章信息，若不存在直接返回
	var articles []dbArticle
	var article Article
	db.Select(&articles, "select ID, AuthorID, Url, FavoriteCount, CommentCount, IsFavorite, Title, Text, Introduction from Article where ID=?", articleID)
	if articles == nil {
		c.JSON(http.StatusOK, JumpResponse{
			Response:        Response{StatusCode: 1, StatusMsg: "article not found"},
			ArticleResponse: article,
		})
	} else {
		var users []dbUser
		db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", articles[0].AuthorID)
		var user = User{
			Id:            users[0].ID,
			Name:          users[0].Name,
			FollowCount:   users[0].FollowCount,
			FollowerCount: users[0].FollowerCount,
			IsFollow:      users[0].IsFollow,
		}
		article = Article{
			Id:            articles[0].ID,
			Author:        user,
			Url:           articles[0].Url,
			FavoriteCount: articles[0].FavoriteCount,
			CommentCount:  articles[0].CommentCount,
			IsFavorite:    articles[0].IsFavorite,
			Text:          articles[0].Text,
			Introduction:  articles[0].Introduction,
			Title:         articles[0].Title,
		}
		c.JSON(http.StatusOK, JumpResponse{
			Response:        Response{StatusCode: 0, StatusMsg: ""},
			ArticleResponse: article,
		})
	}
}
