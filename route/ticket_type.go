package route

import (
	"github.com/newPepole123/srv-yy/controller"
	"github.com/webpkg/web"
)

func ticketTypeRoute(app *web.Application, prefix string) {

	ticketType := controller.CreateTicketTypeController()

	app.Get(prefix+"/type/", middleware.Chain(ticketType.Index, "ticketType.all"))
	app.Post(prefix+"/type/", middleware.Chain(ticketType.Create, "ticketType.edit"))
	app.Get(prefix+"/type/:id", middleware.Chain(ticketType.Detail, "ticketType.all"))
	app.Patch(prefix+"/type/:id", middleware.Chain(ticketType.Update, "ticketType.edit"))
	app.Put(prefix+"/type/:id", middleware.Chain(ticketType.Update, "ticketType.edit"))
	app.Put(prefix+"/type/:id/status/", middleware.Chain(ticketType.UpdateStatus, "ticketType.edit"))
	app.Delete(prefix+"/type/:id", middleware.Chain(ticketType.Destroy, "ticketType.edit"))
}
