package api

import (
	"net/http"

	db "sqlc"
	token "token"
	util "util"

	"github.com/gin-gonic/gin"
)

type postsResponse struct {
	db.Post
	Community struct {
		db.Community
		Rule []db.Rule
	}
	Comments []CommentWithReplies
	Account  []db.GetAccountbyIDRow
}

type CommentWithReplies struct {
	db.Comment
	Replies []db.Reply
}

func (server *Server) postHandler(c *gin.Context) {
	posts, err := server.DbHandler.GetAllPosts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []postsResponse
	for _, post := range posts {
		// Retrieve community details for the post
		community, err := server.DbHandler.GetCommunitybyID(c.Request.Context(), post.CommunityID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Retrieve rules for the community
		rule, err := server.DbHandler.GetRuleFromCommunity(c.Request.Context(), community.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Retrieve account details for the post user
		account, err := server.DbHandler.GetAccountbyID(c.Request.Context(), post.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Retrieve comments for the post
		comments, err := server.DbHandler.GetCommentFromPost(c.Request.Context(), post.ID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		// Group comments and their replies together
		var commentDetails []CommentWithReplies
		for _, comment := range comments {
			replies, err := server.DbHandler.GetReplyFromComment(c.Request.Context(), comment.ID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
			commentDetails = append(commentDetails, CommentWithReplies{
				Comment: comment,
				Replies: replies,
			})
		}

		// Append post details to the response
		response = append(response, postsResponse{
			Post: post,
			Community: struct {
				db.Community
				Rule []db.Rule
			}{Community: community, Rule: rule},
			Comments: commentDetails,
			Account:  account,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (server *Server) CreatePost(c *gin.Context) {
	msg, err := postForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// msg.UserID = c.GetHeader("id")
	claim, err := token.GetClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg.UserID, err = server.DbHandler.GetAccountIDbyUsername(c, claim["Username"].(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg.CommunityID, err = server.DbHandler.GetCommunityIDbyName(c, msg.CommunityName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// fmt.Println("CreatePost msg before CreatePostTx: ", msg)

	if submit, err := server.DbHandler.CreatePostTx(c.Request.Context(), *msg); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		msg.ID = submit.Post.ID
	}

	c.JSON(http.StatusOK, msg)
}

func postForm(r *http.Request) (*db.CreatePostTxParams, error) {
	msg := &db.CreatePostTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
