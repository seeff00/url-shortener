package http

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) Middleware(ginCtx *gin.Context) {
	ginCtx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ginCtx.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
	ginCtx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	ginCtx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	ginCtx.Next()
}
