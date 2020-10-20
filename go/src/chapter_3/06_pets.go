// В приюте для животных есть только собаки и кошки, а работа осуществляется в порядке очереди.
// Люди должны каждый раз забирать «самое старое» (по времени пребывания в питомнике) животное,
// но могут выбрать кошку или собаку (животное в любом случае будет «самым старым»).
// Нельзя выбрать любое понравившееся животное. Создайте структуру данных, которая
// обеспечивает функционирование этой системы и реализует операции enqueue, dequeueAny,
// dequeueDog и dequeueCat. Разрешается использование встроенной структуры данных LinkedList.
package chapter_3

import "math/rand"

type Pet int

const (
	dog Pet = iota
	cat
)

type Pets struct {
	dogs *Queue
	cats *Queue
}

func GetPets() *Pets {
	return &Pets{GetQueue(), GetQueue()}
}

func (p *Pets) Enqueue(pet Pet, name int) {
	if pet == dog {
		p.dogs.Add(name)
	}
	if pet == cat {
		p.cats.Add(name)
	}
}

func (p *Pets) DequeueDog() (int, error) {
	return p.dogs.Remove()
}

func (p *Pets) DequeueCat() (int, error) {
	return p.cats.Remove()
}

func (p *Pets) DequeueAny() (int, error) {
	if rand.Intn(10) < 5 {
		return p.dogs.Remove()
	} else {
		return p.cats.Remove()
	}
}
