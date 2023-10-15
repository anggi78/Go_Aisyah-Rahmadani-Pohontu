package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"task/usecase"
)

type Handler struct {
	LaptopUsecase usecase.Usecase
}

type LaptopResponse struct {
	Status string `json:"status"`
	Data   string `json:"data"`
}

func NewLaptopHandler(usecase usecase.Usecase) *Handler {
	return &Handler{
		LaptopUsecase: usecase,
	}
}

func (h *Handler) RecommendLaptop(c echo.Context) error {
	var requestData map[string]interface{}
	err := c.Bind(&requestData)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON format")
	}

	budget, okBudget := requestData["budget"].(float64)
	purpose, okPurpose := requestData["purpose"].(string)
	brand, okBrand := requestData["brand"].(string)
	ram, okRAM := requestData["ram"].(string)
	cpu, okCPU := requestData["cpu"].(string)
	screenSize, okScreenSize := requestData["screen_size"].(string)

	if !okBudget || !okPurpose || !okBrand || !okRAM || !okCPU || !okScreenSize {
		return c.JSON(http.StatusBadRequest, "Invalid request format")
	}

	userInput := fmt.Sprintf("Recommendation of a %s laptop for %s with %s RAM, %s CPU, %s screen size, and a budget of Rp %.0f.", brand, purpose, ram, cpu, screenSize, budget)

	value, err := h.LaptopUsecase.Recommend(userInput, brand, ram, cpu, screenSize, os.Getenv("KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error in making recommendations")
	}

	responseData := LaptopResponse{
		Status: "success",
		Data:   value,
	}

	return c.JSON(http.StatusOK, responseData)
}