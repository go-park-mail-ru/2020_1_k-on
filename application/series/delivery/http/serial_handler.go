package http

import (
	"errors"
	"github.com/go-park-mail-ru/2020_1_k-on/application/microservices/series/client"
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
	person "github.com/go-park-mail-ru/2020_1_k-on/application/person"
	"github.com/go-park-mail-ru/2020_1_k-on/application/series"
	"github.com/labstack/echo"
	"strconv"
)

type SeriesHandler struct {
	rpcSeriesFilter *client.SeriesFilterClient
	usecase         series.Usecase
	pusecase        person.UseCase
}

func NewSeriesHandler(e *echo.Echo,
	rpcSeriesFilter *client.SeriesFilterClient,
	usecase series.Usecase, pusecase person.UseCase) {
	handler := &SeriesHandler{
		rpcSeriesFilter: rpcSeriesFilter,
		usecase:         usecase,
		pusecase:        pusecase,
	}
	e.GET("/series/:id", handler.GetSeries)
	e.GET("/series/:id/seasons", handler.GetSeriesSeasons)
	e.GET("/seasons/:id/episodes", handler.GetSeasonEpisodes)

	e.GET("/series", handler.FilterSeriesList)
	e.GET("/series/filter", handler.FilterSeriesData)
}

func (sh SeriesHandler) FilterSeriesData(ctx echo.Context) error {
	d, ok := sh.rpcSeriesFilter.GetFilterFields()
	if !ok {
		resp, _ := models.Generate(500, "can't get data", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return errors.New("can't get data")
	}
	resp, _ := models.Generate(200, d, &ctx).MarshalJSON()
	_, err := ctx.Response().Write(resp)
	return err
}

func (sh SeriesHandler) FilterSeriesList(ctx echo.Context) error {
	query := ctx.QueryParams()
	s, ok := sh.rpcSeriesFilter.GetFilteredSeries(query)
	if !ok {
		resp, _ := models.Generate(500, "can't get series", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return errors.New("can't get series")
	}
	var sl models.ListSeriesArr
	resp, _ := models.Generate(200, sl.Convert(s), &ctx).MarshalJSON()
	_, err := ctx.Response().Write(resp)
	return err
}

func (sh SeriesHandler) GetSeries(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp, _ := models.Generate(400, "not number", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return err
	}
	serial, ok := sh.usecase.GetSeriesByID(uint(id))
	if !ok {
		resp, _ := models.Generate(404, "Not Found", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return errors.New("Not Found")
	}
	a, err := sh.pusecase.GetActorsForSeries(serial.ID)
	g, _ := sh.usecase.GetSeriesGenres(serial.ID)
	r := make(map[string]interface{})
	r["object"] = serial
	r["actors"] = a
	r["genres"] = g
	resp, _ := models.Generate(200, r, &ctx).MarshalJSON()
	_, err = ctx.Response().Write(resp)
	return err
}

func (sh SeriesHandler) GetSeriesSeasons(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp, _ := models.Generate(400, "not number", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return err
	}
	seasons, ok := sh.usecase.GetSeriesSeasons(uint(id))
	if !ok {
		resp, _ := models.Generate(404, "Not Found", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return errors.New("Not Found")
	}
	resp, _ := models.Generate(200, seasons, &ctx).MarshalJSON()
	_, err = ctx.Response().Write(resp)
	return err
}

func (sh SeriesHandler) GetSeasonEpisodes(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		resp, _ := models.Generate(400, "not number", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return err
	}
	seasons, ok := sh.usecase.GetSeasonEpisodes(uint(id))
	if !ok {
		resp, _ := models.Generate(404, "Not Found", &ctx).MarshalJSON()
		ctx.Response().Write(resp)
		return errors.New("Not Found")
	}
	resp, _ := models.Generate(200, seasons, &ctx).MarshalJSON()
	_, err = ctx.Response().Write(resp)
	return err
}
