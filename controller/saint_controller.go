package controller

import (
	"math/rand"
	"net/http"
	"saints-api/model"
	"saints-api/usecase"
	"strconv"

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

	saints, err := p.saintUseCase.GetSaints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, saints)
}

func (p *saintController) CreateSaint(ctx *gin.Context) {

	var saint model.Saint
	err := ctx.BindJSON(&saint)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertedSaint, err := p.saintUseCase.CreateSaint(saint)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedSaint)
}

func (p *saintController) GetSaintById(ctx *gin.Context) {

	id := ctx.Param("saintId")
	if id == "" {
		response := model.Response{
			Message: "Saint Id can't be null",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	saintId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Saint Id must be a number",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	saint, err := p.saintUseCase.GetSaintById(saintId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if saint == nil {
		response := model.Response{
			Message: "Saint wasn't found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, saint)
}

func (p *saintController) GetRandomSaint(ctx *gin.Context) {

	saints, err := p.saintUseCase.GetSaints()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if len(saints) == 0 {
		response := model.Response{
			Message: "No saints found",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	randomIndex := rand.Intn(len(saints))
	ctx.JSON(http.StatusOK, saints[randomIndex])
}
