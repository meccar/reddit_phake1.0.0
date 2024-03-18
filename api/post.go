package api

import (
	"net/http"

	db "sqlc"
	token "token"
	util "util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Define a new struct to represent the combination of post and community details
type postsResponse struct {
	Post      db.Post
	Community []db.Community
	Account   []db.GetAccountbyIDRow
}

func (server *Server) postHandler(c *gin.Context) {
	// Retrieve posts from the database
	posts, err := server.DbHandler.GetAllPosts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create a map to store post ID as keys and their associated communities as values
	postCommunityMap := make(map[uuid.UUID][]db.Community)
	postAccountMap := make(map[uuid.UUID][]db.GetAccountbyIDRow)

	// Retrieve community details for each post's community ID
	for _, post := range posts {
		// Retrieve community details for the current post's community ID
		community, err := server.DbHandler.GetCommunitybyID(c.Request.Context(), post.CommunityID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		postCommunityMap[post.ID] = community

		account, err := server.DbHandler.GetAccountbyID(c.Request.Context(), post.UserID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		postAccountMap[post.ID] = account
	}

	// Construct the response by combining post and community details
	var response []postsResponse
	for _, post := range posts {
		community := postCommunityMap[post.ID]
		account := postAccountMap[post.ID]

		// Append post and associated community details to the response
		response = append(response, postsResponse{Post: post, Account: account, Community: community})
	}

	// Return the response to the client
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
