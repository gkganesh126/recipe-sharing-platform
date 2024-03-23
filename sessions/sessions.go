package sessions

import (
	"errors"
	"net/http"
	"time"

	db "github.com/gkganesh126/recipe-sharing-platform/db-ops"
	"github.com/gorilla/securecookie"
)

var cookieHandler *securecookie.SecureCookie

func InitCookieHandler() {
	cookieHandler = securecookie.New([]byte("12345678901234567890123456789012"), []byte("01234567890123456789012345678901"))
	cookieHandler.MaxAge(0)
}

func GetSessionUserName(request *http.Request) (string, string, error) {
	var cookie *http.Cookie
	var userName string
	var err error
	if cookie, err = request.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		} else {
			return "", "", errors.New("error decoding userID: " + err.Error())
		}
	} else {
		return "", "", errors.New("error parsing cookie session: " + err.Error())
	}
	return cookie.Value, userName, nil
}

func IsValidSession(request *http.Request) (string, error) {
	requestSession, userName, err := GetSessionUserName(request)
	if err != nil {
		return "", err
	}
	context := NewContext()
	defer context.Close()
	c := context.RecipeCollectionPlatformDbCollection("users")
	repo := &db.UserRepository{C: c}
	// Get all users form repository
	dbSession, err := repo.GetSession(userName)
	if err != nil {
		return "", err
	}
	//fmt.Println("dbSession: ", dbSession, "requestSession: ", requestSession)
	if dbSession != requestSession {
		return "", errors.New("session mismatch between db and request")
	}
	return userName, nil
}
func SetSession(userName string, response *http.ResponseWriter) error {
	value := map[string]string{
		"name": userName,
	}
	var cookie *http.Cookie
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie = &http.Cookie{
			Name:    "session",
			Value:   encoded,
			Expires: time.Now().Add(50 * 365 * 24 * time.Hour),
			Secure:  false,
			//Path:    "/",
			//expires can be given
		}
	} else {
		return err
	}

	context := NewContext()
	defer context.Close()
	c := context.RecipeCollectionPlatformDbCollection("users")
	repo := &db.UserRepository{C: c}
	if v := cookie.Value; v != "" {
		userID, err := repo.GetUserID(userName)
		if err != nil {
			return err
		}
		err = repo.UpdateSession(userID.UserID.Hex(), v)
		if err != nil {
			return err
		}
	} else {
		return errors.New("cookie is empty")
	}

	http.SetCookie(*response, cookie)
	return nil
}

func ClearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}
