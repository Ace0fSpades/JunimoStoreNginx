package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"uniStore/Backend/internal/domain/services"
	"uniStore/Backend/internal/interfaces/dto"
)

// UserHandler handles HTTP requests related to users
type UserHandler struct {
	userService services.UserService
	roleService services.RoleService
	authService services.AuthService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService services.UserService, roleService services.RoleService, authService services.AuthService) *UserHandler {
	return &UserHandler{
		userService: userService,
		roleService: roleService,
		authService: authService,
	}
}

// Register handles user registration
// @Summary Registers a new User
// @Description This endpoint allows you to register a new User by providing required fields
// @Tags Users
// @Accept json
// @Produce json
// @Param signup body dto.UserSignupDTO true "User data to register"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 409 {object} map[string]interface{} "Email or nickname already in use"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/signup [post]
func (h *UserHandler) Register(c *gin.Context) {
	var userDTO dto.UserSignupDTO
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Input validation
	if userDTO.Email == "" || userDTO.Nickname == "" || userDTO.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email, nickname and password are required"})
		return
	}

	// Validate password confirmation
	if userDTO.Password != userDTO.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "passwords do not match"})
		return
	}

	// Register user using DTO
	userResponseDTO, err := h.userService.Register(&userDTO)
	if err != nil {
		if err.Error() == "email already in use" || err.Error() == "nickname already in use" {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    userResponseDTO,
	})
}

// Login handles user login
// @Summary Logs in a User and returns user data
// @Description This endpoint allows the user to log in by providing email and password
// @Tags Users
// @Accept json
// @Produce json
// @Param login body dto.UserLoginDTO true "User credentials (email and password)"
// @Success 200 {object} dto.UserResponseDTO "Successfully logged in and returned user data"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Email or password is incorrect"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginDTO dto.UserLoginDTO
	if err := c.BindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login user with DTO
	authResponse, err := h.userService.Login(&loginDTO)
	if err != nil {
		if err.Error() == "email or password is incorrect" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authResponse)
}

// RefreshToken handles token refresh
// @Summary Refreshes an access token
// @Description This endpoint refreshes an access token using a refresh token
// @Tags Users
// @Accept json
// @Produce json
// @Param refresh_token body object true "Refresh token object" Schema(object,required=refresh_token,properties={refresh_token=string})
// @Success 200 {object} map[string]interface{} "Successfully refreshed token"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Invalid token"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/refresh [post]
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "refresh token is required"})
		return
	}

	// Refresh token using user service
	authResponse, err := h.userService.RefreshToken(requestBody.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Token refreshed successfully",
		"token":         authResponse.Token,
		"refresh_token": authResponse.RefreshToken,
	})
}

// GetUserByID handles getting a user by ID
// @Summary Get a User by ID
// @Description Fetches a User by their ID from the database
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} dto.UserResponseDTO "Successfully retrieved the user data"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Failed to fetch user"
// @Security ApiKeyAuth
// @Router /api/v1/users/{user_id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user is authorized to access this resource
	if err := h.authService.MatchUserTypeToID(id, c.GetString("role")); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	// Get user by ID
	userResponseDTO, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userResponseDTO)
}

// GetAllUsers handles getting all users with pagination
// @Summary Get users with pagination
// @Description Fetches a paginated list of users from the database
// @Tags Users
// @Accept json
// @Produce json
// @Param limit query int false "Limit number of results" default(10)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {array} dto.UserResponseDTO "Successfully retrieved the paginated users"
// @Failure 400 {object} map[string]interface{} "Invalid query parameters"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Failed to fetch users"
// @Security ApiKeyAuth
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	// Parse pagination parameters
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
		return
	}

	// Get all users with pagination
	users, err := h.userService.GetAllUsers(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser handles updating a user
// @Summary Update a User
// @Description Updates a User's information
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Param user body dto.UserUpdateDTO true "User data to update"
// @Success 200 {object} map[string]interface{} "User updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid input or user ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Failed to update user"
// @Security ApiKeyAuth
// @Router /api/v1/users/{user_id} [patch]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Check if the user is authorized to access this resource
	if err := h.authService.MatchUserTypeToID(id, c.GetString("role")); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	var userDTO dto.UserUpdateDTO
	if err := c.BindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Only admin can update role
	if c.GetString("role") != "admin" {
		userDTO.RoleID = 0 // Reset role ID if not admin
	}

	// Update user using DTO
	updatedUser, err := h.userService.UpdateUser(id, &userDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}
