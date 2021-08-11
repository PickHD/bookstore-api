package utils

import (
	"math"
	"strconv"
)

type Pagination struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func GeneratePaginator(getQLimit, getQPage string) (Pagination, error) {
	limit, err := strconv.Atoi(getQLimit)
	if err != nil {
		return Pagination{}, err
	}

	page, err := strconv.Atoi(getQPage)
	if err != nil {
		return Pagination{}, err
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}, nil
}

func CountTotalPage(totalData int, limit *Pagination) int {
	if totalData == 0 {
		return 1
	}
	return int(math.Ceil(float64(totalData) / float64(limit.Limit)))
}
