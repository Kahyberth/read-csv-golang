package main

import (
	"context"
	"fmt"
	"github.com/Kahyberth/read-csv-golang/lib"
	"github.com/jackc/pgx/v5"

	_func "github.com/Kahyberth/read-csv-golang/func"
)

func main() {

	conn, err := pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/readCV")
	if err != nil {
		fmt.Println("Error conectando a la base de datos:", err)
		return
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())

	usuarios, err := _func.ReadUtils()
	if err != nil {
		fmt.Println("Error leyendo el CSV:", err)
		return
	}

	err = lib.UploadUsersToDB(conn, usuarios, 1000)

	if err != nil {
		fmt.Println("Error al intentar subir los datos a la base de datos", err)
		return
	}

	fmt.Println("Usuarios cargados correctamente!")

}
