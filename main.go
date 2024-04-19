package main

import (
	"fmt"
	"os"
)

func main() {
	InitSystem()
}

var (
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
		CloseSession()
	}
}

func AdminSession() {
	/*Admin Credentials Hardcoded*/
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
	/* user credentials harcoded*/
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
	fmt.Println("6 Finalizar Sesion y salir")
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
	case 6:
		CloseSession()
	}
}

func Userfunctions() {
	fmt.Println("\n User menu \n")
	fmt.Println(" 1- leer archivo \n 2- cambiar sesion\n 3 cerrar sesion")
	var userActionSelect int
	fmt.Scan(&userActionSelect)

	switch userActionSelect {
	case 1:
		ReadFile()
	case 2:
		InitSystem()
	case 3:
		CloseSession()
	}

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

	/* validating lenght name*/

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
		Adminfunctions()
	}

}

func DeleteFile() {
	fmt.Println("Ingrese el nombre del archivo que desea eliminar")
	var userFileToDelete string
	fmt.Scan(&userFileToDelete)
	err := os.Remove(userFileToDelete + ".txt")

	if err != nil {
		fmt.Println("Error:", err)
		Adminfunctions()
	} else {
		fmt.Println("archivo borrado", userFileToDelete+".txt\n")
		Adminfunctions()
	}

}

func FileErrorFound(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func ReadFile() {
	fmt.Println("Ingresa nombre del archivo a leer")
	/* taking user input*/
	var toReadFile string
	fmt.Scan(&toReadFile)

	/* handling file to read in console*/
	file, err := os.ReadFile(toReadFile + ".txt")
	if err != nil {
		fmt.Println("Error: ", err)
		UserTypeValidation()
	} else {
		os.Stdout.Write(file)
		UserTypeValidation()
	}
}

func UserTypeValidation() {
	/* check usertype from global variable */
	if usertype == "admin" {
		Adminfunctions()
	} else {
		Userfunctions()
	}
}

func HandlingErrorsSessions() {
	/* register of users attemps to login */
	userErrors++
	/* Handling errors logins limits  */
	if userErrors >= 2 {
		fmt.Println("\n muchos intentos fallidos redirigiendo al menu principal... \n ")
		userErrors = 0
		InitSystem()
	}
}

func CloseSession() int {
	/* reset variables to finish session */
	usertype = ""
	filename = ""
	fmt.Println("\n Cerrando Sesion \n")
	return 0
}
