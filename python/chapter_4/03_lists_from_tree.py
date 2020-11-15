# Для бинарного дерева разработайте алгоритм, который создает связный
# список всех узлов, находящихся на каждой глубине (для дерева с глубиной
# D должно получиться D связных списков).
from chapter_4.binary_tree import BinaryTree
from chapter_4.linked_list import LinkedList


# depth-first search case
def createListsDFS(tree):
    lists = []
    createLevelLinkedListDFS(tree, lists, 0)
    return lists


# Модифицированный вариант поиска в глубину
def createLevelLinkedListDFS(tree, lists, level):
    if not tree:
        return
    if len(lists) == level:
        # Если это новый уровень, создаем новый связанный список для хранения
        # узлов уровня и добавляем в него первый узел
        llist = LinkedList()
        llist.insert_at_end(tree.value)
        lists.append(llist)
    else:
        # Список этого уровня уже существует, добавляем в массив узлов новый
        # узел уровня
        lists[level].insert_at_end(tree.value)
    createLevelLinkedListDFS(tree.left, lists, level+1)
    createLevelLinkedListDFS(tree.right, lists, level+1)


# Модифицированный вариант поиска в ширину.
# чтобы узнать, какие узлы находятся на уровне i, можно просто
# посмотреть на дочерние узлы уровня i–1.
def createListsBFS(tree):
    result = []
    current = LinkedList()
    # добавляем корень
    if tree:
        current.insert_at_end(tree)
    while len(current) > 0:
        # добавляем уровень i-1
        result.append(current)
        parents = current
        # переходим на уровень i
        current = LinkedList()
        for parent in parents:
            # для уровня i сохраняем потомков узла i-1
            if parent.item.left:
                current.insert_at_end(parent.item.left)
            if parent.item.right:
                current.insert_at_end(parent.item.right)
    return result


if __name__ == '__main__':
    init = [-20, -50, -15, -60, 50, 60, 55, 85, 15, 5, -10]
    tree = BinaryTree(50)
    for i in init:
        tree.insert(i)
    print(tree)

    print('Result with modified DFS algorithm:')
    lists = createListsDFS(tree)
    for i in lists:
        print(i)

    print('Result with modified BFS algorithm:')
    lists = createListsBFS(tree)
    for i in lists:
        print([j.item.value for j in i])
