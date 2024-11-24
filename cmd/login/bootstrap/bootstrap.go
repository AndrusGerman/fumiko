package bootstrap

import (
	"github.com/AndrusGerman/fumiko/internal/adapters/social/whatsapp"
	"github.com/AndrusGerman/fumiko/internal/core/ports"
	"go.uber.org/fx"
)

// deps
func privide() fx.Option {
	return fx.Provide(
		// social manager
		whatsapp.New,
	)
}

func Run() {
	var app = fx.New(
		privide(),
		fx.Invoke(start),
	)

	app.Run()
}

func start(social ports.Social, lc fx.Lifecycle, app fx.Shutdowner) {
	var prepare = make(chan struct{})
	lc.Append(fx.StartHook(func() {
		prepare <- struct{}{}
	}))

	go func() {
		<-prepare
		var err = social.Register()
		if err == nil {
			app.Shutdown(fx.ExitCode(0))
		} else {
			app.Shutdown(fx.ExitCode(1))
		}
	}()

}
