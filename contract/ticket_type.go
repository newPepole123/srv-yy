package contract

import "github.com/newPepole123/srv-yy/model"

// TicketTypeRepository interface
type TicketTypeRepository interface {
	// CreateTicketTypeID generate a new ticketTypeID
	CreateTicketTypeID() uint64
	// GetTicketTypesByKey get ticketTypes by key
	GetTicketTypesByKey(key string, page int, pageSize int) (*model.TicketTypeCollection, error)
	// GetTicketType by id uint64
	GetTicketType(id uint64) (*model.TicketType, error)
	// CreateTicketType ID, TicketTypeName, TicketTypeDescription, Status, CreatedAt
	// return uint64, error
	CreateTicketType(ticketType *model.TicketType) (uint64, error)
	// UpdateTicketType return rowsAffected, error
	// SET TicketTypeName, TicketTypeDescription, Status, UpdatedAt
	// WHERE ID
	UpdateTicketType(ticketType *model.TicketType) (int64, error)
	// UpdateTicketTypeStatus return rowsAffected, error
	// SET status
	// WHERE ID
	UpdateTicketTypeStatus(ticketType *model.TicketType) (int64, error)
	// DestroyTicketType return rowsAffected, error
	// WHERE id uint64
	DestroyTicketType(id uint64) (int64, error)
	// DestroyTicketTypeSoft return rowsAffected, error
	// WHERE id uint64
	DestroyTicketTypeSoft(id uint64) (int64, error)
}
