package interfaces

import (
	"encoding/binary"
	"github.com/Lenstack/clean-architecture/internal/domain"
	"github.com/Lenstack/clean-architecture/internal/usecases"
	"github.com/Lenstack/clean-architecture/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserInteractor usecases.UserInteractor
	Logger         usecases.Logger
}

func NewUserHandler(logger usecases.Logger, mongo usecases.Mongo) *UserHandler {
	return &UserHandler{
		UserInteractor: usecases.UserInteractor{
			UserRepository: &UserRepository{
				Mongo: mongo,
			},
		},
		Logger: logger,
	}
}

func (uc *UserHandler) Index(ctx *fiber.Ctx) error {
	users, err := uc.UserInteractor.Index()
	if err != nil {
		uc.Logger.LogError("%s", err)
	}
	if len(users) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(
			domain.Response{
				Status:  fiber.StatusNotFound,
				Message: "Error",
				Data:    "No Users Found.",
			},
		)
	}
	return ctx.Status(fiber.StatusFound).JSON(users)
}

func (uc *UserHandler) Show(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	user, err := uc.UserInteractor.Show(userId)

	if err != nil {
		if binary.Size(user) == -1 {
			return ctx.Status(fiber.StatusNotFound).JSON(
				domain.Response{
					Status:  fiber.StatusNotFound,
					Message: "Error",
					Data:    "No User Found.",
				},
			)
		}
	}

	return ctx.Status(fiber.StatusFound).JSON(user)
}

func (uc *UserHandler) Create(ctx *fiber.Ctx) error {
	var body domain.User
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	errValidate := utils.ValidateStruct(body)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			domain.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Error",
				Data:    errValidate,
			},
		)
	}

	_, err := uc.UserInteractor.Store(body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			domain.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Error",
				Data:    "The User Is Not Been Created.",
			},
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		domain.Response{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    "The User Has Been Created.",
		},
	)

}

func (uc *UserHandler) Update(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var body domain.User
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	errValidate := utils.ValidateStruct(body)
	if errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			domain.Response{
				Status:  fiber.StatusBadRequest,
				Message: "Error",
				Data:    errValidate,
			},
		)
	}

	_, err := uc.UserInteractor.Update(userId, body)
	if err != nil {
		return ctx.Status(fiber.StatusNotModified).JSON(
			domain.Response{
				Status:  fiber.StatusNotModified,
				Message: "Error",
				Data:    "The User Is Not Been Updated.",
			},
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		domain.Response{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    "The User Has Been Updated.",
		},
	)
}

func (uc *UserHandler) Destroy(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	
	_, err := uc.UserInteractor.Destroy(userId)
	if err != nil {
		return ctx.Status(fiber.StatusNotModified).JSON(
			domain.Response{
				Status:  fiber.StatusNotModified,
				Message: "Error",
				Data:    "The User Is Not Been Deleted.",
			},
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		domain.Response{
			Status:  fiber.StatusOK,
			Message: "Success",
			Data:    "The User Has Been Deleted.",
		},
	)
}
