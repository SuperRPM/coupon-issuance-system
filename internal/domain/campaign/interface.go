package campaign

import "context"

type Handler interface {
	CreateCampaign(ctx context.Context, campaign *Campaign) error
	GetCampaign(ctx context.Context, id int) (*Campaign, error)
}

type Service interface {
	CreateCampaign(ctx context.Context, campaign *Campaign) error
	GetCampaign(ctx context.Context, id int) (*Campaign, error)
}

type Repository interface {
	Create(campaign *Campaign) error
	Get(id int) (*Campaign, error)
}
