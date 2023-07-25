package api

import (
	mockdb "barbershop/db/mock"
	db "barbershop/db/sqlc"
	"barbershop/db/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
	"gopkg.in/guregu/null.v4"
)

func TestGetBarberWithName(t *testing.T) {
	barber := createBarber()

	testCases := []struct {
		name          string
		nameBarber    null.String
		buildStubs    func(store *mockdb.MockStoreMain)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:       "OK",
			nameBarber: null.StringFrom(barber.Username),
			buildStubs: func(store *mockdb.MockStoreMain) {
				store.EXPECT().
					GetEmployeeWithName(gomock.Any(), gomock.Eq(barber.Username)).
					Times(1).
					Return(barber, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder.Body, barber)
			},
		},
		{
			name:       "Not Found",
			nameBarber: null.StringFrom(barber.Username),
			buildStubs: func(store *mockdb.MockStoreMain) {
				store.EXPECT().
					GetEmployeeWithName(gomock.Any(), gomock.Eq(barber.Username)).
					Times(1).
					Return(db.Barber{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:       "InternalError",
			nameBarber: null.StringFrom(barber.Username),
			buildStubs: func(store *mockdb.MockStoreMain) {
				store.EXPECT().
					GetEmployeeWithName(gomock.Any(), gomock.Eq(barber.Username)).
					Times(1).
					Return(db.Barber{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:       "InvalidID",
			nameBarber: null.StringFrom("a"),
			buildStubs: func(store *mockdb.MockStoreMain) {
				store.EXPECT().
					GetEmployeeWithName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {

		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStoreMain(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/employee/%s", tc.nameEmployee.String)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStoreMain(ctrl)
	store.EXPECT().
		GetEmployeeWithName(gomock.Any(), gomock.Eq(barber.Username)).
		Times(1).
		Return(barber, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()
	url := fmt.Sprintf("/employee/%s", barber.Username)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)
	require.Equal(t, http.StatusOK, recorder.Code)

	var book map[string]json.RawMessage

	errs := json.Unmarshal(recorder.Body.Bytes(), &book)
	require.NoError(t, errs)
	fmt.Print(string(book["barber"]))

	requireBodyMatchAccount(t, recorder.Body, barber)

}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, employee db.Barber) {
	value := gjson.Get(body.String(), "employee")

	data := []byte(value.Raw)
	var gotEmployee db.Barber
	err := json.Unmarshal(data, &gotEmployee)

	require.NoError(t, err)
	require.Equal(t, employee, gotEmployee)
}

func createBarber() db.Barber {
	return db.Barber{
		FullName:       util.RandomName(),
		Email:          util.RandomEmail(),
		HashedPassword: util.RandomName(),
		Avatar:         null.StringFrom(util.RandomName()),
		Status:         null.StringFrom(util.RandomString(1)),
		StoreManager:   []uuid.UUID{},
		Username:       util.RandomName(),
		Role:           null.StringFrom(util.RandomString(1)),
		StoreID:        uuid.NullUUID{},
	}
}
