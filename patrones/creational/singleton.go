package main

import (
	"fmt"
	"sync"
	"time"
)

// Patron de diseño creacional que se asegura que solo exista una instancia de una clase

type Database struct{}

func (Database) GetConnection() {
	println("Conectando a base de datos")
	time.Sleep(2 * time.Second)
	println("Conexion establecida")
}

var db *Database
var lock sync.Mutex // Mutex para evitar que se cree más de una instancia de la base de datos

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()

	// No hay instancia, la creamos
	if db == nil {
		fmt.Println("Creando instancia de base de datos")
		db = &Database{}
		db.GetConnection()
	} else {
		fmt.Println("Usando instancia existente")
	}
	return db
}

func main() {
	var wg sync.WaitGroup

	// Lanzamos 10 gorutinas para pedir la instancia de la base de datos
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}

	wg.Wait()
}
