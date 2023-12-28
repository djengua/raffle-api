package core_test

import (
	"testing"

	"github.com/djengua/raffle-api/core"
	"github.com/stretchr/testify/require"
)

func TestRaffle(t *testing.T) {
	// Create new Raffle 3 turns, 5 boletos con 3 participantes , prize of 100

	raffle := core.Raffle{
		Name:       "Rifa de fin de año",
		Prize:      "100 pesos",
		Turns:      20,
		MaxTickets: 10,
		Open:       true,
	}

	// require.Equal(t, test.Expected, result)
	require.Equal(t, "Rifa de fin de año", raffle.Name)
	require.Equal(t, "100 pesos", raffle.Prize)
	require.Equal(t, 20, raffle.Turns)

	// Add participants
	participantOne := core.Participant{Name: "David J"}
	participantTwo := core.Participant{Name: "Eliott"}
	participantThree := core.Participant{Name: "Karen"}
	participantFour := core.Participant{Name: "David"}
	participantFive := core.Participant{Name: "David"}

	err := raffle.AddParticipant(participantOne)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("", false, participantOne.Name)
	require.Error(t, err)

	addedTicket, err := raffle.AddTicketToParticipant("", true, participantOne.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("A", false, participantOne.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("A", false, participantOne.Name)
	require.Equal(t, false, addedTicket)
	require.Error(t, err)

	err = raffle.AddParticipant(participantTwo)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantTwo.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("A-00001", false, participantTwo.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("A-00002", false, participantTwo.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)

	err = raffle.AddParticipant(participantThree)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantThree.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantThree.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("B-0012", false, participantThree.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)

	err = raffle.AddParticipant(participantFour)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantFour.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantFour.Name)
	require.Equal(t, true, addedTicket)
	require.NoError(t, err)
	addedTicket, err = raffle.AddTicketToParticipant("", true, participantFour.Name)
	require.Equal(t, false, addedTicket)
	require.Error(t, err)

	err = raffle.AddParticipant(participantFive)
	require.Error(t, err)

	raffle.PrintParticipants()
	err = raffle.DeleteParticipant("David F")
	require.Error(t, err)
	_ = raffle.DeleteParticipant("David J")
	require.Equal(t, 3, len(raffle.Participants))
	raffle.PrintParticipants()

	// raffle.Prepare()
	// require.Equal(t, 8, len(raffle.Tickets))

	// // passing 3 turns,
	require.Equal(t, 3, len(raffle.Participants))

	err = raffle.SelectWinner()
	require.NoError(t, err)
	err = raffle.SelectWinner()
	require.Error(t, err)

	err = raffle.DeleteParticipant("David")
	require.Error(t, err)

	addedTicket, err = raffle.AddTicketToParticipant("", true, participantFour.Name)
	require.Equal(t, false, addedTicket)
	require.Error(t, err)
}
