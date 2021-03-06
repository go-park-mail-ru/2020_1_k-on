package http

import (
	"bytes"
	mockfilm "github.com/go-park-mail-ru/2020_1_k-on/application/film/mocks"
	"github.com/go-park-mail-ru/2020_1_k-on/application/film/usecase"
	"github.com/go-park-mail-ru/2020_1_k-on/application/microservices/film/client"
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
	mock_p "github.com/go-park-mail-ru/2020_1_k-on/application/person/mocks"
	usecase2 "github.com/go-park-mail-ru/2020_1_k-on/application/person/usecase"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/microcosm-cc/bluemonday"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

var image = "image"

var mg = "mg"
var rn = "rn"
var en = "en"
var sumvotes = 0
var totalvotes = 0
var tl = "tl"
var rating = 1.2
var imdbrating = 9.87
var d = "d"
var c = "c"
var year = 2012
var agelimit = 10
var fid = uint(1)

var testFilm = models.Film{
	ID:              fid,
	MainGenre:       mg,
	RussianName:     rn,
	EnglishName:     en,
	TrailerLink:     tl,
	Rating:          rating,
	ImdbRating:      imdbrating,
	Description:     d,
	Image:           image,
	Country:         c,
	Year:            year,
	AgeLimit:        agelimit,
	SumVotes:        sumvotes,
	TotalVotes:      totalvotes,
	BackgroundImage: image,
}

func setupEcho(t *testing.T, url, method string) (echo.Context,
	FilmHandler, *mockfilm.MockRepository, *mock_p.MockRepository, *client.MockIFilmFilterClient) {
	e := echo.New()
	r := e.Router()
	r.Add(method, url, func(echo.Context) error { return nil })
	var req *http.Request
	switch method {
	case http.MethodPost:
		f, _ := testFilm.MarshalJSON()
		req = httptest.NewRequest(http.MethodGet, url, bytes.NewBuffer(f))
	case http.MethodGet:
		req = httptest.NewRequest(http.MethodGet, url, nil)
	}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath(url)
	ctrl := gomock.NewController(t)
	films := mockfilm.NewMockRepository(ctrl)
	usecase := usecase.NewFilmUsecase(films)
	person := mock_p.NewMockRepository(ctrl)
	pusecase := usecase2.NewPerson(person, nil)
	rpc := client.NewMockIFilmFilterClient(ctrl)
	fh := FilmHandler{rpcFilmFilter: rpc, fusecase: usecase, pusecase: pusecase, sanitizer: bluemonday.UGCPolicy()}
	return c, fh, films, person, rpc

}

func TestFilmHandler_GetFilm(t *testing.T) {
	c, fh, films, person, _ := setupEcho(t, "/films/:id", http.MethodGet)
	c.SetParamNames("id")
	c.SetParamValues("1")
	films.EXPECT().GetById(gomock.Eq(testFilm.ID)).Return(&testFilm, true)
	films.EXPECT().GetFilmGenres(gomock.Eq(testFilm.ID)).Return(nil, true)
	films.EXPECT().GetSimilarFilms(gomock.Any()).Return(nil, true)
	films.EXPECT().GetSimilarSeries(gomock.Any()).Return(nil, true)
	person.EXPECT().GetActorsForFilm(testFilm.ID).Return(nil, nil)
	err := fh.GetFilm(c)
	require.Equal(t, err, nil)
}

func TestFilmHandler_GetFilmList(t *testing.T) {
	c, fh, films, _, _ := setupEcho(t, "/", http.MethodGet)
	films.EXPECT().GetFilmsArr(uint(13), uint(0)).Return(&models.Films{testFilm, testFilm, testFilm, testFilm, testFilm}, true)
	err := fh.GetFilmList(c)
	require.Equal(t, err, nil)
}

func TestFilmHandler_GetFilmList2(t *testing.T) {
	c, fh, films, _, _ := setupEcho(t, "/", http.MethodGet)
	films.EXPECT().GetFilmsArr(uint(13), uint(0)).Return(&models.Films{}, false)
	err := fh.GetFilmList(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_CreateFilm(t *testing.T) {
	c, fh, films, _, _ := setupEcho(t, "/films", http.MethodPost)
	films.EXPECT().Create(&testFilm).Return(testFilm, true)
	err := fh.CreateFilm(c)
	require.Equal(t, err, nil)
}

func TestFilmHandler_CreateFilm2(t *testing.T) {
	c, fh, films, _, _ := setupEcho(t, "/films", http.MethodGet)
	films.EXPECT().Create(&testFilm).Return(testFilm, true)
	err := fh.CreateFilm(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_CreateFilm3(t *testing.T) {
	c, fh, films, _, _ := setupEcho(t, "/films", http.MethodPost)
	films.EXPECT().Create(&testFilm).Return(models.Film{}, false)
	err := fh.CreateFilm(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_GetFilm2(t *testing.T) {
	c, fh, films, person, _ := setupEcho(t, "/films/:id", http.MethodGet)
	films.EXPECT().GetById(gomock.Eq(testFilm.ID)).Return(&testFilm, true)
	person.EXPECT().GetActorsForFilm(testFilm.ID).Return(nil, nil)
	err := fh.GetFilm(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_GetFilm3(t *testing.T) {
	c, fh, films, person, _ := setupEcho(t, "/films/:id", http.MethodGet)
	c.SetParamNames("id")
	c.SetParamValues("1")
	films.EXPECT().GetById(testFilm.ID).Return(&models.Film{}, false)
	person.EXPECT().GetActorsForFilm(testFilm.ID).Return(nil, nil)
	err := fh.GetFilm(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_FilterFilmData(t *testing.T) {
	c, fh, _, _, rpc := setupEcho(t, "/films/filter", http.MethodGet)
	rpc.EXPECT().GetFilterFields().Return(nil, true)
	err := fh.FilterFilmData(c)
	require.Equal(t, err, nil)
}

func TestFilmHandler_FilterFilmData2(t *testing.T) {
	c, fh, _, _, rpc := setupEcho(t, "/films/filter", http.MethodGet)
	rpc.EXPECT().GetFilterFields().Return(nil, false)
	err := fh.FilterFilmData(c)
	require.NotEqual(t, err, nil)
}

func TestFilmHandler_FilterFilmList(t *testing.T) {
	q := make(map[string][]string)
	q["year"] = []string{"year"}
	c, fh, _, _, rpc := setupEcho(t, "/films", http.MethodGet)
	c.QueryParams().Add("year", "year")
	rpc.EXPECT().GetFilteredFilms(q).Return(models.Films{}, true)
	err := fh.FilterFilmList(c)
	require.Equal(t, err, nil)
}

func TestFilmHandler_FilterFilmList2(t *testing.T) {
	q := make(map[string][]string)
	q["year"] = []string{"year"}
	c, fh, _, _, rpc := setupEcho(t, "/films", http.MethodGet)
	c.QueryParams().Add("year", "year")
	rpc.EXPECT().GetFilteredFilms(q).Return(models.Films{}, false)
	err := fh.FilterFilmList(c)
	require.NotEqual(t, err, nil)
}
