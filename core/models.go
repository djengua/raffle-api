package core

type Participant struct {
	Name    string   `json:"name"`
	Tickets []string `json:"tickets"`
}

type Raffle struct {
	Name         string        `json:"name"`
	Prize        string        `json:"prize"`
	MaxTickets   int           `json:"maxTickets"`
	Participants []Participant `json:"participants"`
	FirstTaken   bool          `json:"firstTaken"`
	Turns        int           `json:"turns"`
	Tickets      []string      `json:"tickets"`
	Log          []string      `json:"log"`
}

type Tanda struct {
}
