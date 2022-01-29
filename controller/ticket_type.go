package controller

import (
	"sync"

	"github.com/newPepole123/srv-yy/model"
	"github.com/newPepole123/srv-yy/proxy"
	"github.com/newPepole123/srv-yy/validator"
	"github.com/webpkg/web"
)

var (
	_ticketTypeController     *TicketTypeController
	_onceTicketTypeController sync.Once
)

// CreateTicketTypeController return TicketTypeController
func CreateTicketTypeController() *TicketTypeController {

	_onceTicketTypeController.Do(func() {
		_ticketTypeController = &TicketTypeController{}
	})

	return _ticketTypeController
}

// TicketTypeController struct
type TicketTypeController struct {
}

// Index get ticketTypes
func (c *TicketTypeController) Index(ctx *web.Context) (web.Data, error) {
	var (
		page     int
		pageSize int
	)

	key := ctx.Query("key")
	ctx.TryParseQuery("page", &page)
	ctx.TryParseQuery("pagesize", &pageSize)

	return proxy.GetTicketTypesByKey(key, page, pageSize)
}

// Create create ticketType
func (c *TicketTypeController) Create(ctx *web.Context) (web.Data, error) {
	ticketType := model.CreateTicketType()

	if err := ctx.TryParseBody(ticketType); err != nil {
		return nil, err
	}

	if err := validator.CreateTicketType(ticketType); err != nil {
		return nil, err
	}

	return proxy.CreateTicketType(ticketType)
}

// Detail get ticketType detail by id
func (c *TicketTypeController) Detail(ctx *web.Context) (web.Data, error) {
	var (
		id uint64
	)

	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	return proxy.GetTicketType(id)
}

// Update update ticketType by id
func (c *TicketTypeController) Update(ctx *web.Context) (web.Data, error) {
	ticketType := model.CreateTicketType()

	if err := ctx.TryParseBody(ticketType); err != nil {
		return nil, err
	}

	if err := ctx.TryParseParam("id", &ticketType.ID); err != nil {
		return nil, err
	}

	if err := validator.UpdateTicketType(ticketType); err != nil {
		return nil, err
	}

	return proxy.UpdateTicketType(ticketType)
}

// UpdateStatus update ticketType status by id
func (c *TicketTypeController) UpdateStatus(ctx *web.Context) (web.Data, error) {
	ticketType := model.CreateTicketType()

	if err := ctx.TryParseBody(ticketType); err != nil {
		return nil, err
	}

	if err := ctx.TryParseParam("id", &ticketType.ID); err != nil {
		return nil, err
	}

	if err := validator.UpdateTicketTypeStatus(ticketType); err != nil {
		return nil, err
	}

	return proxy.UpdateTicketTypeStatus(ticketType)
}

// Destroy delete ticketType by id
func (c *TicketTypeController) Destroy(ctx *web.Context) (web.Data, error) {
	var (
		id uint64
	)
	if err := ctx.TryParseParam("id", &id); err != nil {
		return nil, err
	}

	return proxy.DestroyTicketTypeSoft(id)
}
