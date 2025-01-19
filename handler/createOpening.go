package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ldantascruz/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1

// CreateOpeningHandler handles the creation of a new job opening.
// @Summary Create Opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	err := ctx.BindJSON(&request)
	if err != nil {
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	response := schemas.OpeningResponse{
		ID:        opening.ID,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
		DeletedAt: opening.DeletedAt.Time,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
	}

	sendSuccess(ctx, "create-opening", response)
}
