# Реализуйте алгоритм, удаляющий узел из середины односвязного списка
# (то есть узла, не являющегося ни начальным, ни конечным — не обязательно
# находящегося точно в середине).
# Доступ предоставляется только к этому узлу.
# Пример:
# Ввод: узел c из списка a->b->c->d->e->f
# Вывод: ничего не возвращается, но новый список имеет вид: a->b->d->e->f
import unittest
from LinkedList import LinkedList


def delete_middle(node):
    node.item = node.ref.item
    node.ref = node.ref.ref


class Test(unittest.TestCase):
    def test_delete_middle(self):
        ll = LinkedList()
        ll.insert_multiple(1, 2, 3)
        middle = ll.insert_at_end(4)
        ll.insert_multiple(5, 6, 7)
        delete_middle(middle)
        self.assertEqual(str(ll), "1,2,3,5,6,7")


if __name__ == "__main__":
    unittest.main()
