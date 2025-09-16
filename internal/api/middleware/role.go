package middleware

import (
	"dropx/pkg/constants"
	"dropx/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func OnlyRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleRaw, exists := c.Get(constants.TokenRole)
		if !exists {
			utils.WriteErrorResponse(c, fmt.Errorf("missing role in context"), http.StatusForbidden)
			c.Abort()
			return
		}

		role, ok := roleRaw.(string)
		if !ok {
			utils.WriteErrorResponse(c, fmt.Errorf("invalid role type"), http.StatusForbidden)
			c.Abort()
			return
		}

		for _, r := range allowedRoles {
			if strings.EqualFold(r, role) {
				c.Next()
				return
			}
		}

		utils.WriteErrorResponse(c, fmt.Errorf("unauthorized: role '%s' not allowed", role), http.StatusForbidden)
		c.Abort()
	}
}
