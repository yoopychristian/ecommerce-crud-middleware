package middleware

import (
	"ecommerce-crud-middleware/utils"
	"log"
	"net/http"

	ua "github.com/mileusna/useragent"
)

func ActivityLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.UserAgent()
		detect := ua.Parse(userAgent)
		userDevice := utils.DeviceDetect(detect)
		methodDetect := utils.MethodDetect(r)
		log.Printf("Accessing path %v using %v with : %v\n", r.RequestURI, methodDetect, userDevice)
		next.ServeHTTP(w, r)
	})
}
