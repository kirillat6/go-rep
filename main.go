package main

import (
	"context"
	"fmt"
	"study/feature_postgres/simple_connection"
	"study/feature_postgres/simple_sql"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)

	if err != nil {
		panic(err)
	}

	if err := simple_sql.CreateTable(ctx, conn); err !=nil {
		panic(err)
	} 	
	if err := simple_sql.InsertRow(
		ctx,
		conn,
		"Обед", 
		"Покушац надо", 
		false,
		time.Now(),
		); err !=nil {
		panic(err)
	} 

	// if err := simple_sql.UpdateRow(
	// ctx,
	// conn,
	// 2,
	// ); err !=nil {
	// 	panic(err)
	// } 
	
	// if err := simple_sql.DeleteRow(ctx, conn, 2); err != nil {
	// 	panic(err)
	// }

	// tasks, err := simple_sql.SelectRows(ctx, conn) 
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("SUCCEED!")
}