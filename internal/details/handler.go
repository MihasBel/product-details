package details

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

// GetAll godoc
// @Summary Retrieves all product details
// @Produce json
// @Success 200 {array} Details
// @Router /details/all [get]
// @Security ApiKeyAuth
func GetAll(c *fiber.Ctx) error {
	ds, err := AllDetails()
	if err != nil {
		log.Error().Err(err).Msg("error while get all details from db")
	}
	if err := c.JSON(ds); err != nil {
		log.Error().Err(err).Msg("error while marshal all details")
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
func Get(c *fiber.Ctx) error {
	ids := c.Params("id", "")
	if !primitive.IsValidObjectID(ids) {
		return fiber.NewError(http.StatusBadRequest, "wrong id format "+ids)
	}
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		log.Error().Err(err).Msg("error while reading details id")
	}
	d, err := GetById(id)
	if err != nil {
		log.Error().Err(err).Msg("error while getting one details by id")
	}
	if err := c.JSON(d); err != nil {
		log.Error().Err(err).Msg("error while marshal one details")
	}
	return nil
}

// Create godoc
// @Summary Creates a new product-details from the received json document
// @Accept json
// @Param request body _withoutId true "product-details schema"
// @Success 200 {object} Details
// @Router /details/create [post]
// @Security ApiKeyAuth
func Create(c *fiber.Ctx) error {
	d := Details{}
	if err := c.BodyParser(&d); err != nil {
		log.Error().Err(err).Msg("error while decode details")
		return fiber.NewError(http.StatusBadRequest, "error while decode request body")
	}

	d, err := InsertOne(d)
	if err != nil {
		log.Error().Err(err).Msg("error while insert one details to db")
	}
	if err := c.JSON(d); err != nil {
		log.Error().Err(err).Msg("error while marshal inserted one details")
	}
	return nil
}
