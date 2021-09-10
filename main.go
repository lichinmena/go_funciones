package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	saludar_1()

	saludar_2_result := saludar_2()
	fmt.Println(saludar_2_result)

	saludar_3_result := saludar_3("Luis Mena")
	fmt.Println(saludar_3_result)

	saludar_4_result := saludar_4("Luis", "Mena Naal")
	fmt.Println(saludar_4_result)

	saludar_5_result := saludar_5("Luis", "Mena Naal")
	fmt.Println(saludar_5_result)

	saludar_6_result, mayor := saludar_6("Luis", "Mena Naal", 19)
	fmt.Println(saludar_6_result)
	fmt.Println(mayor)

	saludar_7_result, mayor := saludar_7("Luis", "Mena Naal", 10)
	fmt.Println(saludar_7_result)
	fmt.Println(mayor)

	saludar_8_result, _ := saludar_8("Luis", "Mena Naal", 10)
	fmt.Println(saludar_8_result)

	saludar_9_result, _, err := saludar_9("Luis", "RAmns", 10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(saludar_9_result)
	//El estado de salida para una ejecucion exitosa es 0, cualquier otra indica un error

	saludar_10("Luis", "Fernando", "Landy", "Isabel", "Camila", "Emma") //[]string ==> slice

	sumar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

}

//Forma basica
func saludar_1() {
	fmt.Println("saludar_1 :: Hola mi gente")
}

//Con retorno
//Determinamos en la firma de la funcion el tipo de dato
//que retornara
func saludar_2() string {
	return "saludar_2 :: Hola desde una funcion con return"
}

//La firma de una funcion puede tener parametros
// es decir valores que espera una funcion
//****** En go no podemos asignar valores por defecto
func saludar_3(nombre string) string {
	return fmt.Sprintf("Saludar_3 :: Hola %s", nombre)
}

func saludar_4(nombre string, apellido string) string {
	return fmt.Sprintf("Saludar_4 :: Hola %s %s", nombre, apellido)
}

//Con parametros del mismo tipo
func saludar_5(nombre, apellido string) string {
	return fmt.Sprintf("Saludar_5 :: Hola %s %s", nombre, apellido)
}

//Las funciones en go pueden retornar multiples valores
//El orden es muy importante,por que tal cual lo retornara
func saludar_6(nombre, apellido string, edad uint8) (string, bool) {
	booleano := edad >= 18
	texto := fmt.Sprintf("Saludar_6 :: Hola %s %s", nombre, apellido)
	return texto, booleano
}

//Go nos permite nombrar los valores de retorno en la firma dela funcion
func saludar_7(nombre, apellido string, edad uint8) (texto string, booleano bool) {
	booleano = edad >= 18
	texto = fmt.Sprintf("Saludar_7 :: Hola %s %s", nombre, apellido)
	return //Puede estar solo, por que se determina en la firma el retorno
}

//Para ignorar uno de los valores que retorna usamos el guion bajo
func saludar_8(nombre, apellido string, edad uint8) (texto string, booleano bool) {
	booleano = edad >= 18
	texto = fmt.Sprintf("Saludar_8 :: Hola %s %s", nombre, apellido)
	return //Puede estar solo, por que se determina en la firma el retorno
}

//Retornar un error(cuando es de multiples valores), ya que las funciones pueden retornar un error
//Normalmente esta al final
func saludar_9(nombre, apellido string, edad uint8) (texto string, booleano bool, err error) {
	booleano = edad >= 18
	texto = fmt.Sprintf("Saludar_9 :: Hola %s %s", nombre, apellido)

	if len(apellido) <= 2 {
		err = fmt.Errorf("El apellido '%s' no es válido", apellido)
	}
	return //Puede estar solo, por que se determina en la firma el retorno
}

/*Las funciones en go son un tipo de dato, podemos asignarlas a variables*/

//miFuncion := saludar
//saludo,_,err := miFuncion("Luis", "Mena",19)

//las funciones como tipo de dato no son comparables
//No podemos usarlos como llaves de mapa

//Determina el tipo de dato
//fmt.Println("saludo1:", reflect.TypeOf(saludo1))

//Funciones variaticas
//Go permite recibir un numero  variable de parametros en una funcion
func saludar_10(nombres ...string) {
	fmt.Println(reflect.TypeOf(nombres))

	for _, nombre := range nombres {
		fmt.Println("Hola", nombre)
	}
}

/*Declarar un slice*/
// personas := []string{"Uno", "Dos", "Tres"}
//Usar los valores del slice
//saludar(personas...)

func sumar(numeros ...int) (sumatoria int) {
	fmt.Println(reflect.TypeOf(numeros))
	for _, n := range numeros {
		sumatoria += n
	}
	fmt.Println("La suma es: ", sumatoria)
	return
}
