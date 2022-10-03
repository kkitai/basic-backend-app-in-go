package db

import (
	"fmt"
	"time"

	"github.com/kkitai/basic-backend-app-in-go/model"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type TelephoneAccesscer interface {
	GetTelephoneByNumber() (*model.Telephone, error)
	ListTelephone() (*model.Telephone, error)
	PostTelephone() (*model.Telephone, error)
	PutTelephoneByNumber() (*model.Telephone, error)
}

func (d *DB) GetTelephoneByNumber(number string) (*model.Telephone, error) {
	if number == "" {
		return &model.Telephone{}, fmt.Errorf("empty number was given")
	}

	var telephone *model.Telephone
	err := d.Connection.Where("number = ?", number).First(&telephone).Error

	return telephone, err
}

func (d *DB) ListTelephone() ([]*model.Telephone, error) {
	var telephones []*model.Telephone
	err := d.Connection.Find(&telephones).Error

	return telephones, err
}

func (d *DB) PostTelephone(ownerId int, iccId int, number string) error {
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
	tx := d.Connection.Create(telephone)
	return tx.Error
}

func (d *DB) PutTelephoneByNumber(number string, ownerId, iccId int) error {
	var telephone *model.Telephone
	tx := d.Connection.Model(telephone).Where("number = ?", number).Updates(map[string]interface{}{"owner_id": ownerId, "icc_id": iccId})
	return tx.Error
}
