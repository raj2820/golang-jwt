package controllers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "rajauth/helpers"
	"rajauth/models"
	"rajauth/database"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client,"user")
var validate = validator.New()

func HashPassword(){

} 

func VerifyPassword(){

}

func SignUp(){

}

func Login(){

}

func GetUsers(){

}

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")
		err := helper.MatchUserTypeToUid(c, userId) 
		err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
	}
}