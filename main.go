package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Kahyberth/read-csv-golang/lib"
	"github.com/Kahyberth/read-csv-golang/model"
	"github.com/jackc/pgx/v5"
	"os"
	"time"

	_func "github.com/Kahyberth/read-csv-golang/func"
)

func main() {
	// Defino unos flags para elegir qué method usar
	methodPtr := flag.String("metodo", "", "Método de lectura: gocarina, encoding o util")
	flag.Parse()

	// Valid0 el method elegido
	if *methodPtr == "" {
		fmt.Println("Debes especificar un método. Opciones: gocarina, encoding, util")
		flag.PrintDefaults()
		return
	}

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

	var usuarios []model.User
	var readTime time.Duration

	// Ejecuta solo el method elegido
	switch *methodPtr {
	case "gocarina":
		usuarios, readTime, err = _func.ReadCSVOcarina()
	case "encoding":
		usuarios, readTime, err = _func.ReadCSVOne()
	case "util":
		usuarios, readTime, err = _func.ReadUtils()
	default:
		fmt.Printf("Método %s no válido. Opciones: ocarina, encoding, util\n", *methodPtr)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error leyendo el CSV:", err)
		return
	}

	uploadTime, err := lib.UploadUsersToDB(conn, usuarios, 1000)
	if err != nil {
		fmt.Println("Error al intentar subir los datos a la base de datos", err)
		return
	}

	fmt.Println("Usuarios cargados correctamente!")
	totalTime := readTime + uploadTime

	fmt.Printf("Tiempo de lectura: %.2f s\n", readTime.Seconds())
	fmt.Printf("Tiempo de carga: %.2f s\n", uploadTime.Seconds())
	fmt.Printf("Tiempo total: %.2f s\n", totalTime.Seconds())

}
