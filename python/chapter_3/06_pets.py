# В приюте для животных есть только собаки и кошки, а работа осуществляется
# в порядке очереди. Люди должны каждый раз забирать «самое старое»
# (по времени пребывания в питомнике) животное, но могут выбрать кошку
# или собаку (животное в любом случае будет «самым старым»). Нельзя выбрать
# любое понравившееся животное. Создайте структуру данных, которая
# обеспечивает функционирование этой системы и реализует операции
# enqueue, dequeueAny, dequeueDog и dequeueCat. Разрешается использование
# встроенной структуры данных LinkedList.
class LinkedListException(Exception):
    pass


class Pet:
    def __init__(self, name):
        self.name = name
        self.next = None
        self.prev = None


class PetQueue:
    def __init__(self):
        self.first_pet = None
        self.last_pet = None
        self.length = 0

    def __str__(self):
        if self.first_pet is None:
            return ""
        pets = []
        pet = self.first_pet
        while pet:
            pets.append(pet.name)
            pet = pet.next
        return ",".join(pets)

    def push(self, name):
        pet = Pet(name)
        self.length += 1
        if self.length == 1:
            self.first_pet = self.last_pet = pet
            return
        pet.next = self.first_pet
        self.first_pet.prev = pet
        self.first_pet = pet

    def pop(self):
        if self.length == 0:
            raise LinkedListException("Cannot pop. Queue is empty.")
        self.length -= 1
        name = self.last_pet.name
        if self.length == 0:
            self.first_pet = self.last_pet = None
        else:
            n = self.last_pet.prev
            self.last_pet = n
            self.last_pet.next = None
        return name


class Pets:
    def __init__(self):
        self.dogs = PetQueue()
        self.cats = PetQueue()

    def __str__(self):
        return f"Dogs: {str(self.dogs)}\nCats: {str(self.cats)}"

    def enqueue(self, pet, name):
        if pet == "dog":
            self.dogs.push(name)
        if pet == "cat":
            self.cats.push(name)

    def dequeueDog(self):
        return self.dogs.pop()

    def dequeueCat(self):
        return self.cats.pop()

    def dequeueAny(self):
        if self.dogs.length > self.cats.length:
            return self.dogs.pop()
        else:
            return self.cats.pop()


if __name__ == "__main__":
    p = Pets()
    p.enqueue("dog", "black")
    p.enqueue("cat", "little")
    p.enqueue("dog", "gray")
    p.enqueue("cat", "big")
    p.enqueue("dog", "white")
    p.enqueue("cat", "fat")
    p.enqueue("dog", "brown")
    print(p)
    print(p.dequeueDog())
    print(p.dequeueAny())
    print(p.dequeueCat())
    print(p.dequeueAny())
    print(p.dequeueCat())
    p.enqueue("cat", "thin")
    print(p)
