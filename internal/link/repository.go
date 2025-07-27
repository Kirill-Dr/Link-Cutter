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

func (repo *LinkRepository) Create(link *Link) {

}
