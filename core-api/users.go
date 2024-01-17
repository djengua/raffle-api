package coreapi

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/djengua/raffle-api/core"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const USERS_COLLECTION = "users"

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func (s *Server) createUsers(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	role, _ := primitive.ObjectIDFromHex(req.Role)

	user := core.User{
		ID:        primitive.NewObjectID(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		Role:      role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Active:    true,
	}

	_, err := s.database.Collection(USERS_COLLECTION).InsertOne(ctx, user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  user,
	})
}

func (s *Server) fetchAllUsers(ctx *gin.Context) ([]*core.User, error) {
	var users []*core.User

	filter := bson.M{}
	cur, err := s.database.Collection(USERS_COLLECTION).Find(ctx, filter)
	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var r core.User
		err := cur.Decode(&r)
		if err != nil {
			return users, err
		}
		users = append(users, &r)
	}

	return users, nil
}

func (s *Server) fetchUserByEmailPassword(ctx *gin.Context, email string, password string) (*core.User, error) {
	var user core.User

	filter := bson.M{"email": email, "password": password}

	err := s.database.Collection(USERS_COLLECTION).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return &user, err
	}

	return &user, nil
}

func (s *Server) fetchUserByEmail(ctx *gin.Context, email string) (*core.User, error) {
	var user core.User

	filter := bson.M{"email": email}

	err := s.database.Collection(USERS_COLLECTION).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		return &user, err
	}

	return &user, nil
}

func (s *Server) getAllUsers(ctx *gin.Context) {
	users, err := s.fetchAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  users,
	})

}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Server) loginUser(ctx *gin.Context) {

	var req LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(req)

	user, err := s.fetchUserByEmailPassword(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// err = util.CheckPassword(req.Password, user.Password)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	accessToken, err := s.tokenMaker.CreateToken(
		user.Email,
		s.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data": map[string]interface{}{
			"access_token": accessToken,
			"user":         user,
		},
	})
}

func (s *Server) refreshToken(ctx *gin.Context) {

	authHeader := ctx.GetHeader(authorizationHeaderKey)
	if len(authHeader) == 0 {
		err := errors.New("authorization header is not provided")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationTypeBearer {
		err := fmt.Errorf("unsupported authorization type %s", authType)
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	accessToken := fields[1]

	payload, err := s.tokenMaker.GetPayload(accessToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	user, err := s.fetchUserByEmail(ctx, payload.Username)
	if err != nil {
		// if err == sql.ErrNoRows {
		// 	ctx.JSON(http.StatusNotFound, errorResponse(fmt.Errorf("incorrect username or password")))
		// 	return
		// }
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	accessToken, err = s.tokenMaker.CreateToken(
		user.Email,
		s.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// resp := loginUserResponse{
	// 	AccessToken: accessToken,
	// 	User:        newUserResponse(user),
	// }

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data": map[string]interface{}{
			"access_token": accessToken,
			"user":         user,
		},
	})
}
