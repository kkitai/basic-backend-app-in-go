package http

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
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
			// TODO: assert about response body
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
