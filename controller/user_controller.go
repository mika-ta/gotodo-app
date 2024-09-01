package controller

import "github.com/labstack/echo/v4"

type IUserController interface {
	SignUp(c echo.Contect) error // echoで定義されてるContext型
	LogIn(c echo.Contect) error
	LogOut(c echo.Contect) error
}

type UserController struct {
	uu usecase.IUserUsecase	
}

func NewUserController(uu usecaseIUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Contect) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cokkie := new(http.Cokkie)
	cokkie.Name = "token"
	cokkie.Value = tokenString
	cokkie.Expires = time.Now().Add(24 * time.Hour)
	cokkie.Path = "/"
	cokkie.Domain = os.Getenv("API_DOMAIN")
	// cookie.secure = true
	cokkie.HttpOnly = true
	cokkie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) LogOut(c echo.Context) error {
	cookie *= new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}