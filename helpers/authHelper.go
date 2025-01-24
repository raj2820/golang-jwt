package helpers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

func MatchUserTypeToUid(c *gin.Context, userId string)