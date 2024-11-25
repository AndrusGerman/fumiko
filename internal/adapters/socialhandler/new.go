package socialhandler

import (
	"go.uber.org/fx"
)

func NewHandlers(constructors ...any) fx.Option {

	var handlers = make([]any, len(constructors))

	for i := range handlers {
		handlers[i] = fx.Annotate(
			constructors[i],
			fx.ResultTags(`group:"socialHandlers"`),
		)
	}

	return fx.Module("socialHandlers",
		fx.Provide(
			handlers...,
		),
	)

}
