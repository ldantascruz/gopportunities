package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ldantascruz/gopportunities/schemas"
	"net/http"
)

func ListOpeningsHandler(ctx *gin.Context) {
	var openings []schemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	var responses []schemas.OpeningResponse
	for _, opening := range openings {
		responses = append(responses, schemas.OpeningResponse{
			ID:        opening.ID,
			CreatedAt: opening.CreatedAt,
			UpdatedAt: opening.UpdatedAt,
			DeletedAt: opening.DeletedAt.Time, // Use .Time para extrair o valor do gorm.DeletedAt
			Role:      opening.Role,
			Company:   opening.Company,
			Location:  opening.Location,
			Remote:    opening.Remote,
			Link:      opening.Link,
			Salary:    opening.Salary,
		})
	}

	sendSuccess(ctx, "list-openings", responses)
}
