package coreapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/djengua/raffle-api/core"
	"github.com/djengua/raffle-api/util"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRaffleRequest struct {
	Name       string `json:"name" binding:"required"`
	Prize      string `json:"prize" binding:"required"`
	FirstTaken bool   `json:"first_taken"`
	Turns      int    `json:"turns" binding:"required"`
	MaxTickets int    `json:"max_tickets"`
}

func (s *Server) hello(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  "Hello World",
	})
}

func (s *Server) createRaffle(ctx *gin.Context) {
	var req CreateRaffleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle := core.Raffle{
		ID:         primitive.NewObjectID(),
		Name:       req.Name,
		Prize:      req.Prize,
		FirstTaken: req.FirstTaken,
		Turns:      req.Turns,
		MaxTickets: req.MaxTickets,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Open:       true,
	}

	_, err := s.database.Collection("raffles").InsertOne(ctx, raffle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

type GetRaffleByIdRequest struct {
	Id string `uri:"id" binding:"required"`
}

func (s *Server) fetchById(ctx *gin.Context, id string) (core.Raffle, error) {
	var raffle core.Raffle

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return raffle, err
	}

	filter := bson.M{"_id": objId}
	err = s.database.Collection("raffles").FindOne(ctx, filter).Decode(&raffle)
	if err != nil {
		return raffle, err
	}

	return raffle, nil
}

func (s *Server) fetchAll(ctx *gin.Context) ([]*core.Raffle, error) {
	var raffles []*core.Raffle

	filter := bson.M{}
	cur, err := s.database.Collection("raffles").Find(ctx, filter)
	if err != nil {
		return raffles, err
	}

	for cur.Next(ctx) {
		var r core.Raffle
		err := cur.Decode(&r)
		if err != nil {
			return raffles, err
		}
		raffles = append(raffles, &r)
	}

	return raffles, nil
}

func (s *Server) getRaffleById(ctx *gin.Context) {
	var req GetRaffleByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle, err := s.fetchById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

func (s *Server) getAllRaffle(ctx *gin.Context) {
	raffle, err := s.fetchAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

func (s *Server) update(ctx *gin.Context, raffle *core.Raffle) error {
	raffle.UpdatedAt = time.Now()

	objId, err := primitive.ObjectIDFromHex(raffle.ID.Hex())
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: objId}}

	doc, err := util.StructToDocument(raffle)
	if err != nil {
		return err
	}

	update := bson.D{primitive.E{Key: "$set", Value: doc}}

	result, err := s.database.Collection("raffles").UpdateOne(ctx, filter, update)
	fmt.Printf(" Modified: %d \n", result.ModifiedCount)
	if err != nil {
		return err
	}

	return nil
}

type AddParticipantRaffleRequest struct {
	RaffleID    string           `json:"raffle_id" binding:"required"`
	Participant core.Participant `json:"participant" binding:"required"`
}

func (s *Server) addParticipant(ctx *gin.Context) {
	var req AddParticipantRaffleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle, err := s.fetchById(ctx, req.RaffleID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle.AddParticipant(req.Participant)
	err = s.update(ctx, &raffle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

type AddTicketParticipantRaffleRequest struct {
	RaffleID        string `json:"raffle_id" binding:"required"`
	ParticipantName string `json:"participant_name" binding:"required"`
	IsRandom        bool   `json:"is_random"`
	Ticket          string `json:"ticket"`
}

func (s *Server) addTicketToParticipant(ctx *gin.Context) {
	var req AddTicketParticipantRaffleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle, err := s.fetchById(ctx, req.RaffleID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = raffle.AddTicketToParticipant(req.Ticket, req.IsRandom, req.ParticipantName)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = s.update(ctx, &raffle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

type DeleteParticipantRaffleRequest struct {
	RaffleID string `json:"raffle_id" binding:"required"`
	// Participant core.Participant `json:"participant" binding:"required"`
	ParticipantId string `json:"participant_id" binding:"required"`
}

func (s *Server) deleteParticipant(ctx *gin.Context) {
	var req DeleteParticipantRaffleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// TODO: Eliminar por id

	raffle, err := s.fetchById(ctx, req.RaffleID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = raffle.DeleteParticipant(req.ParticipantId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = s.update(ctx, &raffle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

func (s *Server) winner(ctx *gin.Context) {
	var req GetRaffleByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle, err := s.fetchById(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle.SelectWinner()

	err = s.update(ctx, &raffle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  raffle,
	})
}

type RaffleByIdRequest struct {
	RaffleId string `json:"raffle_id" binding:"required"`
}

func (s *Server) discardTicket(ctx *gin.Context) {
	var req RaffleByIdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	raffle, err := s.fetchById(ctx, req.RaffleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ticket, err := raffle.DiscardTicket()

	err = s.update(ctx, &raffle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"error": nil,
		"data":  ticket,
	})
}
