package main

import "fmt"

func main() {
	fmt.Println("")
	fmt.Println("Println = Exibe um texto, ou uma variável, ou até mais de uma variável e pula para o próximo comando (porém aceita ou somente texto ou somente variáveis);")
	fmt.Println("Print = Exibe um texto e o próximo comando se mantém na mesma linha;")
	fmt.Println("Printf = Exibe um texto que contem uma ou mais variáveis e estas serão formatadas para texto (format string). O conceito de format strings funciona como se a string fosse um link para a variável, mas para entender melhor esse conceito, veja os")
	fmt.Println("")

	var x int = 10
	var y float32 = 20.5
	var z string = "FIM"
	fmt.Print("O próximo comando virá aqui neste mesma linha => ")
	fmt.Println("Posso até estar na mesma linha, mas o próximo comando vai para a linha de baixo!")
	fmt.Printf("Os valores de: x = %d | y = %f | z = %s", x, y, z)

}
