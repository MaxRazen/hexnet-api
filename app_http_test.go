package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"hexnet/api/auth"
	"hexnet/api/common"
	"hexnet/api/notes"
	"hexnet/api/users"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	server    *gin.Engine
	authToken string
	user      *users.UserModel
	userData  *users.UserCreateData
)

func prepareAppForTest() common.Config {
	gin.SetMode(gin.TestMode)

	config := common.LoadConfig("")
	common.InitTestDbConnection()
	common.RegisterCustomValidationRules()
	migrate()
	createTestUser()

	server = setupServer()

	return config
}

var authEndpointTests = []struct {
	path         string
	bodyData     string
	expectedCode int
	responseData string
}{
	// bad payload request
	{
		path:         "/api/auth/authorize",
		bodyData:     `{"login":""`,
		expectedCode: http.StatusBadRequest,
		responseData: "",
	},
	// validation error
	{
		path:         "/api/auth/authorize",
		bodyData:     "{}",
		expectedCode: http.StatusUnprocessableEntity,
		responseData: `{"message":"Validation error","errors":[{"field":"login","message":"required"},{"field":"password","message":"required"}]}`,
	},
	// unknown user
	{
		path:         "/api/auth/authorize",
		bodyData:     `{"login":"john","password":"some-password"}`,
		expectedCode: http.StatusUnprocessableEntity,
		responseData: `{"message":"User not found or password mismatch"}`,
	},
	// wrong password
	{
		path:         "/api/auth/authorize",
		bodyData:     `{"login": "john.doe", "password": "wrong-password"}`,
		expectedCode: http.StatusUnprocessableEntity,
		responseData: `{"message":"User not found or password mismatch"}`,
	},
	// success case
	{
		path:         "/api/auth/authorize",
		bodyData:     `{"login": "john.doe", "password": "pa$w0rd"}`,
		expectedCode: http.StatusOK,
	},
}

func TestAuthEndpoint(t *testing.T) {
	prepareAppForTest()
	asserts := assert.New(t)

	for _, testCase := range authEndpointTests {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodPost, testCase.path, bytes.NewBufferString(testCase.bodyData))
		req.Header.Set("Content-Type", "application/json")
		server.ServeHTTP(w, req)

		asserts.NoError(err)
		asserts.Equal(testCase.expectedCode, w.Code)
		if w.Code == http.StatusOK {
			responseData := &auth.AuthorizeResponseData{}
			_ = json.Unmarshal(w.Body.Bytes(), responseData)

			asserts.NotEmpty(responseData.Jwt)
			asserts.Greater(responseData.Exp, time.Now().Add(time.Minute).Unix())
		} else {
			asserts.Equal(testCase.responseData, w.Body.String())
		}
	}
}

var usersMeEndpointTests = []struct {
	path         string
	expectedCode int
	responseData string
	getToken     func() string
}{
	// unauthorized request
	{
		path:         "/api/users/me",
		expectedCode: http.StatusUnauthorized,
		getToken:     func() string { return "" },
	},
	// request with bad auth token
	{
		path:         "/api/users/me",
		expectedCode: http.StatusUnauthorized,
		getToken:     func() string { return "bad-token" },
	},
	// successful authorized request
	{
		path:         "/api/users/me",
		expectedCode: http.StatusOK,
		getToken:     func() string { return getAuthorizeToken() },
	},
}

func TestUsersMeEndpoint(t *testing.T) {
	prepareAppForTest()
	asserts := assert.New(t)

	for _, testCase := range usersMeEndpointTests {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, testCase.path, bytes.NewReader([]byte{}))
		req.Header.Set("Content-Type", "application/json")
		token := testCase.getToken()
		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
		server.ServeHTTP(w, req)

		asserts.NoError(err)
		asserts.Equal(testCase.expectedCode, w.Code)
	}
}

var notesEndpointsTest = []struct {
	path         string
	method       string
	expectedCode int
	requestBody  string
	responseData string
	authorized   bool
}{
	// unauthorized request
	{
		path:         "/api/notes",
		method:       http.MethodPost,
		expectedCode: http.StatusUnauthorized,
		requestBody:  "",
		responseData: `{"code":401,"message":"auth header is empty"}`,
		authorized:   false,
	},
	// request with bad payload
	{
		path:         "/api/notes",
		method:       http.MethodPost,
		expectedCode: http.StatusBadRequest,
		requestBody:  `-`,
		responseData: "",
		authorized:   true,
	},
	// validation error
	{
		path:         "/api/notes",
		method:       http.MethodPost,
		expectedCode: http.StatusUnprocessableEntity,
		requestBody:  `{"note":"text","title":"-"}`,
		responseData: `{"message":"Validation error","errors":[{"field":"title","message":"length"},{"field":"content","message":"required"}]}`,
		authorized:   true,
	},
	// successful request
	{
		path:         "/api/notes",
		method:       http.MethodPost,
		expectedCode: http.StatusCreated,
		requestBody:  `{"title":"::title::","content":"::content::"}`,
		responseData: `{"id":1,"title":"::title::","content":"::content::"}`,
		authorized:   true,
	},
}

func TestNotesEndpoints(t *testing.T) {
	prepareAppForTest()
	asserts := assert.New(t)

	for _, testCase := range notesEndpointsTest {
		w := httptest.NewRecorder()
		req, err := http.NewRequest(testCase.method, testCase.path, bytes.NewBufferString(testCase.requestBody))
		req.Header.Set("Content-Type", "application/json")
		if testCase.authorized {
			req.Header.Set("Authorization", "Bearer "+getAuthorizeToken())
		}
		server.ServeHTTP(w, req)

		asserts.NoError(err)
		asserts.Equal(testCase.expectedCode, w.Code)
		if testCase.expectedCode == http.StatusCreated {
			d := &notes.NoteModel{}
			r := &notes.NoteModel{}
			_ = json.Unmarshal(w.Body.Bytes(), d)
			_ = json.Unmarshal(w.Body.Bytes(), r)
			asserts.NotEmpty(r.CreatedAt)
			asserts.Equal(r.CreatedAt, r.UpdatedAt)
			d.CreatedAt, d.UpdatedAt = r.CreatedAt, r.UpdatedAt
			asserts.Equal(d, r)
		} else {
			asserts.Equal(testCase.responseData, w.Body.String())
		}
	}
}

func createTestUser() {
	data := users.UserCreateData{
		Login:    "john.doe",
		Name:     "John Doe",
		Password: "pa$w0rd",
	}
	userData = &data

	user, _ = users.CreateUserAction(data)
}

func getAuthorizeToken() string {
	if authToken != "" {
		return authToken
	}

	m := auth.NewAuthMiddleware()
	payload := map[string]uint{
		"sub": user.ID,
	}
	token, _, err := m.TokenGenerator(payload)
	if err != nil {
		panic("Auth Token could not be generated")
	}

	authToken = token
	return authToken
}
