package api

import (
	"github.com/gin-gonic/gin"

	"uniStore/Backend/internal/domain/services"
)

// LibraryHandler handles HTTP requests related to the user's game library
type LibraryHandler struct {
	libraryService services.LibraryService
}

// NewLibraryHandler creates a new library handler
func NewLibraryHandler(libraryService services.LibraryService) *LibraryHandler {
	return &LibraryHandler{
		libraryService: libraryService,
	}
}

// GetLibrary retrieves a user's game library
// @Summary Get user's game library
// @Description Returns the list of games owned by the user
// @Tags Library
// @Accept json
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} dto.GameDTO "List of games in the user's library"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 403 {object} map[string]interface{} "Forbidden"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security ApiKeyAuth
// @Router /api/v1/library/{user_id} [get]
func (h *LibraryHandler) GetLibrary(c *gin.Context) {
	// Implementation to be added
}
