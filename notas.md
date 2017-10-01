# Características de GoLang

## Paquetes

```go
package main /* Define un paquete que puede ser compilado y ejecutado */
```

```go
package anything /* Define una dependencia compilable */
```
## Compilación y ejecución

```bash
go build main.go # Compila y genera un ejecutable a partir de los archivos .go incluidos
```

```bash
go run main.go # Compila y ejecuta inmediatamente el código incluido 
```

## Tipado

El tipado en Go es **estático**. O sea que en tiempo de compilación ya se conocen los tipos de todas las variables y argumentos de funciones.

## Tipos de variables

Tipos básicos:
- string
- int
- float64
- array (array estático)
- slice[type] (array dinámico), con valores de tipo "type"

Otros:
- byte (representación en bytes de una variable)

## Conversión de tipos

Ejemplo de conversión de string a byte:
```Go
func main() {
	greeting := "Hi there!" 
	// Go implícitamente entiende que *greeting* es un string

	fmt.Println([]byte(greeting))
	// Imprime => [72 105 32 116 104 101 114 101 33]
}
```
Esto es similar a lo conocido como casteo.

## Tipos por valor, por referencia, y punteros

Como sucede en otros lenguajes **siempre que una función recibe un argumento, en la ejecución lo que le llega es una copia del argumento**.

### Estructuras, punteros y _receiver functions_

Go permite escribir funciones llamadas **receiver functions, aplicadas a los tipos y estructuras de datos que se definan**.

Ahora, debido al comportamiento de las funciones que reciben una copia de sus argumentos, **si una estructura se modifica en un una _reciever function_,** para que sea modificada la estructura original __se debe definir que el tipo a recibir sea un puntero a la instancia de la estructura. *De lo contrario sólo se modificaría la copia de la estructura o tipo que recibe la función.*__ Go simplifica el tratamiento de punteros requiriendo sólamente especificar los punteros en la _receiver function_. Ejemplo:

```Go
type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
	}
	jim.updateName("jimmy")
	fmt.Printf("%v", jim)
	// Se imprime >> {jimmy Party}
}

// Se define la receiver function como receptora de un puntero a "person"
func (p *person) updateName(newFirstname string) {
	(*p).firstName = newFirstname
}
```

### Tipos por valor y por referencia:


**Tipos por valor:**

Los tipos por valor se comportan como las estructuras: **si una función los modifica, la función debe recibir un puntero, o directamente retornar el nuevo valor.**

- int
- float
- string
- bool
- structs

**Tipos por referencias:**

Su definición consta de la estructura interna del tipo, más una referencia. Cuando se utilizan estos tipos **ya se está operando sobre un puntero.** Por ende **no hace falta usar punteros en las funciones que modifican este tipo de variables.**

- slices
- maps
- channels
- pointers
- functions

## Paradigma

Definitivamente Go **no** es un lenguaje orientado a objetos. 

Provee ciertas herramientas que permiten al desarrollador familiarizado con POO definir tipos y **estructuras con datos y comportamiento asociados**.

También implementa algunos features típicos de los lenguajes orientados a objetos, tales como las **_interfaces_**.

```Go
// Definición de la interfaz
type bot interface {
	getGreeting() string
}

// Definición de los tipos que implementan las funciones de la interfaz
type englishBot struct{}
type spanishBot struct{}

// Implementaciones
func (englishBot) getGreeting() string {
	return "Hi There!"
}
func (spanishBot) getGreeting() string {
	return "Hola!"
}

// Función que acepta tipos comprendidos en la nueva interfaz
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
// Ahora se puede llamar a las funciones:
// - printGreeting(eb englishBot) 
// - printGreeting(sb spanishBot)
```

El comportamiento del código anterior es similar a cuando en POO dos clases _implementan_ una funcionalidad encerrada en una **interfaz**.

## Misc

- Go permite definir funciones con más de un retorno

```Go
// Se define la funcion que retorna multiples valores
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func main(){
    // Llamado a la funcion capturando todos los valores de retorno
    hand, remainingCards := deal(cards, 5)
}
```