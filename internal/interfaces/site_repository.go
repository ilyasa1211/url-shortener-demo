package interfaces

import "github.com/ilyasa1211/url-shortener-demo/internal/domain"

type SiteRepository interface {
	All() *[]domain.Site
	FindByAlias(aliasUrl string) (*domain.Site, error)
	Create(site *domain.Site) error
	UpdateByAlias(aliasUrl string, targetUrl string) error
	DeleteByAlias(aliasUrl string) error
}
