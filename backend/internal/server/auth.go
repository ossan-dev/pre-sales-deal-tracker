package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ekefan/deal-tracker/internal/token"
	"github.com/gin-gonic/gin"
)

const (
	authHeaderKey = "authorization"
	authHeaderType = "bearer"
	authPayloadKey = "authorization_payload"
)

// authMiddleware checks if the user authorized to access resource
// it checks the request header for the authorization field
// validates the authorization used, and on success gets the tokens payload
// sets it in the current server context instance
func authMiddleware(maker token.TokenMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Check header for authorization header
		authHeader := ctx.GetHeader(authHeaderKey); 
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("no auth header passed")))
			return
		}

		//check for format validation
		fields := strings.Fields(authHeader)
		if len(fields) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("invalid or missing bearer token")))
			return
		}

		//check authorization type
		authType := fields[0]
		if strings.ToLower(authType) != authHeaderType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(fmt.Errorf("authorization type not supported")))
			return
		}
		
		//token verification
		payload, err := maker.VerifyToken(fields[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		//if verification is successful pass to the next handler func
		ctx.Set(authPayloadKey, payload)
		ctx.Next()
	}
}