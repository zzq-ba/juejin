package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	userID := c.Query("user_id")
	actionType := c.Query("action_type")
	articleID := c.Query("article_id")

	dbInit()
	defer db.Close()
	var users []dbUser
	var articles []dbArticle
	//查询
	db.Select(&users, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", userID)
	db.Select(&articles, "select ID, AuthorID, Url, FavoriteCount, CommentCount, IsFavorite, Title, PublishTime from Article where ID=?", articleID)

	if users != nil {
		if actionType == "1" {
			text := c.Query("comment_text")
			newID := makeId()
			_, month, day := time.Now().Date()
			createData := string(month) + "-" + string(day)

			db.Exec("update Article set CommentCount=? where ID=?", articles[0].CommentCount+1, articleID)
			db.Exec("insert into Comment(CommentText, CreateDate, ID, UserID, ArticleID)value(?, ?, ?, ?, ?)", text, createData, newID, userID, articleID)

			var user = User{
				Id:            users[0].ID,
				Name:          users[0].Name,
				FollowCount:   users[0].FollowCount,
				FollowerCount: users[0].FollowerCount,
				IsFollow:      users[0].IsFollow,
			}

			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "Comment success"},
				Comment: Comment{
					Id:         newID,
					User:       user,
					Content:    text,
					CreateDate: string(month) + "-" + string(day),
				}})
			return
		} else if actionType == "2" {
			commentID := c.Query("comment_id")
			db.Exec("update Article set CommentCount=? where ID=?", articles[0].CommentCount-1, articleID)
			db.Exec("delete from Comment where ID=?", commentID)

			c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Delete comment success"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	articleID := c.Query("article_id")
	dbInit()
	defer db.Close()
	//从数据库查询该视频的评论列表
	var dbComments []dbComment
	db.Select(&dbComments, "select ID, UserID, CommentText, CreateDate from Comment where AticleID=?", articleID)
	var comments []Comment
	//填充返回的评论列表
	for _, comment := range dbComments {
		//查询评论用户
		var user []dbUser
		db.Select(&user, "select ID, Name, FollowCount, FollowerCount, IsFollow from User where ID=?", comment.UserID)
		comments = append(comments, Comment{
			Id: comment.ID,
			User: User{
				Id:            user[0].ID,
				Name:          user[0].Name,
				FollowCount:   user[0].FollowCount,
				FollowerCount: user[0].FollowerCount,
				IsFollow:      user[0].IsFollow,
			},
			Content:    comment.CommentText,
			CreateDate: comment.CreateDate,
		})
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0, StatusMsg: ""},
		CommentList: comments,
	})
}
