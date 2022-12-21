package config

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func SetupAuth() {
	key := "gilangcy-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30  // 30 days
	isProd := true        // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	if isProd {
		goth.UseProviders(
			google.New("297925112549-flbk68gc74acqr3untoluhikiqtrenst.apps.googleusercontent.com", "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", "https://my-open-source-project-372305.uc.r.appspot.com/auth/google/callback", "email", "profile"),
		)
	} else {
		goth.UseProviders(
			google.New("297925112549-flbk68gc74acqr3untoluhikiqtrenst.apps.googleusercontent.com", "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", "http://localhost:8080/auth/google/callback", "email", "profile"),
		)
	}

}
