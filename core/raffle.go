package core

import (
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/exp/slices"
)

func (r *Raffle) AddParticipant(newParticipant Participant) {
	if len(r.Participants) == 0 {
		r.Participants = []Participant{}
	}
	r.Participants = append(r.Participants, newParticipant)
}

func (r *Raffle) AddTicketToParticipant(ticketSelected string, isRandom bool, participant string) error {
	if ticketSelected == "" && !isRandom {
		return errors.New("cannot add ticket void")
	}

	randomTicket := ticketSelected
	if isRandom {
		randomTicket = fmt.Sprintf("%04d", rand.Intn(r.MaxTickets))
	}

	// Add ticket to participant
	for i := 0; i < len(r.Participants); i++ {
		if len(r.Participants[i].Tickets) == 0 {
			r.Participants[i].Tickets = []string{}
		}
		if r.Participants[i].Name == participant {
			r.Participants[i].Tickets = append(r.Participants[i].Tickets, randomTicket)
			return nil
		}
	}
	return fmt.Errorf("not found participant %s", participant)
}

func (r *Raffle) PrintParticipants() {
	fmt.Println("== Participants ==")
	for _, participant := range r.Participants {
		fmt.Printf(" - %s [%d tickets (%s)] \n", participant.Name, len(participant.Tickets), participant.Tickets)
	}
}

func (r *Raffle) Prepare() {
	r.Tickets = []string{}
	fmt.Println("preparing tickets of all participants...")
	r.Log = append(r.Log, "preparing tickets of all participants...")
	for _, p := range r.Participants {
		r.Tickets = append(r.Tickets, p.Tickets...)
	}
	fmt.Printf("total of tickets: %d \n", len(r.Tickets))
	r.Log = append(r.Log, fmt.Sprintf("total of tickets: %d \n", len(r.Tickets)))
	fmt.Println("  Tickets ")
	fmt.Println(r.Tickets)
}

func (r *Raffle) DeleteParticipant(name string) error {
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

func (r *Raffle) SelectWinner() {
	r.Prepare()

	if r.Turns > len(r.Tickets) {
		fmt.Println("The turns is greater than total of tickets, redim to -1.")
		r.Turns = len(r.Tickets)
	}

	for turn := 0; turn < int(r.Turns); turn++ {
		// Obtenemos random de los tickets
		i := rand.Intn(len(r.Tickets))
		newSlice, ticketTaken, err := deleteTicketAtIndex(r.Tickets, i)

		if err != nil {
			panic(err)
		}

		if turn == int(r.Turns)-1 {
			fmt.Printf(" The ticket winner is: '%s' \n", ticketTaken)
			r.Log = append(r.Log, fmt.Sprintf(" The ticket winner is: '%s' \n", ticketTaken))

			for i := 0; i < len(r.Participants); i++ {
				if slices.Contains(r.Participants[i].Tickets, ticketTaken) {
					fmt.Printf(" The participant winner is: '%s' \n", r.Participants[i].Name)
					r.Log = append(r.Log, fmt.Sprintf(" The participant winner is: '%s' \n", r.Participants[i].Name))
					break
				}
			}
		} else {
			fmt.Printf(" Discard: %s \n", ticketTaken)
			r.Log = append(r.Log, fmt.Sprintf(" Discard: %s \n", ticketTaken))
		}
		r.Tickets = newSlice
	}

	fmt.Println("  Non-winning Tickets ")
	fmt.Println(r.Tickets)
}
