package repositories

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"errors"
	"fmt"

	"github.com/phuslu/log"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) contract.IPostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) Create(post *models.Post) error {
	if err := p.db.
		Create(&post).
		Preload("Author").
		Last(&post).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO CREATE POST : " + err.Error())).Msg("")
		return err
	}

	return nil
}

func (p *postRepository) GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string, preloads ...string) (*models.Post, error) {
	db := p.db

	for _, p := range preloads {
		db = db.Preload(p)
	}

	for field, value := range values {
		db = db.Where(fmt.Sprintf("%s = ?", field), value)
	}

	var post *models.Post
	if err := db.First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error().Err(errors.New("FAILED TO GET DETAIL POST : " + err.Error())).Msg("")
		return nil, err
	}

	return post, nil
}

func (p *postRepository) GetAllByCursorAndSelectedFieldsPaginate(
	values map[string]interface{},
	cursor *helpers.Cursor,
	selectedFields string,
	preloads ...string) ([]*models.Post, *helpers.CursorPagination, error) {
	db := p.db
	db = helpers.EqualCondition(db, values)

	var total int64
	if err := db.Table("posts").Select("id").Count(&total).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO GET TOTAL COMMENTS : " + err.Error())).Msg("")
		return nil, nil, err
	}

	if total == 0 {
		return nil, &helpers.CursorPagination{}, nil
	}

	db = helpers.Preloads(db, preloads...)

	var posts []*models.Post
	if err := db.
		Select(selectedFields).
		Offset(cursor.CurrentPage*cursor.PerPage - cursor.PerPage).
		Limit(cursor.PerPage).
		Find(&posts).
		Order("created DESC").Error; err != nil {
		log.Error().Err(errors.New("FAILED TO GET ALL COMMENTS : " + err.Error())).Msg("")
		return nil, nil, err
	}

	cursorPagination := cursor.GeneratePager(total)

	return posts, cursorPagination, nil
}

func (p *postRepository) Update(post *models.Post) error {
	if err := p.db.Updates(&post).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO UPDATE POST : " + err.Error())).Msg("")
		return err
	}

	return nil
}

func (p *postRepository) Delete(post *models.Post) error {
	if err := p.db.Delete(&post).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO DELETE POST : " + err.Error())).Msg("")
		return err
	}

	return nil
}
