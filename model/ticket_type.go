package model

import "time"

// CreateTicketType return *TicketType
func CreateTicketType() *TicketType {

	ticketType := &TicketType{}

	return ticketType
}

// TicketType model
// @table ticket_types
type TicketType struct {
	// @column PrimaryKey
	ID uint64 `json:"id"`
	// @column $dataType=varchar(127)
	TicketTypeName        string  `json:"ticketTypeName"`
	TicketTypeDescription *string `json:"ticketTypeDescription"`
	// lt 0 deleted, 0 pendding, 1 valid
	Status    int        `json:"status"`
	DeletedAt *time.Time `json:"-"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

// CreateTicketTypeCollection return *TicketTypeCollection
func CreateTicketTypeCollection() *TicketTypeCollection {

	ticketTypeCollection := &TicketTypeCollection{}

	return ticketTypeCollection
}

// TicketTypeCollection TicketType list
type TicketTypeCollection []TicketType

// Len return len
func (o *TicketTypeCollection) Len() int { return len(*o) }

// Swap swap i, j
func (o *TicketTypeCollection) Swap(i, j int) { (*o)[i], (*o)[j] = (*o)[j], (*o)[i] }

// Less compare i, j
func (o *TicketTypeCollection) Less(i, j int) bool {
	return (*o)[i].TicketTypeName < (*o)[j].TicketTypeName
}
