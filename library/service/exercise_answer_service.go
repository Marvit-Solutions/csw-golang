package service

import (
	"context"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// The implementation of ExerciseAnswerRepository interface.
type exerciseAnswerService struct {
	db     *gorm.DB
	client *elastic.Client
}

// Creates a new instance of ExerciseAnswerService.
func NewExerciseAnswerService(db *gorm.DB, client *elastic.Client) repository.ExerciseAnswerRepository {
	return &exerciseAnswerService{db, client}
}

// Finds a single record by given criteria.
func (srv *exerciseAnswerService) FindOneBy(criteria map[string]interface{}) (*model.ExerciseAnswer, error) {
	m := new(model.ExerciseAnswer)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Finds records by given criteria, with pagination support.
func (srv *exerciseAnswerService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseAnswer, error) {
	var data []*model.ExerciseAnswer

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// Counts the number of records based on given criteria.
func (srv *exerciseAnswerService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.ExerciseAnswer{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}


// Creates a new record.
func (srv *exerciseAnswerService) Create(model *model.ExerciseAnswer, tx *gorm.DB) (*model.ExerciseAnswer, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Updates an existing record.
func (srv *exerciseAnswerService) Update(model *model.ExerciseAnswer, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

// Deletes an existing record.
func (srv *exerciseAnswerService) Delete(model *model.ExerciseAnswer, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

// Creates or updates an index for the model.
func (srv *exerciseAnswerService) CreateOrUpdateIndex(model *model.ExerciseAnswer) error {
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
