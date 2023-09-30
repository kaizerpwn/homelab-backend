package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaizerpwn/homelab-backend/initializers"
	"github.com/kaizerpwn/homelab-backend/models"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /api

// @Summary Get all users from the database
// @Description Fetch all users from the database (administrator permission needed)
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users [get]
func GetAllUsers(c *gin.Context) {

	// >> Get all users from database
	var users []models.User
	initializers.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, users)
}

// @Summary Get a User by ID
// @Description Retrieve a user's information by their unique ID.
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID" Format(int64)
// @Success 200 {object} models.User
// @Failure 404 {string} string "User with that ID doesn't exist"
// @Router /api/users/{id} [get]
func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := initializers.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with that ID doesn't exist",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

// @Summary User Login
// @Description Authenticates a user based on provided email and password.
// @Tags users
// @Accept json
// @Produce json
// @Param email body string true "User's email address" Format(email)
// @Param password body string true "User's password"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Invalid Password"
// @Failure 404 {string} string "User with that credentials doesn't exist."
// @Router /api/users/login [post]
func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// >> Parse JSON request body and bind it to the struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	// >> Find user in the database
	var user models.User
	result := initializers.DB.First(&user, "email = ?", body.Email)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User with that credentials doesn't exist.",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"user": gin.H{
				"image":   user.Image,
				"name":    user.Name,
				"surname": user.Surname,
				"email":   user.Email,
			},
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Password is not valid.",
		})
	}
}

// @Summary User Registration
// @Description Creates a new user account with the provided information.
// @Tags users
// @Accept json
// @Produce json
// @Param name body string true "User's first name"
// @Param surname body string true "User's last name"
// @Param email body string true "User's email address" Format(email)
// @Param password body string true "User's password"
// @Param city body string true "User's city"
// @Success 200 {string} string "Account successfully registered."
// @Failure 400 {string} string "Invalid request data"
// @Failure 400 {string} string "All fields are required"
// @Failure 409 {string} string "User already exists."
// @Failure 500 {string} string "Internal server error."
// @Router /api/users/register [post]
func Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Surname  string `json:"surname" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		City     string `json:"city" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request data",
		})
		return
	}

	// >> Check if any of the fields are empty
	if body.Name == "" || body.Surname == "" || body.Email == "" || body.Password == "" || body.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "All fields are required",
		})
		return
	}

	// >> Check if user already exists in database
	var existingUser models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists.",
		})
		return
	}

	// >> Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	// >> Insert new user into database
	user := models.User{Name: body.Name, Surname: body.Surname, Email: body.Email, City: body.City, Password: string(hashedPassword)}
	result = initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully registered.",
	})
}
