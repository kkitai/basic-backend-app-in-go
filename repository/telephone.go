package repository

import (
	"fmt"
	"time"

	"github.com/kkitai/basic-backend-app-in-go/model"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type TelephoneRepository interface {
	GetTelephoneByNumber(string) (*model.Telephone, error)
	ListTelephone() ([]*model.Telephone, error)
	PostTelephone(int, int, string) error
	PutTelephoneByNumber(string, int, int) error
}

type telephoneRepository struct {
	connection *gorm.DB
}

func NewTelephoneRepository(conn *gorm.DB) TelephoneRepository {
	return &telephoneRepository{
		connection: conn,
	}
}

func (t *telephoneRepository) GetTelephoneByNumber(number string) (*model.Telephone, error) {
	if number == "" {
		return &model.Telephone{}, fmt.Errorf("empty number was given")
	}

	var telephone *model.Telephone
	err := t.connection.Where("number = ?", number).First(&telephone).Error

	return telephone, err
}

func (t *telephoneRepository) ListTelephone() ([]*model.Telephone, error) {
	var telephones []*model.Telephone
	err := t.connection.Find(&telephones).Error

	return telephones, err
}

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

func (t *telephoneRepository) PutTelephoneByNumber(number string, ownerId, iccId int) error {
	var telephone *model.Telephone
	return t.connection.Model(telephone).Where("number = ?", number).Updates(map[string]interface{}{"owner_id": ownerId, "icc_id": iccId}).Error
}
