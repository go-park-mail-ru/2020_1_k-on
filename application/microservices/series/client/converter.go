package client

import (
	api "github.com/go-park-mail-ru/2020_1_k-on/application/microservices/series/api"
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
)

func convertToFilter(fields map[string][]string) *api.Filter {
	result := new(api.Filter)
	result.Fields = make(map[string]*api.Strings)
	for key, value := range fields {
		result.Fields[key] = new(api.Strings)
		result.Fields[key].Value = value
	}
	return result
}

func convertFromSeriesList(slist *api.SeriesList) models.SeriesArr {
	result := models.SeriesArr{}
	for _, value := range slist.List {
		result = append(result, models.Series{
			ID:              uint(value.Id),
			MainGenre:       value.MainGenre,
			RussianName:     value.RussianName,
			EnglishName:     value.EnglishName,
			TrailerLink:     value.TrailerLink,
			Rating:          value.Rating,
			ImdbRating:      value.ImdbRating,
			TotalVotes:      int(value.TotalVotes),
			SumVotes:        int(value.SumVotes),
			Description:     value.Description,
			Image:           value.Image,
			BackgroundImage: value.BackgroundImage,
			Country:         value.Country,
			YearFirst:       int(value.YearFirst),
			YearLast:        int(value.YearLast),
			AgeLimit:        int(value.AgeLimit),
		})
	}
	return result
}

func convertFromFilterFields(fields *api.FilterFields) map[string]models.Genres {
	result := make(map[string]models.Genres)
	for key, value := range fields.Fields {
		genres := models.Genres{}
		for _, g := range value.List {
			genres = append(genres, models.Genre{
				Name:      g.Name,
				Reference: g.Reference,
			})
		}
		result[key] = genres
	}
	return result
}
