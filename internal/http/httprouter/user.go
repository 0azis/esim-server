package httprouter

import (
	"esim/internal/app"
	"esim/internal/http"
	"esim/internal/store"
	"esim/pkg"

	"github.com/labstack/echo/v4"
)

type userControllers interface {
	AuthEmail(c echo.Context) error
	CheckVerificationCode(c echo.Context) error
	AuthTelegram(c echo.Context) error
}

type user struct {
	app app.App
}

func (u user) AuthEmail(c echo.Context) error {
	emailCredentials := struct {
		Email string `json:"email"`
	}{}

	if err := c.Bind(&emailCredentials); err != nil {
		return http.NewErrors(c, 400,
			http.Error(400, "Плохие данные", "Неправильные поля для регистрации или входа пользователя. Попробуйте еще раз"),
		)
	}

	code := pkg.GenerateVerificationCode()
	err := u.app.Redis.SetCode(emailCredentials.Email, code)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Одноразовый код", "Ошибка при добавлении одноразового кода в хранилище"),
		)
	}

	err = u.app.Mailer.SendVerificationCode(emailCredentials.Email, code)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Отправка письма", "Ошибка при отправлении письма с одноразовым кодом"),
		)
	}

	return http.NewResponseWithDescription(c, 200, "Письмо с кодом было отправлено")
}

func (u user) AuthTelegram(c echo.Context) error {
	initDataCredentials := struct {
		InitData string `json:"initData"`
	}{}

	if err := c.Bind(&initDataCredentials); err != nil {
		return http.NewErrors(c, 400,
			http.Error(400, "Плохие данные", "Неправильные поля для регистрации или входа пользователя. Попробуйте еще раз"),
		)
	}

	status := u.app.Bot.Utils().ValidateInitData(initDataCredentials.InitData)
	if !status {
		return http.NewErrors(c, 401,
			http.Error(401, "Плохие данные", "Мы не смогли удостовериться в подлинности данных для входа. Попробуйте позже."),
		)
	}

	userID, err := u.app.Bot.Utils().ParseInitData(initDataCredentials.InitData)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Проблема с расшифровкой", "Мы не смогли расшифровать ваши даных для входа. Попробуйте позже."),
		)
	}

	token, err := u.app.Jwt.SignJwt(userID)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Проблема с выдачей токена", "Ошибка при генерации токена входа для вас. Попробуйте позже"),
		)
	}

	return http.NewResponseWithDescription(c, 200, token)
}

func (u user) CheckVerificationCode(c echo.Context) error {
	verificationPayload := struct {
		Email string `json:"email"`
		Code  int    `json:"code"`
	}{}
	if err := c.Bind(&verificationPayload); err != nil {
		return http.NewErrors(c, 400,
			http.Error(400, "Плохие данные", "Содержимое содержит некорректные данные"),
		)
	}

	// Redis Checking
	code, err := u.app.Redis.GetCode(verificationPayload.Email)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Проблема с проверкой кода", "Ошибка при проверке вашего кода для входа. Попробуйте позже."),
		)
	}

	if verificationPayload.Code != code {
		return http.NewErrors(c, 400,
			http.Error(401, "Код не совпадает", "Вы ввели неправильный код. Попробуйте запросить новый или сверьте старый"),
		)
	}

	user := store.NewEmailUser(verificationPayload.Email)
	id, err := u.app.Store.User().Create(user)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Пользователь не создан", "Ошибка при создании пользователя. Попробуйте позже"),
		)
	}

	token, err := u.app.Jwt.SignJwt(id)
	if err != nil {
		u.app.Logger.Error(err.Error())
		return http.NewErrors(c, 500,
			http.Error(500, "Токен не создан", "Ошибка при создании токена авторизации. Попробуйте позже"),
		)
	}

	return http.NewResponseWithDescription(c, 200, token)
}

func NewUserControllers(app app.App) userControllers {
	return user{app}
}
