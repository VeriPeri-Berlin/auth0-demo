package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Handler for our logged-in user page.
func Handler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	u := struct {
		Profile string
		IDToken string
	}{
		Profile: profile.(string),
		IDToken: session.Get("id_token").(string),
	}

	ctx.HTML(http.StatusOK, "user.html", u)
}
