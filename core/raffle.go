package core

import (
	"errors"
	"fmt"
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/exp/slices"
)

func (r *Raffle) AddParticipant(newParticipant Participant) error {
	if !r.Open {
		return errors.New("the raffle is closed")
	}

	if len(r.Participants) == 0 {
		r.Participants = []Participant{}
	}

	if newParticipant.ID == primitive.NilObjectID {
		newParticipant.ID = primitive.NewObjectID()
	}

	for i := 0; i < len(r.Participants); i++ {
		if r.Participants[i].Name == newParticipant.Name || r.Participants[i].ID == newParticipant.ID {
			return errors.New("participant duplicated")
		}
	}

	r.Participants = append(r.Participants, newParticipant)
	return nil
}

func (r *Raffle) randomTicket() string {
	var randomTicket string // := fmt.Sprintf("%06d", rand.Intn(r.MaxTickets))
	for {
		randomTicket = fmt.Sprintf("%06d", rand.Intn(r.MaxTickets))
		exists := r.existTicket(randomTicket)
		if !exists {
			break
		}
	}
	return randomTicket
}

func (r *Raffle) existTicket(ticket string) bool {

	for i := 0; i < len(r.Participants); i++ {
		if slices.Contains(r.Participants[i].Tickets, ticket) {
			return true
		}
	}
	return false
}

func (r *Raffle) AddTicketToParticipant(ticketSelected string, isRandom bool, participant string) (bool, error) {
	if !r.Open {
		return false, errors.New("the raffle is closed")
	}

	if ticketSelected == "" && !isRandom {
		return false, errors.New("cannot add ticket void")
	}

	if len(r.Tickets) >= r.MaxTickets {
		return false, errors.New("cannot add more tickets")
	}

	for i := 0; i < len(r.Participants); i++ {
		if r.Participants[i].Name == participant {
			randomTicket := ticketSelected
			if isRandom {
				randomTicket = r.randomTicket()
			}
			if !r.existTicket(randomTicket) {
				r.Participants[i].Tickets = append(r.Participants[i].Tickets, randomTicket)
				r.Prepare()
				return true, nil
			}
		}
	}

	return false, fmt.Errorf("not found participant %s or ticket is duplicated", participant)
}

func (r *Raffle) PrintParticipants() {
	fmt.Println(" == Participants == ")
	for _, participant := range r.Participants {
		fmt.Printf(" - %s [%d tickets (%s)] \n", participant.Name, len(participant.Tickets), participant.Tickets)
	}
}

func (r *Raffle) Prepare() {
	r.Tickets = []string{}
	// fmt.Println("preparing tickets of all participants...")
	for _, p := range r.Participants {
		r.Tickets = append(r.Tickets, p.Tickets...)
	}
	// fmt.Printf("total of tickets: %d \n", len(r.Tickets))
}

func (r *Raffle) DeleteParticipant(name string) error {
	if !r.Open {
		return errors.New("the raffle is closed")
	}
	index := -1

	for i := 0; i < len(r.Participants); i++ {
		if r.Participants[i].Name == name {
			index = i
			break
		}
	}

	newSlice, _, err := deleteParticipantAtIndex(r.Participants, index)
	if err != nil {
		return err
	}
	r.Participants = newSlice
	return nil
}

// Get participant of slice of participants
func deleteTicketAtIndex(tickets []string, index int) ([]string, string, error) {
	if index > len(tickets) {
		return nil, "", errors.New("index is greater than to total of tickets")
	}
	p := tickets[index]
	return append(tickets[:index], tickets[index+1:]...), p, nil
}

func deleteParticipantAtIndex(participants []Participant, index int) ([]Participant, Participant, error) {
	if index > len(participants) {
		return nil, Participant{}, errors.New("index is greater than to total of participants")
	}

	if index < 0 {
		return nil, Participant{}, errors.New("not found participant")
	}

	p := participants[index]
	return append(participants[:index], participants[index+1:]...), p, nil
}

func (r *Raffle) SelectWinner() error {
	r.Log = append(r.Log, "preparing all tickets.")
	r.Prepare()
	r.Log = append(r.Log, fmt.Sprintf("total of tickets: %d", len(r.Tickets)))
	if !r.Open {
		return errors.New("the raffle is closed")
	}

	r.Open = false

	if r.Turns > len(r.Tickets) {
		fmt.Println("The turns is greater than total of tickets, redim to -1.")
		r.Turns = len(r.Tickets)
	}

	for turn := 0; turn < int(r.Turns); turn++ {
		r.Log = append(r.Log, " Stirring and taking a ticket. ")
		// Obtenemos random de los tickets
		i := rand.Intn(len(r.Tickets))
		newSlice, ticketTaken, err := deleteTicketAtIndex(r.Tickets, i)

		if err != nil {
			panic(err)
		}

		if turn == int(r.Turns)-1 {
			fmt.Printf(" The ticket winner is: '%s' \n", ticketTaken)
			r.Log = append(r.Log, fmt.Sprintf(" The ticket winner is: '%s'", ticketTaken))

			for i := 0; i < len(r.Participants); i++ {
				if slices.Contains(r.Participants[i].Tickets, ticketTaken) {
					r.Winner = r.Participants[i]
					r.TicketWinner = ticketTaken
					fmt.Printf(" The participant winner is: '%s' \n", r.Participants[i].Name)
					r.Log = append(r.Log, fmt.Sprintf(" The participant winner is: '%s'", r.Participants[i].Name))
					break
				}
			}
		} else {
			fmt.Printf(" Discard: %s \n", ticketTaken)
			r.Log = append(r.Log, fmt.Sprintf(" Discard: %s", ticketTaken))
		}
		r.Tickets = newSlice
	}

	fmt.Println("  Non-winning Tickets ")
	fmt.Println(r.Tickets)
	return nil
}
