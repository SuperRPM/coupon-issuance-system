package campaign

import "time"

type Campaign struct {
	ID          int
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Limit       int
	IssuedCount int
}

func NewCampaign(name string, limit int, startDate time.Time, endDate time.Time) *Campaign {
	return &Campaign{
		Name:        name,
		Limit:       limit,
		StartDate:   startDate,
		EndDate:     endDate,
		IssuedCount: 0,
	}
}
