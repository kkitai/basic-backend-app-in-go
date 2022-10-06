// Package repository implements data access layer for Telephone ORM structure
package repository

import (
	"fmt"
	"time"

	"github.com/kkitai/basic-backend-app-in-go/model"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// A TelephoneRepository interface is an interface of Telephone
type TelephoneRepository interface {
	GetTelephoneByNumber(string) (*model.Telephone, error)
	ListTelephone() ([]*model.Telephone, error)
	PostTelephone(int, int, string) error
	PutTelephoneByNumber(string, int, int) error
}

// A telephoneRepository serves db connection of gorm
type telephoneRepository struct {
	connection *gorm.DB
}

// NewTelephoneRepository returns an address of telephoneRepository which implements TelephoneRepository interface.
func NewTelephoneRepository(conn *gorm.DB) TelephoneRepository {
	return &telephoneRepository{
		connection: conn,
	}
}

// GetTelephoneByNumber returns an address of Telephone and error.
// this function issues a select query that have where clause with "number" and limits first record.
func (t *telephoneRepository) GetTelephoneByNumber(number string) (*model.Telephone, error) {
	if number == "" {
		return &model.Telephone{}, fmt.Errorf("empty number was given")
	}

	var telephone *model.Telephone
	err := t.connection.Where("number = ?", number).First(&telephone).Error

	return telephone, err
}

// ListTelephone returns an array of an address of Telephone and error.
// this function issues a select query to get all telephone records.
func (t *telephoneRepository) ListTelephone() ([]*model.Telephone, error) {
	var telephones []*model.Telephone
	err := t.connection.Find(&telephones).Error

	return telephones, err
}

// PostTelephone returns error.
// this function issues a insert query to register telephone record.
func (t *telephoneRepository) PostTelephone(ownerId int, iccId int, number string) error {
	telephone := &model.Telephone{
		OwnerId: ownerId,
		Number:  number,
		ICCId:   iccId,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
	}
	return t.connection.Create(telephone).Error
}

// PutTelephoneByNumber returns error.
// this function issues update query that have where clause with "number" to update single record of telephone.
func (t *telephoneRepository) PutTelephoneByNumber(number string, ownerId, iccId int) error {
	var telephone *model.Telephone
	return t.connection.Model(telephone).Where("number = ?", number).Updates(map[string]interface{}{"owner_id": ownerId, "icc_id": iccId}).Error
}
