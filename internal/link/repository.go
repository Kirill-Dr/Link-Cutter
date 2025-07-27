package link

import "link-cutter/pkg/db"

type LinkRepository struct {
	Database *db.DB
}

func NewLinkRepository(database *db.DB) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}
