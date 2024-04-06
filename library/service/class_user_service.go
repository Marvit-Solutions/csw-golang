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

// The implementation of ClassUserRepository interface.
type classUserService struct {
	db     *gorm.DB
	client *elastic.Client
}

// Creates a new instance of ClassUserService.
func NewClassUserService(db *gorm.DB, client *elastic.Client) repository.ClassUserRepository {
	return &classUserService{db, client}
}

// Finds a single record by given criteria.
func (srv *classUserService) FindOneBy(criteria map[string]interface{}) (*model.ClassUser, error) {
	m := new(model.ClassUser)
	res := srv.db.Where(criteria).First(&m)
	if err := res.Error; err != nil {
		return nil, err
	}
	return m, nil
}

// Finds records by given criteria, with pagination support.
func (srv *classUserService) FindBy(criteria map[string]interface{}, page, size int) ([]*model.ClassUser, error) {
	var data []*model.ClassUser

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Order("id DESC").Limit(limit).Find(&data); res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}

// Counts the number of records based on given criteria.
func (srv *classUserService) Count(criteria map[string]interface{}) int {
	var result int64

	if res := srv.db.Model(model.ClassUser{}).Where(criteria).Count(&result); res.Error != nil {
		return 0
	}

	return int(result)
}


// Creates a new record.
func (srv *classUserService) Create(model *model.ClassUser, tx *gorm.DB) (*model.ClassUser, error) {
		db := tx.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

// Updates an existing record.
func (srv *classUserService) Update(model *model.ClassUser, tx *gorm.DB) error {
	err := tx.Save(&model).Error
	return err
}

// Deletes an existing record.
func (srv *classUserService) Delete(model *model.ClassUser, tx *gorm.DB) error {
	err := tx.Delete(&model).Error
	return err
}

// Creates or updates an index for the model.
func (srv *classUserService) CreateOrUpdateIndex(model *model.ClassUser) error {
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
