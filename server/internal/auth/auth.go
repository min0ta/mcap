package auth

import (
	"fmt"
	"io"
	"mcap/internal/config"
	"mcap/internal/errors"
	"mcap/internal/log"
	"mcap/internal/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const cookieName = "auth-jwt"
const (
	RoleGuest = iota
	RoleModerator
	RoleAdmin
)

type Authoriztaion struct {
	db  *JsonDB
	cfg *config.Config
	// logger *log.Logger
}

type Role int

func New(cfg *config.Config, logger *log.Logger) *Authoriztaion {
	a := &Authoriztaion{
		cfg: cfg,
		db:  newJsonDb(),
		// logger: ,
	}
	a.db.Connect(cfg.PATH_TO_JSON_DB)
	return a
}

func (s *Authoriztaion) Authorize(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	query := AuthoriztaionQuery{}
	err := utils.ReadJson(r, &query)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	role := s.db.contains(func(u usersRecord) bool {
		return (u.Username == query.User) && (u.Password == query.Password)
	})

	err = s.setJwtToken(role, w)
	if err != nil {
		errors.HttpError(w, errors.ErrorBadLoginOrPassword, 400)
		return
	}
	if role != RoleGuest {
		utils.WriteResult(w, utils.Response{"succes": true}, 200)
		return
	}
	errors.HttpError(w, errors.ErrorBadLoginOrPassword, 401)
}

func (s *Authoriztaion) TestIfAuth(w http.ResponseWriter, r *http.Request) {
	role := s.AuthCheck(r)

	if role == RoleAdmin {
		io.WriteString(w, "Admin")
		return
	}
	if role == RoleModerator {
		io.WriteString(w, "Moder")
		return
	}
	io.WriteString(w, "Guest")
}

func (s *Authoriztaion) Unauthorize(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     cookieName,
		Expires:  time.Now().Add(time.Second * -1),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, &cookie)
}

func (s *Authoriztaion) AuthCheck(r *http.Request) Role {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return RoleGuest
	}
	unparsedToken := cookie.Value

	token, err := jwt.Parse(unparsedToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT_SIGNING_KEY), nil
	})
	if err != nil {
		// s.logger.WriteFormat("ACCESS DENY! REQUEST FROM IP: %s", r.RemoteAddr)
		fmt.Println("access deny")
		return RoleGuest
	}
	roleStr, err := token.Claims.GetSubject()
	if err != nil {
		// s.logger.WriteFormat("ACCESS DENY! REQUEST FROM IP: %s", r.RemoteAddr)
		fmt.Println("access deny")

		return RoleGuest
	}
	role, _ := strconv.ParseInt(roleStr, 10, 32)
	if role == RoleModerator {
		return RoleModerator
	}
	if role == RoleAdmin {
		return RoleAdmin
	}

	// s.logger.WriteFormat("ACCESS DENY! REQUEST FROM IP: %s", r.RemoteAddr)
	fmt.Println("access deny")

	return RoleGuest
}

func (s *Authoriztaion) setJwtToken(role Role, w http.ResponseWriter) error {
	signKey := []byte(s.cfg.JWT_SIGNING_KEY)
	expires := time.Now().Add(time.Duration(s.cfg.JWT_EXPIRES) * time.Second)
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expires),
		Subject:   fmt.Sprint(role),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(signKey)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    signed,
		Expires:  expires,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}

	http.SetCookie(w, cookie)

	return nil
}
