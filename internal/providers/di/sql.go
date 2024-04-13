package diProvider

import (
	"github.com/BobrePatre/kozodoy-parser/internal/config/datasources"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
)

func (p *Provider) PostgresqlConfig() *datasources.PostgresqlConfig {
	if p.postgresqlConfig == nil {
		cfg, err := datasources.NewPostgresConfig(p.Validate())
		if err != nil {
			slog.Error("failed to get postgresql config", "err", err.Error())
			os.Exit(1)
		}
		p.postgresqlConfig = cfg
	}

	return p.postgresqlConfig
}

func (p *Provider) SqlDatabase() *sqlx.DB {
	if p.sqlDatabase == nil {
		db, err := sqlx.Connect("postgres", p.PostgresqlConfig().Datasource("disable"))
		slog.Debug("postgres config", "config", p.PostgresqlConfig())
		if err != nil {
			slog.Error("failed to connect to sql database", "err", err.Error())
			os.Exit(1)
		}
		p.sqlDatabase = db
	}

	return p.sqlDatabase
}
