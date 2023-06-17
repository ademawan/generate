package middleware

import (
	"api_simple/helper"
	exception "api_simple/helper/exceptions"
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	redisV9 "github.com/redis/go-redis/v9"
)

type CustomMiddleware struct {
	redis *redisV9.Client
}

func NewCustomMiddleware(redisClient *redisV9.Client) *CustomMiddleware {
	return &CustomMiddleware{redis: redisClient}
}

func (m *CustomMiddleware) RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		var key string
		var max_pertime int
		var limit int
		now := time.Now().UnixNano()
		ip := r.Header.Get("X-Forwarded-For")
		fmt.Println("ddd:", ip)
		key = "request"
		max_pertime, _ = strconv.Atoi(os.Getenv("MAX_RATE_PERTIME_REQUEST"))
		limit, _ = strconv.Atoi(os.Getenv("MAX_RATE_REQUEST"))
		userCntKey := fmt.Sprint(ip, ":", key)

		limitTime := 1 * max_pertime

		slidingWindow := time.Duration(limitTime) * time.Second

		m.redis.ZRemRangeByScore(ctx, userCntKey, "0", fmt.Sprint(now-(slidingWindow.Nanoseconds()))).Result()

		reqs, _ := m.redis.ZRange(ctx, userCntKey, 0, -1).Result()
		fmt.Println(reqs)

		if len(reqs) >= limit {
			helper.ResponseWithErr(w, http.StatusTooManyRequests, exception.ErrTomanyRequest.Error(), nil)
			return
		}

		next.ServeHTTP(w, r)

		m.redis.ZAddNX(ctx, userCntKey, redisV9.Z{Score: float64(now), Member: float64(now)})
		m.redis.Expire(ctx, userCntKey, slidingWindow)
	})
}

// func (redis *CustomMiddleware) CaptcaValidation(next http.Handler) http.Handler {

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		response := r.Header.Get("captcha-id")
// 		isSuccess, _ := captcha.GoogleRecaptcha(response)

// 		if r.Header.Get("CHANNELID") == "TESTING" {
// 			next.ServeHTTP(w, r)
// 			return
// 		}
// 		if !isSuccess {
// 			helper.ResponseWithErr(w, http.StatusUnauthorized, "invalid signature", nil)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
