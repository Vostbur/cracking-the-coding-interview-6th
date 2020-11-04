# Бинарное дерево поиска было создано обходом массива слева направо и вставкой каждого элемента.
# Для заданного бинарного дерева поиска с разными элементами выведите все возможные массивы,
# которые могли привести к созданию этого дерева.
# Пример:
# Ввод: 2
#      1 3
# Вывод: {2, 1, 3}, {2, 3, 1}
from chapter_4.binary_tree import BinaryTree
from chapter_4.linked_list import LinkedList


def all_sequences(node: BinaryTree) -> list:
    # получаем связанный список узлов левого поддерева
    left_seq = LinkedList()
    in_order_traversal(node.left, left_seq)
    # получаем связанный список узлов правого поддерева
    right_seq = LinkedList()
    in_order_traversal(node.right, right_seq)
    # переплетаем два списка с сохранением порядка следования элементов
    weaved = []
    weave_lists(left_seq, right_seq, weaved, LinkedList())
    # в результат каждого переплетения добавляем корневой элемент
    for i in weaved:
        i.insert_at_start(node.value)
    return weaved


def in_order_traversal(n: BinaryTree, l: LinkedList) -> None:
    if n:
        in_order_traversal(n.left, l)
        l.insert_at_end(n.value)
        in_order_traversal(n.right, l)


# Алгоритм работы для переплетения {1, 2, 3} с {4, 5, 6}
# weave(first, second, prefix):
# 	weave({1, 2}, {3, 4}, {})
# 	 weave({2}, {3, 4}, {1})
#    weave({}, {3, 4}, {1, 2})
#     {1, 2, 3, 4}
#    weave({2}, {4}, {1, 3})
#     weave({}, {4}, {1, 3, 2})
#  	{1, 3, 2, 4}
#     weave({2}, {}, {1, 3, 4})
#      {1, 3, 4, 2}
#  weave({1, 2}, {4}, {3})
#   weave({2}, {4}, {3, 1})
#    weave({}, {4}, {3, 1, 2})
#     {3, 1, 2, 4}
#    weave({2}, {}, {3, 1, 4})
#     {3, 1, 4, 2}
#   weave({1, 2}, {}, {3, 4})
#    {3, 4, 1, 2}
def weave_lists(first: LinkedList, second: LinkedList,
                results: list, prefix: LinkedList) -> None:
    if not (len(first) * len(second)):
        result = LinkedList()
        # print('prefix', prefix)
        for i in prefix:
            result.insert_at_end(i)
        # print('first', first)
        for i in first:
            result.insert_at_end(i)
        # print('second', second)
        for i in second:
            result.insert_at_end(i)
        results.append(result)
        return

    head_first = first.start_node.item
    first.delete_at_start()
    prefix.insert_at_end(head_first)
    weave_lists(first, second, results, prefix)
    prefix.delete_at_end()
    first.insert_at_start(head_first)

    head_second = second.start_node.item
    second.delete_at_start()
    prefix.insert_at_end(head_second)
    weave_lists(first, second, results, prefix)
    prefix.delete_at_end()
    second.insert_at_start(head_second)


if __name__ == '__main__':
    init = [1, 2, 4, 5]
    tree = BinaryTree(3)
    for i in init:
        tree.insert(i)
    res = all_sequences(tree)
    print(tree)
    for i in res:
        print(i)

