4.7.
====

Имеется список проектов и список зависимостей (список пар проектов, для
которых первый проект зависит от второго проекта). Проект может быть
построен только после построения всех его зависимостей. Найдите такой
порядок построения, который позволит построить все проекты. Если действительного порядка не существует, верните признак ошибки.
Пример:
Ввод:
 проекты: a, b, c, d, e, f
 зависимости: (d, a), (b, f), (d, b), (a, f), (c, d)
Вывод:
 f, e, a, b, d, c
