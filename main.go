package main

import (
	"net/http"

	_ "iwa-work5/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Contact struct {
	ID    string `json:"id" example:"1"`
	Name  string `json:"name" example:"Иван Иванов"`
	Phone string `json:"phone" example:"+79161234567"`
	Email string `json:"email" example:"ivan@mail.ru"`
}

var contacts = []Contact{
	{ID: "1", Name: "Иван Иванов", Phone: "+79161234567", Email: "ivan@mail.ru"},
	{ID: "2", Name: "Петр Петров", Phone: "+79169876543", Email: "petr@mail.ru"},
}

// @title Contacts API
// @version 1.0
// @description API для управления телефонными контактами
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/contacts", getContacts)
		v1.GET("/contacts/:id", getContact)
		v1.POST("/contacts", createContact)
		v1.PUT("/contacts/:id", updateContact)
		v1.DELETE("/contacts/:id", deleteContact)
	}

	r.Run(":8080")
}

// GetContacts godoc
// @Summary Получить все контакты
// @Description Возвращает список всех телефонных контактов
// @Tags contacts
// @Produce json
// @Success 200 {array} Contact
// @Router /contacts [get]
func getContacts(c *gin.Context) {
	c.JSON(http.StatusOK, contacts)
}

// GetContact godoc
// @Summary Получить контакт по ID
// @Description Возвращает контакт по указанному ID
// @Tags contacts
// @Produce json
// @Param id path string true "ID контакта"
// @Success 200 {object} Contact
// @Failure 404 {object} map[string]string
// @Router /contacts/{id} [get]
func getContact(c *gin.Context) {
	id := c.Param("id")

	for _, contact := range contacts {
		if contact.ID == id {
			c.JSON(http.StatusOK, contact)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Контакт не найден"})
}

// CreateContact godoc
// @Summary Создать новый контакт
// @Description Создает новый телефонный контакт
// @Tags contacts
// @Accept json
// @Produce json
// @Param contact body Contact true "Данные контакта"
// @Success 201 {object} Contact
// @Failure 400 {object} map[string]string
// @Router /contacts [post]
func createContact(c *gin.Context) {
	var newContact Contact

	if err := c.BindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Генерируем простой ID
	newContact.ID = string(rune(len(contacts) + 1 + '0'))
	contacts = append(contacts, newContact)

	c.JSON(http.StatusCreated, newContact)
}

// UpdateContact godoc
// @Summary Обновить контакт
// @Description Обновляет данные контакта по ID
// @Tags contacts
// @Accept json
// @Produce json
// @Param id path string true "ID контакта"
// @Param contact body Contact true "Обновленные данные контакта"
// @Success 200 {object} Contact
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /contacts/{id} [put]
func updateContact(c *gin.Context) {
	id := c.Param("id")

	var updatedContact Contact
	if err := c.BindJSON(&updatedContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			updatedContact.ID = id
			contacts[i] = updatedContact
			c.JSON(http.StatusOK, updatedContact)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Контакт не найден"})
}

// DeleteContact godoc
// @Summary Удалить контакт
// @Description Удаляет контакт по ID
// @Tags contacts
// @Produce json
// @Param id path string true "ID контакта"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /contacts/{id} [delete]
func deleteContact(c *gin.Context) {
	id := c.Param("id")

	for i, contact := range contacts {
		if contact.ID == id {
			contacts = append(contacts[:i], contacts[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Контакт удален"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Контакт не найден"})
}
