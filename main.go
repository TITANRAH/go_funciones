package main

import (
	"fmt"
)

func main() {

	// interdace

	e := Estructura{}
	resultado := e.Calculo(1,2)

	fmt.Println("resultado funcion de interface",resultado)


	// miFuncion()

	// "ejemplo gorutine 1 o como funciona"
	// fmt.Println(miFuncionNueva("Sergio"))
	// time.Sleep(time.Second * 5)
	// fmt.Println(miFuncionNueva("Maria"))

	// ejemplo real de gorutine
	// chanel
	// miCanal := make(chan string)

	// go func() {
	// 	// te sirve para generar diferente funciones todas en distinto tiempos o sucesivas
	// 	// obtejer cliente por ejem luego ejecutar otra funcion con los datos del cliente y asi sucesivamente
	// 	miCanal <- miFuncionNueva("Pedro")
	// }()

	// fmt.Println(<-miCanal)

	// recursividad LLAMAR A LA FUNCION DENTRO DE SI MISMA

	// miFuncionRecursiva(0)

}

func miFuncionRecursiva(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncionRecursiva(dato)
	}
}

func miFuncionPanicDefer() {
	fmt.Println("Este es mi primer mensaje")
}

func miFuncionNueva(parametro string) string {
	return "hola " + parametro
}

func miFuncion() {


	// interfaces 
	// como declare una estrucutura y luego declare una interfaces quee contiene metodos que 
	// puedo asociar a dicha estructura ,  ahora puedo instanciar la estrucutra y llamar a los metodos de la interfface

	// estrucutura


	// estructura := Persona{
	// 	Id:     1,
	// 	Nombre: "Sergio",
	// 	Correo: "info@sergio.com",
	// 	Edad:   45,
	// }

	// fmt.Println(estructura)

	// // nueva instancia de Persona
	// p := new(Persona)

	// p.Id = 1
	// p.Correo = "jasdjasdi@asdasdl.com"
	// p.Nombre = "Juan"
	// p.Edad = 41

	// fmt.Println(p)

	// variable tipo valores
	// categoria := Categoria{Id: 1, Nombre: "Categoría 1", Slug: "categoria-1"}
	// producto := Producto{Id: 3, Nombre: "Mesa de computador", Slug: "mesa-de-computador", Precio: 1231, CategoriaId: categoria}

	// fmt.Printf("%+v \n", producto)

	// interfaces

	// defer fmt.Println("MENSAJE FINAL")
	// fmt.Println("PRIMER MENSAJE")
	// fmt.Println("Hola mundo")
	// miFunctionConParametros(6, 2)
	// nombre, apellido, numero := miFuncionConRetornoMultiple("Sergio", "Miranda", 45)

	// fmt.Printf("Hola %v %v, tu edad es %v", nombre, apellido, numero)
	// fmt.Println("La suma es", suma(10, 12))

	// num := 6

	// Tabla := tabla(num)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("%v x %v = %v \n", num, i, Tabla())
	// }
	// a := 1
	// if a == 1 {
	// 	// panic detiene la continuidad de la funcion
	// 	panic("fallo por que fallo")
	// }
}

func miFunctionConParametros(n1 int, n2 int) {

	resultado := n1 + n2
	fmt.Println("El resultado es ", resultado)

	fmt.Println(miFuncionConRetorno("Sergio"))

}

// interfaces

// una interdaces es el contrato de metodos de una estructura
type EjemploInterface interface {
	miFuncionInterface() string
	Calculo(n1 int, n2 int) int
}

type Estructura struct {
}

func (*Estructura) miFuncionInterface() string {
	return "Texto desde mi funcion"
}

func  (*Estructura) Calculo(n1 int,n2 int) int {
	return n1 + n2
}

// estructura es parecido a interface
type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

// estructura anidada
type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}

func miFuncionConRetorno(nombre string) string {

	return "tu nombre es" + nombre

}

func miFuncionConRetornoMultiple(nombre string, apellido string, edad int) (string, string, int) {

	return nombre, apellido, edad

}

// funcion anonima genrealmente son variables
var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// funciuon clousure , siempre retornara una funcion
// como esta siendo recorrida con un for aumentara segun los ciclos declarados en el for
func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
