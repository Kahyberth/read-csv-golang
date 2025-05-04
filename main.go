package main

import (
	"fmt"
	_func "github.com/Kahyberth/read-csv-golang/func"
)

func main() {

	//conn, err := pgx.Connect(context.Background(), "postgres://postgres:admin@localhost:5432/readCV")
	//
	//if err != nil {
	//	fmt.Println("Error conectando a la base de datos:", err)
	//	return
	//}
	//defer conn.Close(context.Background())

	//TODO: Modificar
	err := _func.ReadCSVOne()

	if err != nil {
		fmt.Println("Error leyendo el CSV:", err)
		return
	}

	//
	//err = lib.UploadUsersToDB(conn, usuarios, 1000)
	//
	//if err != nil {
	//	fmt.Println("Error al intentar subir los datos a la base de datos", err)
	//	return
	//}
	//
	//fmt.Println("Usuarios cargados correctamente!")

}
