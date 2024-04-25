package diProvider

import (
	parserHandler "github.com/BobrePatre/kozodoy-parser/internal/delivery/http/handlers/parser"
	parserCategoryRepository "github.com/BobrePatre/kozodoy-parser/internal/repository/categories"
	parserDishRepository "github.com/BobrePatre/kozodoy-parser/internal/repository/dish"
	parserMenuRepository "github.com/BobrePatre/kozodoy-parser/internal/repository/menu"
	parserService "github.com/BobrePatre/kozodoy-parser/internal/service/parser"
)

func (p *Provider) ParserMenuRepository() *parserMenuRepository.Repository {
	if p.parserMenuRepository == nil {
		p.parserMenuRepository = parserMenuRepository.NewRepository(p.ServiceAccessProvider(), p.NetworkDatacourceConfig())
	}
	return p.parserMenuRepository
}

func (p *Provider) ParserCategoryRepository() *parserCategoryRepository.Repository {
	if p.parserCategoryRepository == nil {
		p.parserCategoryRepository = parserCategoryRepository.NewRepository(p.ServiceAccessProvider(), p.NetworkDatacourceConfig())
	}
	return p.parserCategoryRepository
}

func (p *Provider) ParserDishRepository() *parserDishRepository.Repository {
	if p.parserDishRepository == nil {
		p.parserDishRepository = parserDishRepository.NewRepository(p.ServiceAccessProvider(), p.NetworkDatacourceConfig())
	}
	return p.parserDishRepository
}

func (p *Provider) ParserService() *parserService.Service {
	if p.parserService == nil {
		p.parserService = parserService.NewService(
			p.ParserMenuRepository(),
			p.ParserCategoryRepository(),
			p.ParserDishRepository(),
		)
	}
	return p.parserService
}

func (p *Provider) ParserHandler() *parserHandler.Handler {
	if p.parserHandler == nil {
		p.parserHandler = parserHandler.NewHandler(p.ParserService())
	}
	return p.parserHandler
}
