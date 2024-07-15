package service

import (
	"context"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"strconv"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

// The implementation of ExerciseSubmissionsSubModuleRepository interface.
type exerciseSubmissionsSubModuleService struct {
	db     *gorm.DB
	client *elastic.Client
}

// Creates a new instance of ExerciseSubmissionsSubModuleService.
func NewExerciseSubmissionsSubModuleService(db *gorm.DB, client *elastic.Client) repository.ExerciseSubmissionsSubModuleRepository {
	return &exerciseSubmissionsSubModuleService{db, client}
}

// Finds a single record by given criteria.
func (srv *exerciseSubmissionsSubModuleService) FindOneBy(criteria map[string]interface{}) (*model.ExerciseSubmissionsSubModule, error) {
	m := new(model.ExerciseSubmissionsSubModule)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Finds records by given criteria, with pagination support.
func (srv *exerciseSubmissionsSubModuleService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseSubmissionsSubModule, error) {
	var data []*model.ExerciseSubmissionsSubModule

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// Counts the number of records based on given criteria.
func (srv *exerciseSubmissionsSubModuleService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.ExerciseSubmissionsSubModule{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}


// Creates a new record.
func (srv *exerciseSubmissionsSubModuleService) Create(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) (*model.ExerciseSubmissionsSubModule, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Updates an existing record.
func (srv *exerciseSubmissionsSubModuleService) Update(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

// Deletes an existing record.
func (srv *exerciseSubmissionsSubModuleService) Delete(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

// Creates or updates an index for the model.
func (srv *exerciseSubmissionsSubModuleService) CreateOrUpdateIndex(model *model.ExerciseSubmissionsSubModule) error {
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
		Id(strconv.Itoa(model.ID)).
		Refresh("true").
		Do(ctx)

	if err != nil && !elastic.IsNotFound(err) {
		return err
	}

	_, err = srv.client.Index().
		Index(model.TableName()).
		Id(strconv.Itoa(model.ID)).
		BodyJson(&model).
		Do(context.Background())
	return err
}
