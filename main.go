package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\x1b[42m------------------------------\x1b[0m")
	fmt.Println("\x1b[42mDijkstra Grafo Generator 2000 \x1b[0m")
	fmt.Println("\x1b[42m------------------------------\x1b[0m\n")
	fmt.Print("É um grafo dirigido?[yes/no]\n")
	resp, _, _ := reader.ReadLine()
	var dirigido bool
	if string(resp) == "yes" {
		fmt.Println("Ok! teremos um grafico dirigido.")
		dirigido = true
	}

	g := CriarGrafo(dirigido)

	fmt.Print("Quantos vertices?")
	resp, _, _ = reader.ReadLine()

	vertices, err := strconv.Atoi(string(resp))
	if err != nil {
		fmt.Println("NÚMERO INVALIDO")
		os.Exit(1)
	}
	fmt.Printf("\n\nAdicionando %d vertices :\n", vertices)
	for i := 0; i < vertices; i++ {
		fmt.Printf("%c,", 'a'+i)
		g.AdicionarVertice()
	}
	fmt.Println("\n\nVamos adicionar as arestas!")
	fmt.Println("Basta digitar \"[VerticeDeEntrada],[VerticeDeSaida],[Peso]\"")
	fmt.Println("Exemplo: \x1b[35m\"a,b,2\"\x1b[0m (se o peso for 1, pode omitir)")
	fmt.Println("Para finalizar digite ok")
	for resp, _, err := reader.ReadLine(); string(resp) != "ok" && err == nil; resp, _, err = reader.ReadLine() {
		respStr := string(resp)
		respostas := strings.Split(respStr, ",")
		switch len(respostas) {
		case 2:
			g.AdicionarAresta(g.vertices[respostas[0][0]-'a'], g.vertices[respostas[1][0]-'a'], 1)
		case 3:
			peso, _ := strconv.Atoi(respostas[2])
			g.AdicionarAresta(g.vertices[respostas[0][0]-'a'], g.vertices[respostas[1][0]-'a'], int32(peso))
		default:
			fmt.Println("Não entendi seu padrão ;)")
		}
	}

	fmt.Println("Rodando algoritmo de Dikstra")
	start := time.Now()
	err = g.Dijkstra()
	if err != nil {
		fmt.Println("Grafo Invalido")
		os.Exit(1)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println(g)
	fmt.Printf("Tempo total: %d ns\n\n\n", elapsed.Nanoseconds())
	fmt.Println("Até qual vertice você quer saber o caminho?")
	fmt.Println("Para finalizar digite ok")
	for resp, _, err := reader.ReadLine(); string(resp) != "ok" && err == nil; resp, _, err = reader.ReadLine() {
		v := g.vertices[resp[0]-'a']
		fmt.Printf("Até o nodo %c a estimativa é de %d - %s\n", v.id, v.estimativa, g.MenorCaminho(v))

	}
}
