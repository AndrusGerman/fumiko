package social

import (
	"go.uber.org/fx"
)

func NewSocials(constructors ...any) fx.Option {

	var handlers = make([]any, len(constructors))

	for i := range handlers {
		handlers[i] = fx.Annotate(
			constructors[i],
			fx.ResultTags(`group:"social"`),
		)
	}

	return fx.Module("social",
		fx.Provide(
			handlers...,
		),
	)

}
