package chapter_3

import "testing"

func TestPets(t *testing.T) {
	p := GetPets()
	p.Enqueue(dog, 1)
	p.Enqueue(cat, 1)
	p.Enqueue(dog, 2)
	p.Enqueue(cat, 2)
	p.Enqueue(dog, 3)
	p.Enqueue(cat, 3)
	p.Enqueue(dog, 4)

	if pet, _ := p.DequeueDog(); pet != 1 {
		t.Errorf("Expected: 1. Actual: %d", pet)
	}
	if pet, _ := p.DequeueCat(); pet != 1 {
		t.Errorf("Expected: 1. Actual: %d", pet)
	}
}
