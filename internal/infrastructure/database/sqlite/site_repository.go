package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/ilyasa1211/url-shortener-demo/internal/domain"
)

type SiteRepository struct {
	DB *sql.DB
}

func NewSiteRepository(db *sql.DB) *SiteRepository {
	return &SiteRepository{db}
}

func (ur *SiteRepository) All() *[]domain.Site {
	rows, err := ur.DB.Query("SELECT * FROM sites")

	if err != nil {
		return nil
	}

	defer rows.Close()

	sites := make([]domain.Site, 0)
	var i int

	for rows.Next() {
		var site domain.Site

		if err := rows.Scan(&site.ID, &site.AliasUrl, &site.TargetUrl); err != nil {
			fmt.Println(err)
		}

		sites = append(sites, site)
		i++
	}

	return &sites
}
func (ur *SiteRepository) FindByAlias(aliasUrl string) (*domain.Site, error) {
	rows, err := ur.DB.Query("SELECT * FROM sites WHERE alias_url = ? LIMIT 1", aliasUrl)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var site domain.Site

	for rows.Next() {
		rows.Scan(&site.ID, &site.AliasUrl, &site.TargetUrl)
	}

	return &site, nil
}

func (ur *SiteRepository) Create(site *domain.Site) error {
	_, err := ur.DB.Exec("INSERT INTO sites (alias_url, target_url) VALUES (?, ?)", site.AliasUrl, site.TargetUrl)

	if err != nil {
		return err
	}

	return nil
}

func (ur *SiteRepository) UpdateByAlias(aliasUrl string, targetUrl string) error {
	_, err := ur.DB.Exec("UPDATE sites SET target_url = ? WHERE alias_url = ?", targetUrl, aliasUrl)

	if err != nil {
		return err
	}

	return nil
}
func (ur *SiteRepository) DeleteByAlias(aliasUrl string) error {
	_, err := ur.DB.Exec("DELETE sites WHERE alias_url = ?", aliasUrl)

	if err != nil {
		return err
	}

	return nil
}
