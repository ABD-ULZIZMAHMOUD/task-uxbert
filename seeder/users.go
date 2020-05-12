package seeder

import (
	"syreclabs.com/go/faker"
	"task-uxbert/config"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
)

/**
* fake data and create user
 */
func newUser(admin bool) {
	hash, _ := helpers.HashPassword("123456789")
	user := models.User{
		Email:    faker.Internet().Email(),
		Password: hash,
		FullName: faker.Name().Name(),
	}
	// if admin true create admin user else create normal user
	if admin {
		user.Type = 1
	} else {
		user.Type = 2
	}
	config.Db.Create(&user)
}

/***
*	Seed Normal users and admins
 */
func UserSeeder() {
	// seed super admin
	SeedSuperAdmin()
	// seed admin users
	for i := 0; i < 5; i++ {
		newUser(true)
	}
}

/***
*	Seed Super Admin
 */
func SeedSuperAdmin() {
	hash, _ := helpers.HashPassword("123456789")
	user := models.User{
		FullName: "admin",
		Email:    "admin@admin.com",
		Password: hash,
		Type:     1,
	}
	config.Db.Create(&user)
}
