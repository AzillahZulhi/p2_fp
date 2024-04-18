package main

import "p2-fp/config"

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/p2_graded_challange_3?charset=utf8mb4&parseTime=True&loc=Local"
	config.DatabaseConnect(dsn)
}
