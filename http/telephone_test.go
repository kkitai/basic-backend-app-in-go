package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	mock_db "github.com/kkitai/basic-backend-app-in-go/mock/db"
	"github.com/kkitai/basic-backend-app-in-go/model"
)

func TestGetTelephone(t *testing.T) {
	cases := []struct {
		description      string
		inputNumber      string
		mockTelephone    *model.Telephone
		mockError        error
		expectStatusCode int
	}{
		{
			description: "when correct number and Telephone{} was given, it returns 200 response",
			inputNumber: "08033332222",
			mockTelephone: &model.Telephone{
				OwnerId: 0,
				Number:  "08033332222",
				ICCId:   0,
			},
			mockError:        nil,
			expectStatusCode: 200,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			// mocking dependent methods
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_db.NewMockTelephoneAccesscer(ctrl)
			mock.
				EXPECT().
				GetTelephoneByNumber(gomock.Eq(tt.inputNumber)).
				Return(tt.mockTelephone, tt.mockError)

			telephoneRepository = mock

			reqBody := bytes.NewBufferString("")
			req := httptest.NewRequest(http.MethodGet, "/telephones/{number}", reqBody)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("number", tt.inputNumber)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			got := httptest.NewRecorder()

			// exec
			getTelephone(got, req)

			if got.Code != tt.expectStatusCode {
				t.Errorf(`expectStatusCode(%d) was not matched (actual: %d)`, tt.expectStatusCode, got.Code)
			}
		})
	}
}

func TestListTelephone(t *testing.T) {
	cases := []struct {
		description      string
		mockTelephone    []*model.Telephone
		mockError        error
		expectStatusCode int
	}{
		{
			description: "when correct []Telephone{} was given, it returns 200 response",
			mockTelephone: []*model.Telephone{
				&model.Telephone{
					OwnerId: 0,
					Number:  "08033332222",
					ICCId:   0,
				},
				&model.Telephone{
					OwnerId: 1,
					Number:  "08033334444",
					ICCId:   1,
				},
			},
			mockError:        nil,
			expectStatusCode: 200,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			// mocking dependent methods
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := mock_db.NewMockTelephoneAccesscer(ctrl)
			mock.
				EXPECT().
				ListTelephone().
				Return(tt.mockTelephone, tt.mockError)

			telephoneRepository = mock

			reqBody := bytes.NewBufferString("")
			req := httptest.NewRequest(http.MethodGet, "/telephones", reqBody)
			rctx := chi.NewRouteContext()
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			got := httptest.NewRecorder()

			// exec
			listTelephone(got, req)

			if got.Code != tt.expectStatusCode {
				t.Errorf(`expectStatusCode(%d) was not matched (actual: %d)`, tt.expectStatusCode, got.Code)
			}
		})
	}
}

func TestPostTelephone(t *testing.T) {
	cases := []struct {
		description      string
		inputNumber      string
		inputOwnerId     string
		inputIccId       string
		mockError        error
		expectStatusCode int
	}{
		{
			description:      "when correct inputs was given, it returns 201 response",
			inputNumber:      "08033332222",
			inputOwnerId:     "1",
			inputIccId:       "1",
			mockError:        nil,
			expectStatusCode: http.StatusCreated,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			// mocking dependent methods
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ownerId, _ := strconv.Atoi(tt.inputOwnerId)
			iccId, _ := strconv.Atoi(tt.inputIccId)

			mock := mock_db.NewMockTelephoneAccesscer(ctrl)
			mock.
				EXPECT().
				PostTelephone(ownerId, iccId, tt.inputNumber).
				Return(tt.mockError)

			telephoneRepository = mock

			reqBody := bytes.NewBufferString(fmt.Sprintf(`{"owner_id": "%s", "icc_id": "%s"}`, tt.inputOwnerId, tt.inputIccId))
			req := httptest.NewRequest(http.MethodGet, "/telephones/{number}", reqBody)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("number", tt.inputNumber)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			got := httptest.NewRecorder()

			// exec
			postTelephone(got, req)

			if got.Code != tt.expectStatusCode {
				t.Errorf(`expectStatusCode(%d) was not matched (actual: %d)`, tt.expectStatusCode, got.Code)
			}
		})
	}
}

func TestPutTelephone(t *testing.T) {
	cases := []struct {
		description      string
		inputNumber      string
		inputOwnerId     string
		inputIccId       string
		mockTelephone    *model.Telephone
		mockError        error
		expectStatusCode int
	}{
		{
			description:  "when correct inputs was given, it returns 200 response",
			inputNumber:  "08033332222",
			inputOwnerId: "1",
			inputIccId:   "1",
			mockTelephone: &model.Telephone{
				OwnerId: 0,
				Number:  "08033332222",
				ICCId:   0,
			},
			mockError:        nil,
			expectStatusCode: http.StatusOK,
		},
	}

	for _, tt := range cases {
		t.Run(tt.description, func(t *testing.T) {
			// mocking dependent methods
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ownerId, _ := strconv.Atoi(tt.inputOwnerId)
			iccId, _ := strconv.Atoi(tt.inputIccId)

			mock := mock_db.NewMockTelephoneAccesscer(ctrl)
			mock.
				EXPECT().
				GetTelephoneByNumber(tt.inputNumber).
				Return(tt.mockTelephone, tt.mockError)
			mock.
				EXPECT().
				PutTelephoneByNumber(tt.inputNumber, ownerId, iccId).
				Return(tt.mockError)

			telephoneRepository = mock

			reqBody := bytes.NewBufferString(fmt.Sprintf(`{"owner_id": "%s", "icc_id": "%s"}`, tt.inputOwnerId, tt.inputIccId))
			req := httptest.NewRequest(http.MethodGet, "/telephones/{number}", reqBody)
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("number", tt.inputNumber)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))
			got := httptest.NewRecorder()

			// exec
			putTelephone(got, req)

			if got.Code != tt.expectStatusCode {
				t.Errorf(`expectStatusCode(%d) was not matched (actual: %d)`, tt.expectStatusCode, got.Code)
			}
		})
	}
}
