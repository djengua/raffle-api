package core_test

import (
	"fmt"
	"testing"

	"github.com/djengua/raffle-api/core"
	"github.com/djengua/raffle-api/util"
	"github.com/stretchr/testify/require"
)

func TestRaffleSelectWinner(t *testing.T) {
	// Create new Raffle 3 turns, 5 boletos con 3 participantes , prize of 10

	nameRaffle := util.RandomString(20)
	prizeRaffle := util.RandomString(10)

	raffle := core.Raffle{
		Name:       nameRaffle,
		Prize:      prizeRaffle,
		Turns:      3,
		MaxTickets: 10,
		Open:       true,
	}

	require.Equal(t, nameRaffle, raffle.Name)
	require.Equal(t, prizeRaffle, raffle.Prize)
	require.Equal(t, 3, raffle.Turns)

	// Add participants
	participantOne := core.RaffleParticipant{Participant: core.Participant{Name: "David J"}}
	participantTwo := core.RaffleParticipant{Participant: core.Participant{Name: "Eliott"}}
	participantThree := core.RaffleParticipant{Participant: core.Participant{Name: "Karen"}}
	participantFour := core.RaffleParticipant{Participant: core.Participant{Name: "David"}}
	participantFive := core.RaffleParticipant{Participant: core.Participant{Name: "David"}}

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
	require.Equal(t, 4, len(raffle.Participants))
	raffle.PrintParticipants()

	// raffle.Prepare()
	// require.Equal(t, 8, len(raffle.Tickets))

	// // passing 3 turns,
	require.Equal(t, 4, len(raffle.Participants))

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

func TestRaffleDiscardTicket(t *testing.T) {
	nameRaffle := util.RandomString(20)
	prizeRaffle := util.RandomString(10)

	raffle := core.Raffle{
		Name:       nameRaffle,
		Prize:      prizeRaffle,
		Turns:      3,
		MaxTickets: 10,
		Open:       true,
	}

	require.Equal(t, nameRaffle, raffle.Name)
	require.Equal(t, prizeRaffle, raffle.Prize)
	require.Equal(t, 3, raffle.Turns)

	raffle.ToString()

	participantOne := core.RaffleParticipant{}
	participantOne.Name = "Karen"
	// participantTwo := core.RaffleParticipant{}
	// participantTwo.Name = "Eliott"
	participantTwo := core.RaffleParticipant{Participant: core.Participant{Name: "Eliott"}}

	err := raffle.AddParticipant(participantOne)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("A-00001", false, participantOne.Name)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("A-00002", false, participantOne.Name)
	require.NoError(t, err)

	err = raffle.AddParticipant(participantTwo)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("A-00003", false, participantTwo.Name)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("A-00004", false, participantTwo.Name)
	require.NoError(t, err)
	_, err = raffle.AddTicketToParticipant("A-00005", false, participantTwo.Name)
	require.NoError(t, err)

	_, err = raffle.DiscardTicket()
	require.NoError(t, err)
	_, err = raffle.DiscardTicket()
	require.NoError(t, err)
	_, err = raffle.DiscardTicket()
	require.NoError(t, err)
	_, err = raffle.DiscardTicket()
	require.NoError(t, err)
	_, err = raffle.DiscardTicket()
	require.NoError(t, err)
}

func TestRaffleErrors(t *testing.T) {
	nameRaffle := util.RandomString(20)
	prizeRaffle := util.RandomString(10)

	raffle := core.Raffle{
		Name:       nameRaffle,
		Prize:      prizeRaffle,
		Turns:      10,
		MaxTickets: 10,
		Open:       false,
	}

	require.Equal(t, nameRaffle, raffle.Name)
	require.Equal(t, prizeRaffle, raffle.Prize)
	require.Equal(t, 10, raffle.Turns)

	participantOne := core.RaffleParticipant{}
	participantOne.Name = "Karen"

	err := raffle.AddParticipant(participantOne)
	require.Error(t, err)

	result, err := raffle.DiscardTicket()
	require.Error(t, err)
	fmt.Println(result)
	err = raffle.SelectWinner()
	require.Error(t, err)
}

func TestRandomNumber(t *testing.T) {
	result := core.RandomNumber(10, true)
	require.LessOrEqual(t, result, 99999)

	resultStr := core.TicketNumber(5)
	require.NotEmpty(t, resultStr)

	resultMel := core.MelGenerator(3)
	require.NotEmpty(t, resultMel)

}
