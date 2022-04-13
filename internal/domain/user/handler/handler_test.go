package handler

import (
	"faceit-test/internal/domain/user/model"
	mockService "faceit-test/internal/domain/user/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

var (
	userJSON = `{
  "first_name": "firstname1",
  "last_name": "lastname",
  "nick_name": "nick1",
  "password": "pwd1",
  "email": "email",
  "country": "country"
}`
	respJSON = `{"id":"id"}
`
	paginatedJSON = `{"data":[{"id":"id","first_name":"firstname","last_name":"lastname","nick_name":"nickname","password":"password","email":"email","country":"country","created_at":1231411,"updated_at":12311111}],"Page":{"Total":1,"Count":1}}
`
	createUserModel = model.User{
		FirstName: "firstname1",
		LastName:  "lastname",
		NickName:  "nick1",
		Password:  "pwd1",
		Email:     "email",
		Country:   "country",
	}
	updateUserModel = model.User{
		ID:        "id1",
		FirstName: "firstname1",
		LastName:  "lastname",
		NickName:  "nick1",
		Password:  "pwd1",
		Email:     "email",
		Country:   "country",
	}
	paginatedRequest = model.PagedUserRequest{
		Nick:    "nick",
		Country: "country",
		Email:   "email",
		Page:    3,
		Size:    15,
	}
	paginatedResp = []*model.User{{
		ID:        "id",
		FirstName: "firstname",
		LastName:  "lastname",
		NickName:  "nickname",
		Password:  "password",
		Email:     "email",
		Country:   "country",
		CreatedAt: 1231411,
		UpdatedAt: 12311111,
	}}
	id         = "id1"
	errService = errors.New("service error")
	errJSON    = `{"error":"service error"}
`
	errParseJSON = `{"error":"unable to parse request"}
`
	errNoIDJSON = `{"error":"no id provided"}
`
)

func TestUser_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	userService := mockService.NewMockUser(ctrl)
	handler := NewUser(userService)

	t.Run("failed to unmarshal user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("h"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, errParseJSON, rec.Body.String())
	})

	t.Run("failed to create user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		userService.EXPECT().CreateUser(gomock.Any(), createUserModel).Return("", errService)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, errJSON, rec.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		userService.EXPECT().CreateUser(gomock.Any(), createUserModel).Return("id", nil)

		err := handler.Create(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, respJSON, rec.Body.String())
	})

}

func TestUser_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	userService := mockService.NewMockUser(ctrl)
	handler := NewUser(userService)

	t.Run("failed to unmarshal update user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader("h"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)

		err := handler.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, errParseJSON, rec.Body.String())
	})

	t.Run("failed no id provided", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")

		err := handler.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, errNoIDJSON, rec.Body.String())
	})

	t.Run("failed to update user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)
		userService.EXPECT().UpdateUser(gomock.Any(), updateUserModel).Return(errService)

		err := handler.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, errJSON, rec.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)
		userService.EXPECT().UpdateUser(gomock.Any(), updateUserModel).Return(nil)

		err := handler.Update(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "null\n", rec.Body.String())
	})
}

func TestUser_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	userService := mockService.NewMockUser(ctrl)
	handler := NewUser(userService)

	t.Run("failed no id provided", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")

		err := handler.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, errNoIDJSON, rec.Body.String())
	})

	t.Run("failed to delete user", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)
		userService.EXPECT().DeleteUser(gomock.Any(), id).Return(errService)

		err := handler.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, errJSON, rec.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(id)
		userService.EXPECT().DeleteUser(gomock.Any(), id).Return(nil)

		err := handler.Delete(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "null\n", rec.Body.String())
	})
}

func TestUser_Paged(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	e := echo.New()
	userService := mockService.NewMockUser(ctrl)
	handler := NewUser(userService)

	t.Run("failed to get paged users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		userService.EXPECT().PaginatedUsers(gomock.Any(), model.PagedUserRequest{}).Return(nil, 0, errService)

		err := handler.Paged(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, errJSON, rec.Body.String())
	})

	t.Run("success", func(t *testing.T) {
		params := make(url.Values)
		params.Add("email", paginatedRequest.Email)
		params.Add("nick", paginatedRequest.Nick)
		params.Add("country", paginatedRequest.Country)
		params.Add("page", strconv.Itoa(paginatedRequest.Page))
		params.Add("size", strconv.Itoa(paginatedRequest.Size))
		req := httptest.NewRequest(http.MethodGet, "/?"+params.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		userService.EXPECT().PaginatedUsers(gomock.Any(), paginatedRequest).Return(paginatedResp, 1, nil)

		err := handler.Paged(c)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, paginatedJSON, rec.Body.String())
	})
}
