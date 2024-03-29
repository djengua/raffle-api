package core

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Participant struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type RaffleParticipant struct {
	Participant `json:"participants" bson:"participants"`
	Tickets     []string `json:"tickets" bson:"tickets"`
}

type Raffle struct {
	ID           primitive.ObjectID  `json:"id,omitempty" bson:"_id"`
	Name         string              `json:"name" bson:"name"`
	Prize        string              `json:"prize" bson:"prize"`
	MaxTickets   int                 `json:"max_tickets" bson:"max_tickets"`
	Participants []RaffleParticipant `json:"participants,omitempty" bson:"participants"`
	FirstTaken   bool                `json:"first_taken" bson:"first_taken"`
	Turns        int                 `json:"turns" bson:"turns"`
	Tickets      []string            `json:"tickets" bson:"tickets"`
	Log          []string            `json:"log,omitempty" bson:"log"`
	CreatedAt    time.Time           `json:"createdAt" bson:"created_at"`
	UpdatedAt    time.Time           `json:"updatedAt" bson:"updated_at"`
	Open         bool                `json:"open" bson:"open"`
	Winner       RaffleParticipant   `json:"winner,omitempty" bson:"winner"`
	TicketWinner string              `json:"ticket_winner,omitempty" bson:"ticket_winner"`
}

func (r *Raffle) ToString() string {
	return fmt.Sprintf("id: %s, name: %s, participants: %d", r.ID, r.Name, len(r.Participants))
}

type TandaParticipant struct {
	Participant `json:"participants" bson:"participants"`
	Numbers     []int `json:"numbers" bson:"numbers"`
}

type Tanda struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Participants []TandaParticipant `json:"participants" bson:"participants"`
	Name         string             `json:"name" bson:"name"`
	NumbersTotal int                `json:"numbers_total" bson:"numbers_total"`
}
