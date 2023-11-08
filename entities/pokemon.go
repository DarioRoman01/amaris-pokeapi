package entities

import "github.com/lib/pq"

type Pokemon struct {
	ID           int            `json:"id" db:"id" validate:"required"`
	Name         string         `json:"name" db:"name"`
	Generation   string         `json:"generation" db:"generation" validate:"required"`
	Weight       int            `json:"weight" db:"weight" validate:"required"`
	Height       int            `json:"height" db:"height" validate:"required"`
	VersionGroup string         `json:"version_group" db:"version_group"`
	IsDefault    bool           `json:"is_default" db:"is_default"`
	Types        pq.StringArray `json:"types" db:"types"`
}
