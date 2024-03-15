package api

import (
	"net/http"

	jwtauth "token"

	"github.com/gin-gonic/gin"
)

func (server *Server) SetupRoutes() {
	// Load HTML templates
	server.Router.LoadHTMLGlob("./templates/*.html")

	// Serve static files
	server.Router.Static("/public", "./public")

	// Public routes
	server.Router.GET("/lienhe", GetHandler("lienhe"))
	server.Router.GET("/register", GetHandler("register"))
	server.Router.GET("/", GetHandler("home"))
	server.Router.GET("/thankyou", GetHandler("thankyou"))
	server.Router.GET("/login", GetHandler("login"))

	communityRoutes := server.Router.Group("r/community")
	communityRoutes.GET("/", GetHandler("community"))
	communityRoutes.GET("/:community/:post_id/:post", server.handlerWrapper(server.CommunityHandler))

	// Route grouping for authenticated routes
	authRoutes := server.Router.Group("/")
	authRoutes.Use(
		// func(c *gin.Context) {
		// 	fmt.Println("\n Entering Verifier middleware")
		// },

		jwtauth.Verifier(server.TokenAuthRS256),
		// func(c *gin.Context) {
		// 	fmt.Println("\n <<< After Verifier")
		// },

		// jwtauth.VerifyPaseto(server.Pv4),
		// func(c *gin.Context) {
		// 	fmt.Println("\n <<< After VerifyPaseto")
		// },

		// jwtauth.VerifyPaseto(*http.Request),

		jwtauth.Authenticator(server.TokenAuthRS256),
		// func(c *gin.Context) {
		// 	fmt.Println("\n <<< After Authenticator")
		// },

		roleMiddleware(),
		// func(c *gin.Context) {
		// 	fmt.Println("\n <<< After roleMiddleware")
		// },
	)
	{
		authRoutes.GET("/:role/:token", UserHandler)
	}

	// API routes
	apiRoutes := server.Router.Group("/api/v1")
	apiRoutes.POST("/form", server.handlerWrapper(server.formHandler))
	// apiRoutes.Use(
	// 	jwtauth.Verifier(server.TokenAuthRS256),
	// 	jwtauth.Authenticator(server.TokenAuthRS256),
	// )
	{
		apiRoutes.POST("/register", server.handlerWrapper(server.registerHandler))
		apiRoutes.GET("/verify_email", server.handlerWrapper(server.VerifyEmailHandler))
		apiRoutes.POST("/login", server.handlerWrapper(server.loginHandler))
		apiRoutes.POST("/post", server.handlerWrapper(server.CreatePost))
		apiRoutes.POST("/logout", server.handlerWrapper(server.logoutHandler))
		apiRoutes.GET("/tintuc", server.handleNews)

	}

	// No route handler
	server.Router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
}
