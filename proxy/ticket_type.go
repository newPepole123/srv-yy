package proxy

import (
	"github.com/newPepole123/srv-yy/model"
	"github.com/newPepole123/srv-yy/repository"
)

// CreateTicketTypeID generate a new ticketTypeID
func CreateTicketTypeID() uint64 {
	repo := repository.CreateTicketTypeRepository()
	return repo.CreateTicketTypeID()
}

// GetTicketTypesByKey get ticketTypes by key
func GetTicketTypesByKey(key string, page int, pageSize int) (*model.TicketTypeCollection, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.GetTicketTypesByKey(key, page, pageSize)
}

// GetTicketType by id uint64
func GetTicketType(id uint64) (*model.TicketType, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.GetTicketType(id)
}

// CreateTicketType ID, TicketTypeName, TicketTypeDescription, Status, CreatedAt
// return uint64, error
func CreateTicketType(ticketType *model.TicketType) (uint64, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.CreateTicketType(ticketType)
}

// UpdateTicketType return rowsAffected, error
// SET TicketTypeName, TicketTypeDescription, Status, UpdatedAt
// WHERE ID
func UpdateTicketType(ticketType *model.TicketType) (int64, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.UpdateTicketType(ticketType)
}

// UpdateTicketTypeStatus return rowsAffected, error
// SET status
// WHERE ID
func UpdateTicketTypeStatus(ticketType *model.TicketType) (int64, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.UpdateTicketTypeStatus(ticketType)
}

// DestroyTicketType return rowsAffected, error
// WHERE id uint64
func DestroyTicketType(id uint64) (int64, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.DestroyTicketType(id)
}

// DestroyTicketTypeSoft return rowsAffected, error
// WHERE id uint64
func DestroyTicketTypeSoft(id uint64) (int64, error) {
	repo := repository.CreateTicketTypeRepository()
	return repo.DestroyTicketTypeSoft(id)
}
