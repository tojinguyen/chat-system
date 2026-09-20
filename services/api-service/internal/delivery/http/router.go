package http

import (
	"net/http"

	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Auth         *AuthHandler
	User         *UserHandler
	Friend       *FriendHandler
	Conversation *ConversationHandler
}

func SetupRouter(authUsecase usecase.AuthUsecase, handlers *Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// CORS nhẹ
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 1. Auth routes (public)
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", handlers.Auth.Register)
		authGroup.POST("/login", handlers.Auth.Login)
	}

	// Protected routes
	api := r.Group("/")
	api.Use(middleware.AuthMiddleware(authUsecase))
	{
		// 2. Users routes
		users := api.Group("/users")
		{
			users.GET("/me", handlers.User.GetProfile)
			users.PATCH("/me", handlers.User.UpdateProfile)
			users.GET("/search", handlers.User.SearchUsers)
			users.GET("/:id", handlers.User.GetUserByID)
		}

		// 3. Friends routes
		friends := api.Group("/friends")
		{
			friends.GET("", handlers.Friend.GetFriendsList)
			friends.GET("/requests", handlers.Friend.GetPendingRequests)
			friends.POST("/request/:targetUserId", handlers.Friend.SendFriendRequest)
			friends.PATCH("/requests/:requestId/accept", handlers.Friend.AcceptFriendRequest)
			friends.PATCH("/requests/:requestId/decline", handlers.Friend.DeclineFriendRequest)
			friends.POST("/block/:targetUserId", handlers.Friend.BlockUser)
			friends.DELETE("/:relationshipId", handlers.Friend.Unfriend)
		}

		// 4. Conversations routes
		convos := api.Group("/conversations")
		{
			convos.GET("", handlers.Conversation.GetMyConversations)
			convos.POST("/direct", handlers.Conversation.GetOrCreateDirect)
			convos.POST("/group", handlers.Conversation.CreateGroup)
			convos.GET("/:id", handlers.Conversation.GetConversationByID)
			convos.POST("/:id/members", handlers.Conversation.AddMember)
			convos.DELETE("/:id/members/:targetUserId", handlers.Conversation.RemoveMember)
			convos.POST("/:id/read", handlers.Conversation.UpdateReadReceipt)
		}
	}

	return r
}
