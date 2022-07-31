package common

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

type PaginateData struct {
	Limit          int
	Offset         int
	OrderColumn    string
	OrderDirection string
}

func ExtractPaginationData(c *gin.Context, defaultOffset, defaultTake int) PaginateData {
	take, err := strconv.Atoi(c.DefaultQuery("limit", strconv.Itoa(defaultTake)))
	if err != nil {
		take = defaultTake
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", strconv.Itoa(defaultOffset)))
	if err != nil {
		offset = defaultOffset
	}

	orderCol := c.DefaultQuery("orderCol", "")
	orderDir := c.DefaultQuery("orderDir", "")

	return PaginateData{
		Limit:          take,
		Offset:         offset,
		OrderColumn:    orderCol,
		OrderDirection: orderDir,
	}
}

func GenerateRandomString(strLen uint8, allowedCharSet string) string {
	var charSet []rune

	if len(allowedCharSet) > 0 {
		charSet = []rune(allowedCharSet)
	} else {
		charSet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	}

	b := make([]rune, strLen)
	for i := range b {
		b[i] = charSet[rand.Intn(len(charSet))]
	}

	return string(b)
}
