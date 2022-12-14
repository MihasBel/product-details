package rest

import (
	"net/http"

	"github.com/MihasBel/product-details/model"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// getAll godoc
// @Summary Retrieves all product details
// @Produce json
// @Success 200 {array} Details
// @Router /details/ [get]
// @Security ApiKeyAuth
func (r *REST) getAll(c *fiber.Ctx) error {
	ds, err := r.d.All(c.Context())
	if err != nil {
		log.Error().Err(err).Msg("error while get all details from db")
		return err
	}
	if err := c.JSON(ds); err != nil {
		log.Error().Err(err).Msg("error while marshal all details")
		return err
	}
	return nil

}

// Get godoc
// @Summary Retrieves product-details on given ID
// @Produce json
// @Param id path string true "product details ID"
// @Success 200 {object} Details
// @Router /details/one/{id} [get]
// @Security ApiKeyAuth
func (r *REST) getByID(c *fiber.Ctx) error {
	ids := c.Params("id", "")
	if !primitive.IsValidObjectID(ids) {
		return fiber.NewError(http.StatusBadRequest, "wrong id format "+ids)
	}
	d, err := r.d.ByID(c.Context(), ids)
	if err != nil {
		log.Error().Err(err).Msg("error while getting one details by id")
		return err
	}
	if err := c.JSON(d); err != nil {
		log.Error().Err(err).Msg("error while marshal one details")
		return err
	}
	return nil
}

// Create godoc
// @Summary Creates a new product-details from the received json document
// @Accept json
// @Param request body model.Detail true "product-details schema"
// @Success 200 {object} model.Detail
// @Router /details/ [post]
// @Security ApiKeyAuth
func (r *REST) create(c *fiber.Ctx) error {
	d := model.Detail{}
	if err := c.BodyParser(&d); err != nil {
		log.Error().Err(err).Msg("error while decode details")
		return fiber.NewError(http.StatusBadRequest, "error while decode request body")
	}
	d, err := r.d.InsertOne(c.Context(), d) //TODO pointer?
	if err != nil {
		log.Error().Err(err).Msg("error while insert one details to db")
		return err
	}
	if err := c.JSON(d); err != nil {
		log.Error().Err(err).Msg("error while marshal inserted one details")
		return err
	}
	return nil
}
