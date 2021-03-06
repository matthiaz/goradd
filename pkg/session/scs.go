package session

import (
	"context"
	"github.com/alexedwards/scs"
	"github.com/goradd/goradd/pkg/log"
	"net/http"
)

const scsSessionDataKey = "goradd.data"

// ScsManager satisfies the ManagerI interface for the github.com/alexedwards/scs session manager. You can use it as an example
// of how to incorporate a different session manager into your app.
type ScsManager struct {
	*scs.Manager
}

func NewScsManager(mgr *scs.Manager) ManagerI {
	return ScsManager{mgr}
}

// Use is an http handler that wraps the session management process. It will get and put session data
// into the http context.
func (mgr ScsManager) Use(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var data []byte
		var temp string
		// get the session. All of our session data is stored in only one key in the session manager.
		session := mgr.Manager.Load(r)
		if err := session.Touch(w); err != nil { // Make sure to get a cookie in our header if we don't have one
			log.Errorf("Error loading session: %s", err.Error()) // we can't panic here, because our panic handlers have not been set up
		}
		data, _ = session.GetBytes(scsSessionDataKey)
		sessionData := NewSession()
		if data != nil {
			if err := sessionData.UnmarshalBinary(data); err != nil {
				log.Errorf("Error unpacking session data: %s", err.Error())
			}
		}

		if sessionData.Has(sessionResetKey) {
			// Our previous session requested a reset. We can't reset after writing, so we reset here at the start of the next request.
			sessionData.Delete(sessionResetKey)
			if err := session.RenewToken(w); err != nil {
				log.Errorf("Error renewing session token: %s", err.Error())
			}
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, sessionContext, sessionData)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)

		// write out the changed session. The below will attempt to write a cookie, but it can't because headers have already been written.
		// That is OK, because of our Touch above.
		if sessionData.Len() > 0 {
			var err error
			data, err = sessionData.MarshalBinary()
			if err != nil {
				// This is an application error. We put data into the session which is not serializable.
				s := "Error marshalling session data: %s" + err.Error()
				log.Error(s)
				http.Error(w, s, 500)
			}
			temp = string(data)
			_ = temp
			if err := session.PutBytes(w, scsSessionDataKey, data); err != nil {
				log.Errorf("Error putting session data: %s", err.Error())
			}
		} else {
			if err := session.Clear(w); err != nil {
				log.Errorf("Error clearing session: %s", err.Error())
			}
		}
	}
	return mgr.Manager.Use(http.HandlerFunc(fn))
}
