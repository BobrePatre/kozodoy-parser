package app

import (
	"context"
	"github.com/BobrePatre/kozodoy-parser/internal/providers/di"
)

var _ = (*App)(nil)

func (a *App) initDeps(ctx context.Context) error {
	inits := []configFunc{
		a.initDiProvider,
		a.initLogger,
		a.initHTTPServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initDiProvider(_ context.Context) error {
	a.diProvider = diProvider.NewDiProvider()
	return nil
}
