package api

import (
	"fmt"
	"net/http"

	db "sqlc"
	token "token"
	util "util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type postsResponse struct {
	Post      []db.Post
	Community []db.Community
}

func (server *Server) postHandler(c *gin.Context) {
	// Retrieve posts from the database
	posts, err := server.DbHandler.GetAllPosts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("\n postHandler posts: ", posts)

	// Collect community IDs from posts
	var communityIDs []uuid.UUID
	for _, post := range posts {
		communityIDs = append(communityIDs, post.CommunityID)
	}
	fmt.Println("\n postHandler communityIDs: ", communityIDs)

	// Retrieve communities for each community ID
	var communities []db.Community
	for _, id := range communityIDs {
		community, err := server.DbHandler.GetCommunitybyID(c.Request.Context(), id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		// Append each community individually
		communities = append(communities, community...)
	}
	fmt.Println("\n postHandler communities: ", communities)

	// Construct the response
	response := postsResponse{
		Post:      posts,
		Community: communities,
	}
	fmt.Println("\n postHandler response: ", response)

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
	fmt.Println("\n CreatePost c before Parse: ", c)
	fmt.Println("\n CreatePost msg before Parse: ", msg)
	fmt.Println("\n CreatePost claim before Parse: ", claim)
	fmt.Println("\n CreatePost claim['ID'].(string) before Parse: ", claim["ID"].(string))

	// Parse the string value into a UUID
	msg.UserID, err = uuid.Parse(claim["ID"].(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("\n CreatePost msg before GetCommunityIDbyName: ", msg)

	msg.CommunityID, err = server.DbHandler.GetCommunityIDbyName(c, msg.CommunityName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("CreatePost msg before CreatePostTx: ", msg)

	// msg.CommunityID = communityResult[0].ID

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
