package provider

import (
	"os"
	"strconv"
	"task-uxbert/seeder"
)

/***
* seed and truncate data if  TRUNCATE_ALL_TABLES true
 */
func Seed() {
	deleteTables, _ := strconv.ParseBool(os.Getenv("TRUNCATE_ALL_TABLES"))
	if deleteTables {
		truncate("users", "messages", "room")
		seeder.UserSeeder()
	}
}
