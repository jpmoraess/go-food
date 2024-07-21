package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpmoraess/go-food/order-service/internal/application/dto"
	"github.com/jpmoraess/go-food/order-service/internal/application/usecase"
)

type OrderHandler struct {
	createOrderUseCase *usecase.CreateOrderUseCase
}

func NewOrderHandler(createOrderUseCase *usecase.CreateOrderUseCase) *OrderHandler {
	return &OrderHandler{createOrderUseCase: createOrderUseCase}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var input = new(dto.CreateOrderInputDTO)
	if err := c.BodyParser(input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	output, err := h.createOrderUseCase.Execute(c.Context(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusCreated).JSON(output)
}
