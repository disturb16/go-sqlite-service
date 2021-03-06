package dto

import "github.com/disturb16/go-sqlite-service/internal/persons/entity"

// disclaimer: there are non-ok json tag values for some of the *RQ structs, because they are in path or somewhere else, so it is not actual json

// swagger:parameters getThingRQ
type getThingRQ struct {
	// in: query
	ID string `json:"id"`
}

// swagger:model getThingRS
type getThingRS struct {
	// in: body
	Thing *entity.Thing `json:"thing"`
}
