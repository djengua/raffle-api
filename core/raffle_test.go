package core_test

import (
	"testing"

	"github.com/djengua/rifa-api/core"
	"github.com/stretchr/testify/require"
)

func TestRaffle(t *testing.T) {
	// Create new Raffle 3 turns, 5 boletos con 3 participantes , prize of 100

	raffle := core.Raffle{
		Name:       "Rifa de fin de año",
		Prize:      "100 pesos",
		Turns:      20,
		MaxTickets: 100,
	}

	// require.Equal(t, test.Expected, result)
	require.Equal(t, "Rifa de fin de año", raffle.Name)
	require.Equal(t, "100 pesos", raffle.Prize)
	require.Equal(t, 20, raffle.Turns)

	// Add participants

	participantOne := core.Participant{Name: "David J"}
	// participantOne.Tickets = []string{"23", "18"}
	participantTwo := core.Participant{Name: "Eliott"}
	// participantTwo.Tickets = []string{"10", "100", "50"}
	participantThree := core.Participant{Name: "Karen"}
	// participantThree.Tickets = []string{"1", "150", "8"}
	participantFour := core.Participant{Name: "David"}
	// participantFour.Tickets = []string{"189", "3", "66"}

	raffle.AddParticipant(participantOne)
	raffle.AddTicketToParticipant("", true, participantOne.Name)
	raffle.AddTicketToParticipant("A", true, participantOne.Name)
	raffle.AddParticipant(participantTwo)
	raffle.AddTicketToParticipant("", true, participantTwo.Name)
	raffle.AddTicketToParticipant("A-0001", false, participantTwo.Name)
	raffle.AddTicketToParticipant("A-0002", false, participantTwo.Name)
	raffle.AddParticipant(participantThree)
	raffle.AddTicketToParticipant("", true, participantThree.Name)
	raffle.AddTicketToParticipant("", true, participantThree.Name)
	raffle.AddTicketToParticipant("B-0012", false, participantThree.Name)
	raffle.AddParticipant(participantFour)
	raffle.AddTicketToParticipant("", true, participantFour.Name)
	raffle.AddTicketToParticipant("", true, participantFour.Name)
	raffle.AddTicketToParticipant("", true, participantFour.Name)
	// 3 Participants
	raffle.PrintParticipants()
	require.Equal(t, 4, len(raffle.Participants))

	err := raffle.AddTicketToParticipant("", true, "Juanito")
	require.Error(t, err)

	err = raffle.DeleteParticipant("David F")
	require.Error(t, err)
	_ = raffle.DeleteParticipant("David")
	require.Equal(t, 3, len(raffle.Participants))
	raffle.PrintParticipants()

	raffle.Prepare()
	// 7 boletos
	require.Equal(t, 8, len(raffle.Tickets))

	// passing 3 turns,
	require.Equal(t, 3, len(raffle.Participants))

	raffle.SelectWinner()

}
