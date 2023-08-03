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

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
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
					GetBarber(gomock.Any(), gomock.Eq(barber.Username)).
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
					GetBarber(gomock.Any(), barber.Username).
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
					GetBarber(gomock.Any(), gomock.Eq(barber.Username)).
					Times(1).
					Return(db.Barber{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		// {
		// 	name:       "InvalidID",
		// 	nameBarber: null.StringFrom("a"),
		// 	buildStubs: func(store *mockdb.MockStoreMain) {
		// 		store.EXPECT().
		// 			GetBarber(gomock.Any(), gomock.Any()).
		// 			Times(0)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
		// 	},
		// },
	}

	for i := range testCases {

		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storeMain := mockdb.NewMockStoreMain(ctrl)
			tc.buildStubs(storeMain)

			server := NewTestServer(t, storeMain)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/barber/%s", tc.nameBarber.String)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, barber db.Barber) {
	// value := gjson.Get(body.String(), "barber")
	// log.Printf("body barber %s", value)
	// log.Printf("body barber %s", value.Raw)
	// data := []byte(value.Raw)
	var data []byte = body.Bytes()
	var getBarber db.Barber
	err := json.Unmarshal(data, &getBarber)

	require.NoError(t, err)
	require.Equal(t, barber, getBarber)
}

func createBarber() db.Barber {
	return db.Barber{
		FullName:     util.RandomName(),
		Email:        util.RandomEmail(),
		Avatar:       null.StringFrom(util.RandomName()),
		Status:       null.StringFrom(util.RandomString(1)),
		StoreManager: []uuid.UUID{},
		Username:     util.RandomName(),
		Role:         null.StringFrom(util.RandomString(1)),
		StoreWork:    uuid.NullUUID{},
	}
}
