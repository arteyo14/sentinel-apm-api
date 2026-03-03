package cookie

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetAuthCookies(c *gin.Context, accessToken string, refreshToken string) {
	appEnv := os.Getenv("APP_ENV")
	domain := os.Getenv("COOKIE_DOMAIN")

	isProduction := appEnv == "production"

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"access_token",
		accessToken,
		900,
		"/",
		domain,
		isProduction,
		true,
	)

	c.SetCookie(
		"refresh_token",
		refreshToken,
		60*60*24*7,
		"/api/v1/auth/refresh-token",
		domain,
		isProduction,
		true,
	)
}

func ClearAuthCookies(c *gin.Context) {
	appEnv := os.Getenv("APP_ENV")
	domain := os.Getenv("COOKIE_DOMAIN")

	isProduction := appEnv == "production"

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"access_token",
		"",
		-1,
		"/",
		domain,
		isProduction,
		true,
	)

	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/api/v1/auth/refresh-token",
		domain,
		isProduction,
		true,
	)
}
