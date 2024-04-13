package diProvider

import (
	parserHandler "github.com/BobrePatre/kozodoy-parser/internal/delivery/http/handlers/parser"
	parserRepository "github.com/BobrePatre/kozodoy-parser/internal/repository/parser"
	parserService "github.com/BobrePatre/kozodoy-parser/internal/service/parser"
)

func (p *Provider) ParserRepository() *parserRepository.Repository {
	if p.parserRepository == nil {
		p.parserRepository = parserRepository.NewRepository()
	}
	return p.parserRepository
}

func (p *Provider) ParserService() *parserService.Service {
	if p.parserService == nil {
		p.parserService = parserService.NewService(p.ParserRepository())
	}
	return p.parserService
}

func (p *Provider) ParserHandler() *parserHandler.Handler {
	if p.parserHandler == nil {
		p.parserHandler = parserHandler.NewHandler(p.ParserService())
	}
	return p.parserHandler
}
