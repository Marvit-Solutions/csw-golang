package service

import (
	"context"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// The implementation of QuizChoiceRepository interface.
type quizChoiceService struct {
	db     *gorm.DB
	client *elastic.Client
}

// Creates a new instance of QuizChoiceService.
func NewQuizChoiceService(db *gorm.DB, client *elastic.Client) repository.QuizChoiceRepository {
	return &quizChoiceService{db, client}
}

// Finds a single record by given criteria.
func (srv *quizChoiceService) FindOneBy(criteria map[string]interface{}) (*model.QuizChoice, error) {
	m := new(model.QuizChoice)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Finds records by given criteria, with pagination support.
func (srv *quizChoiceService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizChoice, error) {
	var data []*model.QuizChoice

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// Counts the number of records based on given criteria.
func (srv *quizChoiceService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.QuizChoice{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}


// Creates a new record.
func (srv *quizChoiceService) Create(model *model.QuizChoice, tx *gorm.DB) (*model.QuizChoice, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Updates an existing record.
func (srv *quizChoiceService) Update(model *model.QuizChoice, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

// Deletes an existing record.
func (srv *quizChoiceService) Delete(model *model.QuizChoice, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

// Creates or updates an index for the model.
func (srv *quizChoiceService) CreateOrUpdateIndex(model *model.QuizChoice) error {
	ctx := context.Background()

	exists, err := srv.client.IndexExists(model.TableName()).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		_, err = srv.client.CreateIndex(model.TableName()).Do(ctx)
		if err != nil {
			return err
		}
	}

	_, err = srv.client.Delete().
		Index(model.TableName()).
		Id(model.ID).
		Refresh("true").
		Do(ctx)

	if err != nil && !elastic.IsNotFound(err) {
		return err
	}

	_, err = srv.client.Index().
		Index(model.TableName()).
		Id(model.ID).
		BodyJson(&model).
		Do(context.Background())
	return err
}
