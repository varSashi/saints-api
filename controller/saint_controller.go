package controller

import (
	"saints-api/model"
	"saints-api/usecase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type saintController struct {
	saintUseCase usecase.SaintUsecase
}

func NewSaintController(usecase usecase.SaintUsecase) saintController {
	return saintController{
		saintUseCase: usecase,
	}
}

func (p *saintController) GetSaints(ctx *gin.Context) {

	saints := []model.Saint{
		{
			ID:    1,
			Name:  "St Augustine of Hippo",
			Quote: "God loves each of us as if there were only one of us",
		},
	}

	ctx.JSON(http.StatusOK, saints)
}
