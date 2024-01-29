package middleware

import (
	"net/http"
	"strings"

	"gitlab.com/cdv-projects/go-apis/internal/common/entities"
	"gitlab.com/cdv-projects/go-apis/internal/common/models"
	"gitlab.com/cdv-projects/go-apis/internal/common/telemetry"
	"gorm.io/gorm"
)

// serviceJWTReqLogsName : service name
const serviceJWTReqLogsName string = "jwt_req_logs_middleware"

// JWTReqLogs : JWTReqLogs
type JWTReqLogs struct {
	db *gorm.DB
}

// NewJWTReqLogs creates a new JWTReqLogs handler with the provided options.
func NewJWTReqLogs(db *gorm.DB) *JWTReqLogs {
	return &JWTReqLogs{
		db: db,
	}
}

// Handler is a middleware that logs in database jwt validated http request
func (fc *JWTReqLogs) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ignore metrics calls
		if strings.Contains(r.RequestURI, "/metrics") {
			next.ServeHTTP(w, r)
			return
		}

		ctxJWT := r.Context().Value(CtxJWT).(entities.AuthJWT)

		// Log Validated JWT Request in nowaiting goroutine.
		go logJWTRequest(fc.db, r, &ctxJWT)

		next.ServeHTTP(w, r)
	})
}

// checkToken returns error if token is unkown or disabled.
func logJWTRequest(db *gorm.DB, r *http.Request, ctxJWT *entities.AuthJWT) {
	reqJWTLog := &models.ReqJWTLog{
		Method:         r.Method,
		CompanyCode:    ctxJWT.CompanyCode,
		SubscriptionId: ctxJWT.SubscriptionID,
		ApikeyId:       ctxJWT.ApikeyID,
	}

	if err := db.Create(&reqJWTLog).Error; err != nil {
		telemetry.LogError(serviceJWTReqLogsName, "Error when logging JWT validated request.", err)
	}
}
