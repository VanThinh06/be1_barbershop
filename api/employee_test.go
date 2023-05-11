package api

// import (
// 	mockdb "barbershop/db/mock"
// 	db "barbershop/db/sqlc"
// 	"barbershop/db/util"
// 	"bytes"
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// 	"github.com/tidwall/gjson"
// 	"gopkg.in/guregu/null.v4"
// )

// func TestGetEmployeeWithName(t *testing.T) {
// 	employee := createEmployee()

// 	testCases := []struct {
// 		name          string
// 		nameEmployee  null.String
// 		buildStubs    func(store *mockdb.MockStoreMain)
// 		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name:         "OK",
// 			nameEmployee: null.StringFrom(employee.Username),
// 			buildStubs: func(store *mockdb.MockStoreMain) {
// 				store.EXPECT().
// 					GetEmployeeWithName(gomock.Any(), gomock.Eq(employee.Username)).
// 					Times(1).
// 					Return(employee, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, employee)
// 			},
// 		},
// 		{
// 			name:         "Not Found",
// 			nameEmployee: null.StringFrom(employee.Username),
// 			buildStubs: func(store *mockdb.MockStoreMain) {
// 				store.EXPECT().
// 					GetEmployeeWithName(gomock.Any(), gomock.Eq(employee.Username)).
// 					Times(1).
// 					Return(db.Employee{}, sql.ErrNoRows)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusNotFound, recorder.Code)
// 			},
// 		},
// 		{
// 			name:         "InternalError",
// 			nameEmployee: null.StringFrom(employee.Username),
// 			buildStubs: func(store *mockdb.MockStoreMain) {
// 				store.EXPECT().
// 					GetEmployeeWithName(gomock.Any(), gomock.Eq(employee.Username)).
// 					Times(1).
// 					Return(db.Employee{}, sql.ErrConnDone)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusInternalServerError, recorder.Code)
// 			},
// 		},
// 		// {
// 		// 	name:         "InvalidID",
// 		// 	nameEmployee: null.StringFrom("a"),
// 		// 	buildStubs: func(store *mockdb.MockStoreMain) {
// 		// 		store.EXPECT().
// 		// 			GetEmployeeWithName(gomock.Any(), gomock.Any()).
// 		// 			Times(0)
// 		// 	},
// 		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)
// 		// 	},
// 		// },
// 	}

// 	for i := range testCases {

// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			store := mockdb.NewMockStoreMain(ctrl)
// 			tc.buildStubs(store)

// 			server := NewServer(store)
// 			recorder := httptest.NewRecorder()

// 			url := fmt.Sprintf("/employee/%s", tc.nameEmployee.String)
// 			request, err := http.NewRequest(http.MethodGet, url, nil)
// 			require.NoError(t, err)

// 			server.router.ServeHTTP(recorder, request)
// 			tc.checkResponse(t, recorder)
// 		})
// 	}

// 	// ctrl := gomock.NewController(t)
// 	// defer ctrl.Finish()

// 	// store := mockdb.NewMockStoreMain(ctrl)
// 	// store.EXPECT().
// 	// 	GetEmployeeWithName(gomock.Any(), gomock.Eq(employee.Username)).
// 	// 	Times(1).
// 	// 	Return(employee, nil)

// 	// server := NewServer(store)
// 	// recorder := httptest.NewRecorder()
// 	// url := fmt.Sprintf("/employee/%s", employee.Username)
// 	// request, err := http.NewRequest(http.MethodGet, url, nil)
// 	// require.NoError(t, err)
// 	// server.router.ServeHTTP(recorder, request)
// 	// require.Equal(t, http.StatusOK, recorder.Code)

// 	// var book map[string]json.RawMessage

// 	// errs := json.Unmarshal(recorder.Body.Bytes(), &book)
// 	// require.NoError(t, errs)
// 	// fmt.Print(string(book["employee"]))

// 	// requireBodyMatchAccount(t, recorder.Body, employee)

// }

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, employee db.Employee) {
// 	value := gjson.Get(body.String(), "employee")

// 	data := []byte(value.Raw)
// 	var gotEmployee db.Employee
// 	err := json.Unmarshal(data, &gotEmployee)

// 	require.NoError(t, err)
// 	require.Equal(t, employee, gotEmployee)
// }

// func createEmployee() db.Employee {
// 	return db.Employee{
// 		ID:       uuid.New(),
// 		Username: util.RandomName(),
// 		Role:     util.RandomRole(),
// 		Image:    null.String{},
// 		StoreID:  uuid.Nil,
// 	}
// }
