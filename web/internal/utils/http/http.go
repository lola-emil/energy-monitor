package httputil

import (
	"net/http"
	"strconv"
)

func GetUserID(r *http.Request) int64 {
	return r.Context().Value("user_id").(int64)
}

func ParseIntPtr(s string) *int64 {
	if s == "" {
		return nil
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return &v
}
