package model

/**
    @date: 2024/7/13
**/

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
}
