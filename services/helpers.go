package services

import (
	"net/url"
	"strings"

	"github.com/cheqd/did-resolver/types"
	"github.com/labstack/echo/v4"
)

func GetContentType(accept string) types.ContentType {
	typeList := strings.Split(accept, ",")
	for _, cType := range typeList {
		result := types.ContentType(strings.Split(cType, ";")[0])
		if result == "*/*" || result == types.JSONLD {
			result = types.DIDJSONLD
		}
	}

	return ""
}

func PrepareQueries(c echo.Context) (rawQuery string, flag *string) {
	rawQuery = c.Request().URL.RawQuery
	flagIndex := strings.LastIndex(rawQuery, "%23")
	if flagIndex == -1 || strings.Contains(rawQuery[flagIndex:], "&") {
		return rawQuery, nil
	}
	queryFlag := rawQuery[flagIndex:]

	return rawQuery[0:flagIndex], &queryFlag
}

func GetDidParam(c echo.Context) (string, error) {
	return url.QueryUnescape(c.Param("did"))
}
