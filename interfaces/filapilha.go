package main

import (
	"fmt"
)

type Structure interface {
	add(numero int)
	print()
	remove()
}

type Fila struct {
	items []int
}

func (f *Fila) add(numero int) {
	f.items = append(f.items, numero)
}

func (f *Fila) remove() {
	if len(f.items) > 1 {
		f.items = f.items[0 + 1 : len(f.items)]
	} else {
		f.items = []int{}
	}
}

func (f *Fila) print() {
}

type Pilha struct {
	items []int
}

func (p *Pilha) add(numero int) {
	p.items = append(p.items, numero)
}
func (p *Pilha) remove() {
	if len(p.items) > 1 {
		p.items = p.items[0: len(p.items) - 1]
	} else {
		p.items = []int{}
	}
}

func (p *Pilha) print() {}

func incluirNaLoterica(s Structure, numero int) {
	s.add(numero)
}

func sairDaLoterica(s Structure) {
	s.remove()
}

func addPratoNaPia(s Structure, numero int) {
	s.add(numero)
}

func lavaPrato(s Structure) {
	s.remove()
}

func main() {
	fmt.Println("==== Fila da Loterica ====")
	fila := Fila{}

	incluirNaLoterica(&fila, 10)
	incluirNaLoterica(&fila, 1)
	incluirNaLoterica(&fila, 2)
	incluirNaLoterica(&fila, 3)

	fmt.Println(fila)
	
	sairDaLoterica(&fila)
	sairDaLoterica(&fila)
	fmt.Println(fila)
	
	fmt.Println("==== Pilha de Louças ====")
	pilha := &Pilha{}
	addPratoNaPia(pilha, 10)
	addPratoNaPia(pilha, 11)
	addPratoNaPia(pilha, 12)
	addPratoNaPia(pilha, 14)
	addPratoNaPia(pilha, 13)
	fmt.Println(*pilha)
	
	lavaPrato(pilha)
	lavaPrato(pilha)
	lavaPrato(pilha)
	fmt.Println(*pilha)
}

