package repositories

import (
	contract "backend-takehome-blog/contracts"
	"backend-takehome-blog/helpers"
	"backend-takehome-blog/models"
	"errors"

	"github.com/phuslu/log"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) contract.ICommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cm *commentRepository) Create(comment *models.Comment) error {
	if err := cm.db.
		Create(&comment).
		Preload("Author").
		Last(&comment).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO CREATE COMMENT : " + err.Error())).Msg("")
		return err
	}

	return nil
}

func (cm *commentRepository) GetByCustomAndSelectedFields(values map[string]interface{}, selectedFields string, preloads ...string) (*models.Comment, error) {
	db := cm.db
	db = helpers.EqualCondition(db, values)
	db = helpers.Preloads(db, preloads...)

	var comment *models.Comment
	if err := db.Select(selectedFields).First(&comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		log.Error().Err(errors.New("FAILED TO GET COMMENT : " + err.Error())).Msg("")
		return nil, err
	}

	return comment, nil
}

func (cm *commentRepository) GetAllByCursorAndSelectedFieldsPaginate(
	values map[string]interface{},
	cursor *helpers.Cursor,
	selectedFields string,
	preloads ...string) ([]*models.Comment, *helpers.CursorPagination, error) {

	db := cm.db
	db = helpers.EqualCondition(db, values)

	var total int64
	if err := db.Table("comments").Select("id").Count(&total).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO GET TOTAL COMMENTS : " + err.Error())).Msg("")
		return nil, nil, err
	}

	if total == 0 {
		return nil, &helpers.CursorPagination{}, nil
	}

	db = helpers.Preloads(db, preloads...)

	var comments []*models.Comment
	if err := db.
		Select(selectedFields).
		Offset(cursor.CurrentPage*cursor.PerPage - cursor.PerPage).
		Limit(cursor.PerPage).
		Find(&comments).
		Order("created DESC").Error; err != nil {
		log.Error().Err(errors.New("FAILED TO GET ALL COMMENTS : " + err.Error())).Msg("")
		return nil, nil, err
	}

	cursorPagination := cursor.GeneratePager(total)

	return comments, cursorPagination, nil
}

func (cm *commentRepository) GetTotalByCustom(values map[string]interface{}) (int64, error) {
	db := cm.db

	db = helpers.EqualCondition(db, values)

	var counter int64
	if err := db.Select("id").Table("comments").Count(&counter).Error; err != nil {
		log.Error().Err(errors.New("FAILED TO GET TOTAL COMMENTS : " + err.Error())).Msg("")
		return 0, err
	}

	return counter, nil
}
