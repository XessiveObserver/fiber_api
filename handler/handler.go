package handler

import (
	"github.com/XessiveObserver/fiber_api/database"
	"github.com/XessiveObserver/fiber_api/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create User
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	// Store body in a user and return error if any
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something is wrong with your input", "data": err})
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	// Return the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has been created", "data": user})
}

// Get all users
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db
	var users []model.User

	// Find all users in the database
	db.Find(&users)

	// If no user found return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	// Return users
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users Found", "data": users})
}

// Get single user from database
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db

	// Get id parama
	id := c.Params("id")

	var user model.User

	// Find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// Update user
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Username string `json:"username"`
	}
	db := database.DB.Db
	var user model.User

	// get id params
	id := c.Params("id")

	// Find a single user in the database by Id
	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something is wrong with your input", "data": err})
	}

	user.Username = updateUserData.Username

	// Save the Changes
	db.Save(&user)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user Found", "data": user})
}

// delete user in the db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User

	// get id params
	id := c.Params("id")

	// Find single user in the database by id
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
