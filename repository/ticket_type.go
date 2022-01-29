package repository

import (
	"sync"
	"sync/atomic"

	"github.com/newPepole123/srv-yy/contract"
	"github.com/newPepole123/srv-yy/model"
)

var (
	_ticketTypeID             uint64
	_ticketTypeRepository     contract.TicketTypeRepository
	_onceTicketTypeRepository sync.Once
)

// CreateTicketTypeRepository return contract.TicketTypeRepository
func CreateTicketTypeRepository() contract.TicketTypeRepository {

	_onceTicketTypeRepository.Do(func() {
		_ticketTypeRepository = &TicketTypeRepository{}

		if _ticketTypeID == 0 {
			_ticketTypeID, _ = max("ticket_types", "id")

			if _ticketTypeID == 0 {
				_ticketTypeID = WebConfig().App.AppID - WebConfig().App.AppNum
			}
		}
	})

	return _ticketTypeRepository
}

// TicketTypeRepository struct
type TicketTypeRepository struct {
}

// CreateTicketTypeID generate a new ticketTypeID
func (r *TicketTypeRepository) CreateTicketTypeID() uint64 {
	return atomic.AddUint64(&_ticketTypeID, WebConfig().App.AppNum)
}

// GetTicketTypesByKey get ticketTypes by key
func (r *TicketTypeRepository) GetTicketTypesByKey(key string, page int, pageSize int) (*model.TicketTypeCollection, error) {
	sqlx := "SELECT `id`, `ticket_type_name`, `ticket_type_description`, `status`, `created_at`, `updated_at` " +
		"FROM `ticket_types` " +
		"WHERE `ticket_type_name` like ? and `status` > 0 " +
		"limit ? offset ? "

	key = "%" + key + "%"

	if pageSize > _maxPageSize {
		pageSize = _maxPageSize
	} else if pageSize <= 0 {
		pageSize = _pageSize
	}

	offset := 0

	if page > 1 {
		offset = (page - 1) * pageSize
	}

	rows, err := query(sqlx, key, pageSize, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ticketTypes := model.CreateTicketTypeCollection()

	for rows.Next() {

		ticketType := model.CreateTicketType()

		err := rows.Scan(&ticketType.ID, &ticketType.TicketTypeName, &ticketType.TicketTypeDescription, &ticketType.Status, &ticketType.CreatedAt, &ticketType.UpdatedAt)

		if err != nil {
			return nil, err
		}

		*ticketTypes = append(*ticketTypes, *ticketType)
	}

	return ticketTypes, rows.Err()
}

// GetTicketType by id uint64
func (r *TicketTypeRepository) GetTicketType(id uint64) (*model.TicketType, error) {
	sqlx := "SELECT `id`, `ticket_type_name`, `ticket_type_description`, `status`, `created_at`, `updated_at` " +
		"FROM `ticket_types` " +
		"WHERE `id` = ? and `status` > 0 " +
		"limit 1 "

	row := queryRow(sqlx, id)

	ticketType := model.CreateTicketType()

	err := row.Scan(&ticketType.ID, &ticketType.TicketTypeName, &ticketType.TicketTypeDescription, &ticketType.Status, &ticketType.CreatedAt, &ticketType.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return ticketType, nil
}

// CreateTicketType ID, TicketTypeName, TicketTypeDescription, Status, CreatedAt
// return uint64, error
func (r *TicketTypeRepository) CreateTicketType(ticketType *model.TicketType) (uint64, error) {
	sqlx := "INSERT INTO `ticket_types` " +
		"(`id`, `ticket_type_name`, `ticket_type_description`, `status`, `created_at`) " +
		"VALUES(?, ?, ?, ?, ?) "

	if ticketType.ID == 0 {
		ticketType.ID = r.CreateTicketTypeID()
	}

	_, err := exec(sqlx, ticketType.ID, ticketType.TicketTypeName, ticketType.TicketTypeDescription, ticketType.Status, now())

	if err != nil {
		return 0, err
	}

	return ticketType.ID, nil
}

// UpdateTicketType return rowsAffected, error
// SET TicketTypeName, TicketTypeDescription, Status, UpdatedAt
// WHERE ID
func (r *TicketTypeRepository) UpdateTicketType(ticketType *model.TicketType) (int64, error) {
	sqlx := "UPDATE `ticket_types` " +
		"SET `ticket_type_name` = ?, `ticket_type_description` = ?, `updated_at` = ? " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, ticketType.TicketTypeName, ticketType.TicketTypeDescription, now(), ticketType.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// UpdateTicketTypeStatus return rowsAffected, error
// SET status
// WHERE ID
func (r *TicketTypeRepository) UpdateTicketTypeStatus(ticketType *model.TicketType) (int64, error) {
	sqlx := "UPDATE `ticket_types` " +
		"SET `status` = ?, `updated_at` = ? " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, ticketType.Status, now(), ticketType.ID)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTicketType return rowsAffected, error
// WHERE id uint64
func (r *TicketTypeRepository) DestroyTicketType(id uint64) (int64, error) {
	sqlx := "DELETE FROM `ticket_types` WHERE `id` = ? "

	result, err := exec(sqlx, id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}

// DestroyTicketTypeSoft return rowsAffected, error
// WHERE id uint64
func (r *TicketTypeRepository) DestroyTicketTypeSoft(id uint64) (int64, error) {
	sqlx := "UPDATE `ticket_types` SET `deleted_at` = ?, status=-ABS(status) " +
		"WHERE `id` = ? "

	result, err := exec(sqlx, now(), id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
