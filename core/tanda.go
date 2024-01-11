package core

type TandaCore interface {
	AddParticipant() error
	RandomizeNumbers() error
	GetInformation() ([]string, error)
}

func (r *Tanda) AddParticipant(newParticipant TandaParticipant) error {
	// if !r.Open {
	// 	return errors.New(util.RAFFLE_IS_CLOSED)
	// }

	// if len(r.Participants) == 0 {
	// 	r.Participants = []RaffleParticipant{}
	// }

	// if newParticipant.ID == primitive.NilObjectID {
	// 	newParticipant.ID = primitive.NewObjectID()
	// }

	// for i := 0; i < len(r.Participants); i++ {
	// 	if r.Participants[i].Name == newParticipant.Name || r.Participants[i].ID == newParticipant.ID {
	// 		return errors.New(util.PARTICIPANT_DUPLICATED)
	// 	}
	// }

	// r.Participants = append(r.Participants, newParticipant)
	return nil
}
