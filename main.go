package main

import (
	"fmt"
	"os"
)

func main() {
	InitSystem()
}

var (
	session    bool
	usertype   string
	userErrors int = 0
	filename   string
)

func InitSystem() {
	fmt.Println("Presiona Para iniciar Sesion: \n 1 Iniciar Usuario \n 2 Iniciar Admin \n 3 cerrar sesion")
	var userInitInput int
	fmt.Scan(&userInitInput)
	switch userInitInput {
	case 1:
		UserSession()
	case 2:
		AdminSession()
	case 3:
		ExitSession()
	}
}

func AdminSession() {
	var admin string = "admin"
	var adminPass string = "admin"
	fmt.Println("Ingresa tu nombre de usuario")
	var inputUser string
	fmt.Scan(&inputUser)
	fmt.Println("Ingresa tu contraseña")
	var inputPass string
	fmt.Scan(&inputPass)

	if admin == inputUser && adminPass == inputPass {
		fmt.Println("\n cargando...")
		fmt.Println("\n Bienvenido...")
		usertype = "admin"
		Adminfunctions()
	} else {
		fmt.Println("\n Verifique sus credenciales e intente nuevamente")
		fmt.Println(userErrors)
		HandlingErrorsSessions()
		AdminSession()
	}

}

func UserSession() {
	var user string = "user"
	var userPass string = "user"
	fmt.Println("Ingresa tu nombre de usuario")
	var inputUser string
	fmt.Scan(&inputUser)
	fmt.Println("Ingresa tu contraseña")
	var inputPass string
	fmt.Scan(&inputPass)

	if user == inputUser && userPass == inputPass {
		fmt.Println("\n cargando...")
		fmt.Println("\n Bienvenido...")
		usertype = "user"
		Userfunctions()
	} else {
		fmt.Println("\n Verifique sus credenciales e intente nuevamente")
		fmt.Println(userErrors)
		HandlingErrorsSessions()
		UserSession()
	}

}

func Adminfunctions() {
	fmt.Println(" \n Admin menu")
	fmt.Println("1 Leer Doc")
	fmt.Println("2 Crear Doc")
	fmt.Println("3 Escribir en Doc")
	fmt.Println("4 Borrar Doc")
	fmt.Println("5 Cambiar Sesion de usuario")
	var adminSelection int
	fmt.Scan(&adminSelection)
	switch adminSelection {
	case 1:
		ReadFile()
	case 2:
		CreateFile()
	case 3:
		WriteFile()
	case 4:
		DeleteFile()
	case 5:
		InitSystem()
	}
}

func Userfunctions() {
	fmt.Println("\n User menu \n")

}

func ReadFile() {
	fmt.Println("Leyendo archivo")
}

func CreateFile() {
	fmt.Println("Ingrese nombre del Archivo")
	fmt.Scan(&filename)
	var filetFormatName string = filename + ".txt"
	var _, err = os.Stat(filetFormatName)
	if os.IsNotExist(err) {
		var file, err = os.Create(filetFormatName)
		if FileErrorFound(err) {
			return
		}
		fmt.Println("archivo " + filetFormatName + " creado")
		filename = ""
		defer file.Close()
		defer Adminfunctions()
	}

}

func WriteFile() {
	fmt.Println("<<< Ingresa Nombre del Archivo donde vas a escribir >>>")
	fmt.Scan(&filename)
	if len(filename) > 1 {

		var file, err = os.OpenFile(filename+".txt", os.O_RDWR, 000)

		if FileErrorFound(err) {
			return
		}

		defer file.Close()

		var userText string
		fmt.Println("Ingresa tu texto")
		fmt.Scan(&userText)
		_, err = file.WriteString(userText)

		if FileErrorFound(err) {
			return
		}

		err = file.Sync()

		if FileErrorFound(err) {
			return
		}

		fmt.Println("actualizado con exito")
	}

}

func DeleteFile() {
	fmt.Println("Eliminando Archivo")
}

func FileErrorFound(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func HandlingErrorsSessions() {
	userErrors++

	if userErrors >= 2 {
		fmt.Println("\n muchos intentos fallidos redirigiendo al menu principal... \n ")
		userErrors = 0
		InitSystem()
	}
}

func ExitSession() int {
	usertype = ""
	fmt.Println("\n Cerrando Sesion \n")
	return 0
}
