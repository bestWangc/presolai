package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./internal/models",
	})

	gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/presol?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		g.GenerateModel("users"),

		// Generate struct `Employee` based on table `users`
		// g.GenerateModelAs("users", "Employee"),

		// Generate struct `User` based on table `users` and generating options
		// g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),

		// Generate struct `Customer` based on table `customer` and generating options
		// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
		// g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

