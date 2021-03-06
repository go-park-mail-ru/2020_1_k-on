package http

import (
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
	"github.com/go-park-mail-ru/2020_1_k-on/application/playlist"
	"github.com/go-park-mail-ru/2020_1_k-on/application/server/middleware"
	"github.com/go-park-mail-ru/2020_1_k-on/pkg/constants"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type PlaylistHandler struct {
	useCase   playlist.UseCase
	logger    *zap.Logger
	sanitizer *bluemonday.Policy
}

func NewPlaylistHandler(e *echo.Echo,
	uc playlist.UseCase, auth middleware.Auth, logger *zap.Logger, sanitizer *bluemonday.Policy) {
	handler := PlaylistHandler{
		useCase:   uc,
		logger:    logger,
		sanitizer: sanitizer,
	}

	e.POST("/playlist", handler.Create, auth.GetSession, middleware.CSRF, middleware.ParseErrors)
	e.POST("/playlist/:pid/films/:id", handler.AddFilm, auth.GetSession, middleware.CSRF, middleware.ParseErrors)
	e.POST("/playlist/:pid/series/:id", handler.AddSeries, auth.GetSession, middleware.CSRF, middleware.ParseErrors)

	e.GET("/playlist/:pid", handler.Get, auth.GetSession, middleware.ParseErrors)
	e.GET("/playlist", handler.GetUserPlaylists, auth.GetSession, middleware.ParseErrors)
	e.GET("/playlist/user/:id", handler.GetUserPublicPlaylists, middleware.ParseErrors)
	e.GET("/films/:id/playlists", handler.GetPlaylistsWithoutFilm, auth.GetSession, middleware.ParseErrors)
	e.GET("/series/:id/playlists", handler.GetPlaylistsWithoutSer, auth.GetSession, middleware.ParseErrors)
	e.GET("/index", handler.GetAdminPlaylists, middleware.ParseErrors)

	e.DELETE("/playlist/:pid", handler.Delete, auth.GetSession, middleware.CSRF, middleware.ParseErrors)
	e.DELETE("/playlist/:pid/film/:id", handler.DeleteFilm, auth.GetSession, middleware.CSRF, middleware.ParseErrors)
	e.DELETE("/playlist/:pid/series/:id", handler.DeleteSeries, auth.GetSession, middleware.CSRF, middleware.ParseErrors)
}

func (handler *PlaylistHandler) Create(ctx echo.Context) error {
	play, err := handler.parseRequestBody(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "request parser error")
	}

	play, err = handler.useCase.Create(play)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, play)
}

func (handler *PlaylistHandler) AddFilm(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	filmId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	err = handler.useCase.AddFilm(pid, filmId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, "")
}

func (handler *PlaylistHandler) AddSeries(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	seriesId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	err = handler.useCase.AddSeries(pid, seriesId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, "")
}

func (handler *PlaylistHandler) Get(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	play, err := handler.useCase.Get(pid, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, play)
}

func (handler *PlaylistHandler) GetUserPlaylists(ctx echo.Context) error {
	userId := ctx.Get(constants.UserIdKey).(uint)

	plist, err := handler.useCase.GetUserPlaylists(userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, plist)
}

func (handler *PlaylistHandler) GetUserPublicPlaylists(ctx echo.Context) error {
	userId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	plist, err := handler.useCase.GetUserPublicPlaylists(userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, plist)
}

func (handler *PlaylistHandler) GetPlaylistsWithoutSer(ctx echo.Context) error {
	userId := ctx.Get(constants.UserIdKey).(uint)
	serId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	plist, err := handler.useCase.GetPlaylistsWithoutSer(serId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, plist)
}

func (handler *PlaylistHandler) GetPlaylistsWithoutFilm(ctx echo.Context) error {
	userId := ctx.Get(constants.UserIdKey).(uint)
	filmId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	plist, err := handler.useCase.GetPlaylistsWithoutFilm(filmId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, plist)
}

func (handler *PlaylistHandler) GetAdminPlaylists(ctx echo.Context) error {
	plist, err := handler.useCase.GetAdminPlaylists()
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, plist)
}

func (handler *PlaylistHandler) Delete(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	err = handler.useCase.Delete(pid, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, "")
}

func (handler *PlaylistHandler) DeleteFilm(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	filmId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	err = handler.useCase.DeleteFilm(pid, filmId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, "")
}

func (handler *PlaylistHandler) DeleteSeries(ctx echo.Context) error {
	pid, err := handler.getParamId(ctx, "pid")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	seriesId, err := handler.getParamId(ctx, "id")
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	err = handler.useCase.DeleteSeries(pid, seriesId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, "")
}

func (handler *PlaylistHandler) parseRequestBody(ctx echo.Context) (*models.Playlist, error) {
	play := new(models.Playlist)
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, play); err != nil {
		handler.logger.Error("request parser error")
		return nil, err
	}

	play.Name = handler.sanitizer.Sanitize(play.Name)
	play.UserId = ctx.Get(constants.UserIdKey).(uint)

	return play, nil
}

func (handler *PlaylistHandler) getParamId(ctx echo.Context, name string) (uint, error) {
	id, err := strconv.Atoi(ctx.Param(name))
	if err != nil || id < 0 {
		handler.logger.Error("wrong parameter")
		return 0, err
	}

	return uint(id), nil
}
