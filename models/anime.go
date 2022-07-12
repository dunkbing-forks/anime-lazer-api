package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Anime struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	Title         string             `json:"title,omitempty" validate:"required"`
	Image         string             `json:"image,omitempty" validate:"required"`
	Type          string             `json:"type,omitempty" validate:"required"`
	Summary       string             `json:"summary,omitempty" validate:"required"`
	Released      string             `json:"released,omitempty"`
	Genres        []Genre            `json:"genres,omitempty"`
	Status        string             `json:"status,omitempty"`
	TotalEpisodes int                `json:"totalEpisodes,omitempty"`
	OtherNames    []string           `json:"otherNames,omitempty"`
}
