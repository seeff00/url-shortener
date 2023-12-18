package repository

import (
	"github.com/go-xorm/xorm"
	"time"
)

type VisitorsRepository struct {
	Repository
	db *xorm.Engine
}

type Visitors struct {
	ID        int64     `xorm:"id" json:"id,omitempty"`
	IP        string    `xorm:"ip" json:"ip"`
	UrlId     int64     `xorm:"url_id" json:"url_id"`
	VisitedAt time.Time `xorm:"visited_at" json:"visited_at"`
}

func NewVisitorsRepository(db *xorm.Engine) VisitorsRepository {
	return VisitorsRepository{db: db}
}

func (pr *VisitorsRepository) Create(visitor Visitors) (int64, error) {
	return pr.db.Insert(visitor)
}

func (pr *VisitorsRepository) Get(visitor Visitors) (*Visitors, error) {
	isExists, err := pr.db.Get(&visitor)
	if err != nil || !isExists {
		return nil, err
	}

	return &visitor, err
}

func (pr *VisitorsRepository) GetAll() ([]Visitors, error) {
	var images []Visitors
	err := pr.db.Find(&images)
	if err != nil {
		return nil, err
	}

	return images, err
}

func (pr *VisitorsRepository) Update(visitor Visitors) (int64, error) {
	return pr.db.Update(&visitor, &Visitors{ID: visitor.ID})
}

func (pr *VisitorsRepository) Delete(id int64) (int64, error) {
	visitor := &Visitors{ID: id}
	return pr.db.Delete(visitor)
}

func (pr *VisitorsRepository) Exists(visitor Visitors) (bool, error) {
	return pr.db.Exist(&visitor)
}
