package handler

import (
	"faceit-test/internal/domain/user/service"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"log"
	"net/http"
)

const parseRequestErrText = "unable to parse request"
const noIDErrText = "no id provided"

type User struct {
	userService service.User
}

func NewUser(userService service.User) User {
	return User{userService: userService}
}

func (h *User) Create(ctx echo.Context) error {
	var user userModel
	err := ctx.Bind(&user)
	if err != nil {
		log.Println("failed to unmarshal user")
		return ctx.JSON(http.StatusBadRequest, mapErrorToErrorResponse(errors.New(parseRequestErrText)))
	}
	userID, err := h.userService.CreateUser(ctx.Request().Context(), mapToUserModel(user))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, mapErrorToErrorResponse(err))
	}
	return ctx.JSON(http.StatusCreated, userModel{ID: userID})
}

func (h *User) Update(ctx echo.Context) error {
	var user userModel
	err := ctx.Bind(&user)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, mapErrorToErrorResponse(errors.New(parseRequestErrText)))
	}
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, mapErrorToErrorResponse(errors.New(noIDErrText)))
	}
	err = h.userService.UpdateUser(ctx.Request().Context(), mapUpdateToUserModel(user, id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, mapErrorToErrorResponse(err))
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (h *User) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	if id == "" {
		return ctx.JSON(http.StatusBadRequest, mapErrorToErrorResponse(errors.New(noIDErrText)))
	}
	err := h.userService.DeleteUser(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, mapErrorToErrorResponse(err))
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (h *User) Paged(ctx echo.Context) error {
	pagedUsersRequest := mapToPagedUsersRequest(ctx.QueryParams())
	users, count, err := h.userService.PaginatedUsers(ctx.Request().Context(), pagedUsersRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, mapErrorToErrorResponse(err))
	}
	return ctx.JSON(http.StatusOK, mapToPagedUsersResponse(users, count))
}
