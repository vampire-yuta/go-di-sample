package handlers

import (
	"net/http"

	"go-di/internal/usecase"

	"github.com/labstack/echo/v4"
)

// UserHandler はユーザー関連のHTTPハンドラーです
type UserHandler struct {
	userService *usecase.UserService
}

// NewUserHandler は新しいUserHandlerを作成します
func NewUserHandler(userService *usecase.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser はユーザー作成のハンドラーです
func (h *UserHandler) CreateUser(c echo.Context) error {
	var req usecase.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, user)
}

// GetUser はユーザー取得のハンドラーです
func (h *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User ID is required",
		})
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

// GetAllUsers は全ユーザー取得のハンドラーです
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"count": len(users),
	})
}

// UpdateUser はユーザー更新のハンドラーです
func (h *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User ID is required",
		})
	}

	var req usecase.UpdateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// URLパラメータのIDを設定
	req.ID = id

	user, err := h.userService.UpdateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, user)
}

// DeleteUser はユーザー削除のハンドラーです
func (h *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "User ID is required",
		})
	}

	err := h.userService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.NoContent(http.StatusNoContent)
}

// HealthCheck はヘルスチェックのハンドラーです
func (h *UserHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Server is running",
	})
}
