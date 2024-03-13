package service

import (
	"context"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// The implementation of ExerciseChoiceRepository interface.
type exerciseChoiceService struct {
	db     *gorm.DB
	client *elastic.Client
}

// Creates a new instance of ExerciseChoiceService.
func NewExerciseChoiceService(db *gorm.DB, client *elastic.Client) repository.ExerciseChoiceRepository {
	return &exerciseChoiceService{db, client}
}

// Finds a single record by given criteria.
func (srv *exerciseChoiceService) FindOneBy(criteria map[string]interface{}) (*model.ExerciseChoice, error) {
	m := new(model.ExerciseChoice)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Finds records by given criteria, with pagination support.
func (srv *exerciseChoiceService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseChoice, error) {
	var data []*model.ExerciseChoice

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// Counts the number of records based on given criteria.
func (srv *exerciseChoiceService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.ExerciseChoice{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}


// Creates a new record.
func (srv *exerciseChoiceService) Create(model *model.ExerciseChoice, tx *gorm.DB) (*model.ExerciseChoice, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Updates an existing record.
func (srv *exerciseChoiceService) Update(model *model.ExerciseChoice, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

// Deletes an existing record.
func (srv *exerciseChoiceService) Delete(model *model.ExerciseChoice, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

// Creates or updates an index for the model.
func (srv *exerciseChoiceService) CreateOrUpdateIndex(model *model.ExerciseChoice) error {
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
