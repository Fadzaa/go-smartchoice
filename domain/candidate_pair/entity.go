package candidate_pair

import (
	"smartchoice/domain/party"
)

type CandidatePair struct {
	ID                  uint          `json:"id" gorm:"primaryKey"`
	PairName            string        `json:"pair_name" gorm:"maxLength:100"`
	PairSlogan          string        `json:"pair_slogan" gorm:"maxLength:100"`
	PresId              uint          `json:"pres_id"`
	PresCandidate       string        `json:"pres_candidate" gorm:"maxLength:100"`
	PresImage           string        `json:"pres_image"`
	ViceId              uint          `json:"vice_id"`
	ViceCandidate       string        `json:"vice_candidate" gorm:"maxLength:100"`
	ViceImage           string        `json:"vice_image"`
	VisionMissionLink   string        `json:"vision_mission_link"`
	Education           string        `json:"education"`
	Health              string        `json:"health"`
	Job                 string        `json:"job"`
	Corrupt             string        `json:"corrupt"`
	FreedomOfExpression string        `json:"freedom_of_expression"`
	ClimateCrisis       string        `json:"climate_crisis"`
	Parties             []party.Party `json:"parties"` // 1 to Many
}
