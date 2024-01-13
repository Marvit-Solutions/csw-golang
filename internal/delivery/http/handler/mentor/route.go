package mentor

import "github.com/gin-gonic/gin"

func (mentorHandler *MentorHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/mentor")
	authGroup.GET("/topmentors", mentorHandler.ListThreeMentors)
}
