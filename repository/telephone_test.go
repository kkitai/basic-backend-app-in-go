package repository

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestGetTelephoneByNumber(t *testing.T) {

	cases := []struct {
		description    string
		inputNumber    string
		mockIds        []uint
		mockOwnerIds   []int
		mockICCIds     []int
		mockCreatedAts []time.Time
		mockUpdatedAts []time.Time
		mockDeletedAts []time.Time
	}{
		{
			description:    "when string value passed for number argument, it finished successfully",
			inputNumber:    "09011112222",
			mockIds:        []uint{1},
			mockOwnerIds:   []int{1},
			mockICCIds:     []int{111111111111111},
			mockCreatedAts: []time.Time{time.Now().Add(-5 * 24 * time.Hour)},
			mockUpdatedAts: []time.Time{time.Now().Add(-4 * 24 * time.Hour)},
			mockDeletedAts: []time.Time{time.Now()},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			mdb, mock, err := NewDbMock()
			if err != nil {
				t.Fatalf("an error '%s' was not expect when opening a stub database connection", err)
			}

			// mocking the results of db query
			mock.ExpectQuery(regexp.QuoteMeta(
				`SELECT * FROM "telephones" WHERE number = $1`)).
				WithArgs(tt.inputNumber).
				WillReturnRows(NewMockRow(mock, tt.mockOwnerIds, []string{tt.inputNumber}, tt.mockICCIds, tt.mockIds, tt.mockCreatedAts, tt.mockUpdatedAts, tt.mockDeletedAts))

			// execute the test fuction
			tr := telephoneRepository{
				connection: mdb,
			}
			tel, err := tr.GetTelephoneByNumber(tt.inputNumber)

			// assertion
			if tel.ID != tt.mockIds[0] {
				t.Errorf(`expectId(%d) was not matched (actual: %d)`, tt.mockIds[0], tel.ID)
			}
			if tel.OwnerId != tt.mockOwnerIds[0] {
				t.Errorf(`expectOwnerId(%d) was not matched (actual: %d)`, tt.mockOwnerIds[0], tel.OwnerId)
			}
			if tel.Number != tt.inputNumber {
				t.Errorf(`expectNumber(%s) was not matched (actual: %s)`, tt.inputNumber, tel.Number)
			}
			if tel.ICCId != tt.mockICCIds[0] {
				t.Errorf(`expectICCId(%d) was not matched (actual: %d)`, tt.mockICCIds[0], tel.ICCId)
			}
		})
	}
}

func TestListTelephone(t *testing.T) {

	cases := []struct {
		description    string
		mockIds        []uint
		mockNumbers    []string
		mockOwnerIds   []int
		mockICCIds     []int
		mockCreatedAts []time.Time
		mockUpdatedAts []time.Time
		mockDeletedAts []time.Time
	}{
		{
			description:    "when two records are returned from db, it returns the array struct of Telephone",
			mockIds:        []uint{1, 2},
			mockNumbers:    []string{"09011112222", "09022223333"},
			mockOwnerIds:   []int{1, 2},
			mockICCIds:     []int{111111111111111, 222222222222222},
			mockCreatedAts: []time.Time{time.Now().Add(-5 * 24 * time.Hour), time.Now().Add(-5 * 24 * time.Hour)},
			mockUpdatedAts: []time.Time{time.Now().Add(-4 * 24 * time.Hour), time.Now().Add(-4 * 24 * time.Hour)},
			mockDeletedAts: []time.Time{time.Now(), time.Now()},
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			mdb, mock, err := NewDbMock()
			if err != nil {
				t.Fatalf("an error '%s' was not expect when opening a stub database connection", err)
			}

			// mocking the results of db query
			mock.ExpectQuery(regexp.QuoteMeta(
				`SELECT * FROM "telephones"`)).
				WillReturnRows(NewMockRow(mock, tt.mockOwnerIds, tt.mockNumbers, tt.mockICCIds, tt.mockIds, tt.mockCreatedAts, tt.mockUpdatedAts, tt.mockDeletedAts))

			// execute the test fuction
			tr := telephoneRepository{
				connection: mdb,
			}
			tels, err := tr.ListTelephone()

			// assertion
			if len(tels) != len(tt.mockIds) {
				t.Errorf(`expect return length of struct "Telephone" was %d (actual: %d)`, len(tt.mockIds), len(tels))
				for _, tel := range tels {
					t.Errorf("%v", tel)
				}
			}
		})
	}
}

func TestPostTelephone(t *testing.T) {

	cases := []struct {
		description  string
		inputNumber  string
		inputOwnerId int
		inputICCId   int
		mockId       uint
		hasErr       bool
	}{
		{
			description:  "when correct arguments are given, it returns nil",
			inputNumber:  "09011112222",
			inputOwnerId: 1,
			inputICCId:   111111111111111,
			mockId:       1,
			hasErr:       false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			mdb, mock, err := NewDbMock()
			if err != nil {
				t.Fatalf("an error '%s' was not expect when opening a stub database connection", err)
			}

			// mocking the results of db query
			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.mockId)
			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(
				`INSERT INTO "telephones"`)).
				WillReturnRows(rows)
			mock.ExpectCommit()

			// execute the test function
			tr := telephoneRepository{
				connection: mdb,
			}
			err = tr.PostTelephone(tt.inputOwnerId, tt.inputICCId, tt.inputNumber)

			// assertion
			if !tt.hasErr && err != nil {
				t.Errorf(`an error '%s' was not expect`, err)
			}
		})
	}
}

func TestPutTelephoneByNumber(t *testing.T) {

	cases := []struct {
		description  string
		inputNumber  string
		inputOwnerId int
		inputICCId   int
		mockId       uint
		hasErr       bool
	}{
		{
			description:  "when correct arguments are given, it returns nil",
			inputNumber:  "09011112222",
			inputOwnerId: 1,
			inputICCId:   111111111111111,
			mockId:       1,
			hasErr:       true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			mdb, mock, err := NewDbMock()
			if err != nil {
				t.Fatalf("an error '%s' was not expect when opening a stub database connection", err)
			}

			// mocking the results of db query
			rows := sqlmock.NewRows([]string{"id"}).AddRow(tt.mockId)
			mock.ExpectBegin()
			mock.ExpectQuery(regexp.QuoteMeta(
				`UPDATE telephones SET (.+)`)).
				WithArgs(tt.inputICCId, tt.inputOwnerId, AnyTime{}, tt.inputNumber).
				WillReturnRows(rows)
			mock.ExpectCommit()

			// execute the test function
			tr := telephoneRepository{
				connection: mdb,
			}
			err = tr.PutTelephoneByNumber(tt.inputNumber, tt.inputOwnerId, tt.inputICCId)

			// assertion
			if !tt.hasErr && err != nil {
				t.Errorf(`an error '%s' was not expect`, err)
			}
		})
	}
}

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, mock, err
	}

	gormDB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			}), &gorm.Config{})

	return gormDB, mock, err
}

func NewMockRow(mock sqlmock.Sqlmock, ownerIds []int, numbers []string, iccIds []int, ids []uint, createdAts, updatedAts, deletedAts []time.Time) *sqlmock.Rows {
	rows := mock.NewRows([]string{
		"owner_id",
		"number",
		"icc_id",
		"id",
		"created_at",
		"updated_at",
		"deleted_at",
	})

	for i := 0; i < len(ids); i++ {
		rows.AddRow(
			ownerIds[i],
			numbers[i],
			iccIds[i],
			ids[i],
			createdAts[i],
			updatedAts[i],
			deletedAts[i],
		)
	}

	return rows
}
