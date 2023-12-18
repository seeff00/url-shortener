package repository

import (
	"github.com/go-xorm/xorm"
	"time"
)

type UrlsRepository struct {
	Repository
	db *xorm.Engine
}

type Urls struct {
	ID        int64     `xorm:"id" json:"id,omitempty"`
	Code      string    `xorm:"code" json:"code"`
	Url       string    `xorm:"url" json:"url"`
	IP        string    `xorm:"ip" json:"ip"`
	CreatedAt time.Time `xorm:"created_at" json:"created_at"`
}

func NewUrlsRepository(db *xorm.Engine) UrlsRepository {
	return UrlsRepository{db: db}
}

func (pr *UrlsRepository) Create(url Urls) (int64, error) {
	return pr.db.Insert(url)
}

func (pr *UrlsRepository) Get(url Urls) (*Urls, error) {
	isExists, err := pr.db.Get(&url)
	if err != nil || !isExists {
		return nil, err
	}

	return &url, err
}

func (pr *UrlsRepository) GetAll() ([]Urls, error) {
	var images []Urls
	err := pr.db.Find(&images)
	if err != nil {
		return nil, err
	}

	return images, err
}

func (pr *UrlsRepository) Update(url Urls) (int64, error) {
	return pr.db.Update(&url, &Urls{ID: url.ID})
}

func (pr *UrlsRepository) Delete(id int64) (int64, error) {
	url := &Urls{ID: id}
	return pr.db.Delete(url)
}

func (pr *UrlsRepository) Exists(url Urls) (bool, error) {
	return pr.db.Exist(&url)
}
