package main

// Vertice = id , estado (aberto/fechado)

// arestas
// vertice de saida, vertice de chegada e peso
import (
	"errors"
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
	precedente *Vertice
}

type Aresta struct {
	//posso usar ids mas prefiri ponteiros
	id             byte
	entrada, saida *Vertice
	peso           int32
}

type Grafo struct {
	dirigido bool
	raiz     *Vertice
	vertices []*Vertice
	arestas  []*Aresta
}

//criar grafo grafo
func CriarGrafo(dirigido bool) (grafo *Grafo) {
	grafo = &Grafo{dirigido: dirigido}
	return
}

var ultimoVertice byte = 'a'
var ultimoAresta byte = 'A'

func (g *Grafo) AdicionarVertice() *Vertice {
	v := &Vertice{id: ultimoVertice, estado: VerticeAberto, estimativa: math.MaxInt32}
	ultimoVertice++
	g.vertices = append(g.vertices, v)
	return v
}
func (g *Grafo) AdicionarAresta(v1 *Vertice, v2 *Vertice, peso int32) {
	a := &Aresta{id: ultimoAresta, entrada: v1, saida: v2, peso: peso}
	if g.dirigido {

		g.arestas = append(g.arestas, a)
	} else {
		b := &Aresta{id: ultimoAresta, entrada: v2, saida: v1, peso: peso}
		g.arestas = append(g.arestas, a, b)

	}
	ultimoAresta++
}

func (g *Grafo) fecharVertice(vertice *Vertice) {
	for _, aresta := range g.arestas {
		//aresta adjacente
		if aresta.entrada.id == vertice.id {
			estimativa := aresta.entrada.estimativa + aresta.peso
			if aresta.saida.estimativa > estimativa {
				aresta.saida.estimativa = estimativa
				aresta.saida.precedente = aresta.entrada
			}
		}
	}
	vertice.estado = VerticeFechado
}
func (g *Grafo) restaVerticeAberto() bool {
	for _, v := range g.vertices {
		if v.estado == VerticeAberto {
			return true
		}
	}
	return false
}
func (g *Grafo) Dijkstra() error {
	if g.raiz == nil {
		return errors.New("Raiz não definida")
	}
	g.raiz.estimativa = 0
	g.fecharVertice(g.raiz)
	for g.restaVerticeAberto() {
		for _, v := range g.vertices {
			if v.precedente != nil && v.estado == VerticeAberto {
				g.fecharVertice(v)
			}
		}
	}
	return nil
}

func (g *Grafo) MenorCaminho(v *Vertice) string {
	if v.precedente == nil {
		return fmt.Sprintf("%c", v.id)
	}
	return fmt.Sprintf("%c -> ", v.id) + g.MenorCaminho(v.precedente) //+ g.MenorCaminho(v.precedente)
}

func (v *Vertice) String() string {
	estadoStr := [...]string{"aberto", "fechado"}
	var precedente byte = '.'
	if v.precedente != nil {
		precedente = v.precedente.id
	}
	return fmt.Sprintf("%c [%s] : estimativa %d precedente %c\n", v.id, estadoStr[v.estado], v.estimativa, precedente)
}
func (g *Grafo) String() string {
	return fmt.Sprintf("Raiz %s \n-> Arestas %d \n-> Vertices %d\n %v", g.raiz, len(g.arestas), len(g.vertices), g.vertices)
}

func main() {
	grafo := CriarGrafo(false)
	//exemplo slides
	a := grafo.AdicionarVertice()
	b := grafo.AdicionarVertice()
	c := grafo.AdicionarVertice()
	d := grafo.AdicionarVertice()
	e := grafo.AdicionarVertice()
	f := grafo.AdicionarVertice()
	//
	grafo.AdicionarAresta(a, b, 2)
	grafo.AdicionarAresta(a, d, 1)
	grafo.AdicionarAresta(a, c, 5)

	grafo.AdicionarAresta(b, c, 3)
	grafo.AdicionarAresta(b, d, 2)

	grafo.AdicionarAresta(c, d, 3)
	grafo.AdicionarAresta(c, e, 1)
	grafo.AdicionarAresta(c, f, 5)

	grafo.AdicionarAresta(e, d, 1)
	grafo.AdicionarAresta(e, f, 2)
	//seleciono raiz do nodo
	grafo.raiz = a
	//rodo algoritmo de Dijkstra
	grafo.Dijkstra()

	fmt.Println(grafo)

	fmt.Println("Menor caminho até nodo f é ", f.estimativa, grafo.MenorCaminho(f))
}
