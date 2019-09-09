package main

// Vertice = id , estado (aberto/fechado)

// arestas
// vertice de saida, vertice de chegada e peso
import (
	"fmt"
	"math"
)

type Estado int

const (
	VerticeAberto  Estado = 0
	VerticeFechado Estado = 1
)

type Vertice struct {
	id         byte //vai ser char
	estado     Estado
	estimativa int32
}

type Aresta struct {
	//posso usar ids mas prefiri ponteiros
	id             byte
	entrada, saida *Vertice
	peso           int32
}

type Grafo struct {
	raiz     *Vertice
	vertices []Vertice
	arestas  []Aresta
}

//criar grafo grafo
func criarGrafo() (grafo *Grafo) {
	grafo = &Grafo{}
	return
}

var ultimoVertice byte = 'a'
var ultimoAresta byte = 'A'

func (g *Grafo) adicionarVertice() *Vertice {
	v := Vertice{id: ultimoVertice, estado: VerticeAberto, estimativa: math.MaxInt32}
	ultimoVertice++
	g.vertices = append(g.vertices, v)
	return &v
}
func (g *Grafo) adicionarAresta(v1 *Vertice, v2 *Vertice, peso int32) *Aresta {
	a := Aresta{id: ultimoAresta, entrada: v1, saida: v2, peso: peso}
	ultimoAresta++
	g.arestas = append(g.arestas, a)
	return &a
}

func (v *Vertice) String() string {
	estadoStr := []string{"aberto", "fechado"}
	return fmt.Sprintf("%c [%s] : estimativa %d \n", v.id, estadoStr[v.estado], v.estimativa)
}
func (g *Grafo) String() string {
	return fmt.Sprintf("Raiz %s \n-> Arestas %d %v \n-> Vertices %d %v", g.raiz, len(g.arestas), g.arestas, len(g.vertices), g.vertices)
}

//adicionar vertice em grafo

// adicionar aresta em vertices

func main() {
	grafo := criarGrafo()
	v1 := grafo.adicionarVertice()
	v2 := grafo.adicionarVertice()
	grafo.adicionarAresta(v1, v2, 20)
	fmt.Println(grafo)
}
