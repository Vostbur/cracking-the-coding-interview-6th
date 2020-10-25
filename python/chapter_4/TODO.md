4.4.
====

Реализуйте функцию, проверяющую сбалансированность бинарного дерева.
Будем считать дерево сбалансированным, если разница высот двух поддеревьев любого узла не превышает 1.

Подсказки:
----------

Задумайтесь над определением сбалансированного дерева. Возможно ли
проверить это условие для одного узла? Возможно ли проверить его для
всех узлов?

Если вы разработали решение методом «грубой силы», будьте осторожны
с его временем выполнения. Если вы вычисляете высоту поддеревьев для
каждого узла, алгоритм может получиться весьма неэффективным

Нельзя ли изменить класс узла бинарного дерева, чтобы в нем хранилась
высота его поддерева?

Для хранения высоты поддерева не нужно изменять класс бинарного дерева. Сможет ли ваша рекурсивная функция вычислить высоту каждого поддерева одновременно с проверкой сбалансированности узла? Попробуйте
написать функцию, которая возвращает сразу несколько значений

Возможно написать одну функцию checkHeight, которая совмещает вычисление высоты с проверкой сбалансированности. Целочисленное возвращаемое значение может обозначать оба условия

4.5.
====

Реализуйте функцию для проверки того, является ли бинарное дерево бинарным деревом поиска.

Подсказки:
----------

Если при симметричном обходе дерева элементы следовали в правильном
порядке, означает ли это, что дерево действительно правильно упорядочено? Что произойдет при наличии дубликатов? Если дубликаты разрешены,
они должны находиться с конкретной стороны (обычно слева)

Чтобы дерево было бинарным деревом поиска, недостаточно выполнения
условия left.value <= current.value < right.value для каждого
узла. Любой узел в левом поддереве должен быть меньше текущего, а последний должен быть меньше любого узла в правом поддереве

Если каждый узел в левом поддереве должен быть меньше текущего узла
либо равен ему, это фактически означает, что самый большой узел в левом
поддереве должен быть меньше текущего узла либо равен ему

Вместо того чтобы проверять значение текущего узла по значениям leftTree.max и rightTree.min, нельзя ли действовать наоборот — проверить
узлы левого дерева и убедиться в том, что они меньше current.value?

Рассмотрите checkBST как рекурсивную функцию, которая проверяет, что
каждый узел находится в пределах допустимого диапазона (минимум,
максимум). Сначала этот диапазон бесконечен. При обходе левого поддерева минимумом становится отрицательная бесконечность, а максимумом —
root.value. Сможете ли вы реализовать эту рекурсивную функцию и обеспечить необходимую корректировку диапазонов в процессе обхода дерева?

4.6.
====

Напишите алгоритм поиска «следующего» узла для заданного узла в бинарном
дереве поиска. Считайте, что у каждого узла есть ссылка на его родителя.

Подсказки:
----------

Подумайте над тем, как работает симметричный обход, и попробуйте «развернуть» его в обратном направлении

Один из шагов логики: «следующим» является крайний левый узел правого
поддерева. А если правого поддерева нет?

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

Подсказки:
----------

Постройте направленный граф, представляющий зависимости. Каждый узел
представляет проект, и ребро из A в B существует в том случае, если B зависит от A (проект A должен быть построен до B). Также можно построить
дерево наоборот, если вам так будет удобнее

Взгляните на граф. Сможете ли вы найти на нем узел, который можно заведомо безопасно построить первым?

Если вам удастся найти узел, не имеющий входящих ребер, он определенно может быть построен. Найдите такой узел (их может быть несколько)
и включите его в порядок построения. Что это будет означать для его исходящих ребер?

Когда вы решаете построить узел, его исходящее ребро может быть удалено. Как после этого найти другие узлы, которые можно построить?

Попробуйте применить совершенно иной подход: поиск в глубину от произвольного узла. Как связаны между собой поиск в глубину и допустимый порядок построения проектов?

Выберите произвольный узел и проведите от него поиск в глубину. Добравшись до конца пути, мы знаем, что этот узел может быть последним, потому
что никакие другие узлы от него не зависят. Что это значит для узлов, непосредственно предшествующих ему?

4.8.
====

Создайте алгоритм и напишите код поиска первого общего предка двух узлов
бинарного дерева. Постарайтесь избежать хранения дополнительных узлов
в структуре данных. Примечание: бинарное дерево не обязательно является
бинарным деревом поиска.

Подсказки:
----------

Если бы в каждом узле хранилась ссылка на его родителя, то можно было
бы воспользоваться решением к задаче 2.7. Впрочем, интервьюер может отказать вам в этом предположении

Первым общим предком является узел с наибольшей глубиной, потомками
которого являются и p, и q. Подумайте над тем, как идентифицировать этот
узел

Как определить, является ли p потомком узла n?

Начните с корня. Можете ли вы определить, является ли корень первым
общим предком? Если не является, то можете ли вы определить, с какой
стороны корня находится первый общий предок?

Примените рекурсивный подход. Проверьте, являются ли p и q потомками
левого и правого поддерева. Если они являются потомками разных поддеревьев, то текущий узел является первым общим предком. Если они являются
потомками одного поддерева, то это поддерево содержит первого общего
предка. Как эффективно реализовать эту схему?

В «наивном» алгоритме один метод проверял, является ли x потомком n,
а другой метод осуществлял рекурсию для поиска первого общего предка.
Такой подход приводит к повторению поиска по одним элементам поддерева. Поиск следует объединить в одну функцию firstCommonAncestor.
Какие возвращаемые значения предоставят необходимую информацию?

Функция firstCommonAncestor может вернуть первого общего предка
(если и p, и q находятся в дереве); p, если в дереве находится только p,
но не q; q, если в дереве находится только q, но не p; и null в остальных
случаях

Будьте внимательны! А ваш алгоритм обрабатывает случай, в котором существует только один узел? Что произойдет? Возможно, возвращаемые значения придется слегка изменить

4.9.
====

Бинарное дерево поиска было создано обходом массива слева направо и вставкой каждого элемента. Для заданного бинарного дерева поиска с разными
элементами выведите все возможные массивы, которые могли привести
к созданию этого дерева.
Пример:
Ввод:
2
1 3
Вывод: {2, 1, 3}, {2, 3, 1}

Подсказки:
----------

Каким должно быть самое первое значение, присутствующее в каждом массиве?

Корень — самое первое значение, которое должно присутствовать в каждом
массиве. Что можно сказать о порядке значений в левом поддереве по сравнению со значениями в правом поддереве? Должны ли значения в левом
поддереве вставляться до правого поддерева?

Отношения между значениями левого и правого поддерева фактически произвольны. Значения левого поддерева могут быть вставлены до значений
правого поддерева; возможно и обратное (правые значения предшествуют
левым), как и любой другой порядок

Разбейте задачу на подзадачи. Используйте рекурсию. Если известны все
возможные последовательности для левого и правого поддеревьев, как создать все возможные последовательности для всего дерева?

4.10.
=====

T1 и T2 — два очень больших бинарных дерева, причем T1 значительно больше
T2. Создайте алгоритм, проверяющий, является ли T2 поддеревом T1.
Дерево T2 считается поддеревом T1, если существует такой узел n в T1, что
поддерево, «растущее» из n, идентично дереву T2. (Иначе говоря, если вырезать дерево в узле n, оно будет идентично T2.)

Подсказки:
----------

Если T2 является поддеревом T1, как его порядок симметричного обхода
связан с порядком симметричного обхода T1? А как насчет префиксного
и постфиксного обхода?

Порядок симметричного обхода не принесет особой пользы. В конце концов,
у всех бинарных деревьев поиска с одинаковыми значениями (независимо
от структуры) этот порядок одинаков. (А если этот метод не работает в конкретном случае бинарного дерева поиска, то он не сработает и для более
общего бинарного дерева.) С другой стороны, из порядка префиксного обхода можно извлечь намного больше полезной информации

Возможно, вы решили, что если T2.preorderTraversal() является подстрокой T1.preorderTraversal(), то T2 является поддеревом T1. Это
почти правда… если деревья не содержат дубликатов. Допустим, T1 и T2
состоят только из дубликатов, но имеют разную структуру. Порядок префиксного обхода в этом случае будет совпадать, хотя T2 и не является поддеревом T1. Как справиться с подобными ситуациями?

Хотя на первый взгляд может показаться, что проблема возникает из-за
дубликатов, на самом деле она глубже. Порядки префиксного обхода совпадают только из-за наличия пропущенных null-узлов. Рассмотрите возможность вставки условного значения в строку префиксного обхода для
каждого null-узла. Регистрация null-узла как «реального» узла позволит
различать разные структуры

Задачу также можно решать рекурсивно. Если задан конкретный узел в T1,
можно ли проверить, совпадает ли его поддерево с T2?

4.11.
=====

Вы пишете с нуля класс бинарного дерева поиска, который помимо методов
вставки, поиска и удаления содержит метод getRandomNode() для получения
случайного узла дерева. Вероятность выбора всех узлов должна быть одинаковой. Разработайте и реализуйте алгоритм getRandomNode; объясните, как
вы реализуете остальные методы.

Подсказки:
----------

Будьте очень внимательны. Следите за тем, чтобы все узлы выбирались
с равной вероятностью, а ваше решение не замедляло работу стандартных
алгоритмов бинарных деревьев поиска (insert, find и delete). Также
помните, что даже если дерево является сбалансированным бинарным деревом поиска, это не означает его полноты/законченности/идеальности

Вы сами создаете класс бинарного дерева поиска, поэтому можете включить
в него любую информацию о структуре дерева или узлах (если только это не
приведет к нежелательным последствиям — например, замедлению insert).
Вероятно, интервьюер именно поэтому указал, что класс пишется с нуля,
а для эффективной реализации операции потребуется сохранить дополнительную информацию

Можно ли воспользоваться алгоритмом обхода дерева для реализации этого
алгоритма (как наивное решение методом «грубой силы»)? За какое время
он будет выполняться?

Другое возможное решение — выбрать случайную глубину для обхода и выполнить случайный обход, остановившись при достижении выбранной глубины. Однако поразмыслите над этим решением. Будет ли оно работать?

Выбор случайной глубины особой пользы не принесет. Во-первых, на нижних уровнях больше узлов, чем на верхних. Во-вторых, даже если бы эти
вероятности удалось сбалансировать, возможен «тупик», когда необходимо
выбрать узел на глубине 5, а на глубине 3 обнаруживается лист. Впрочем,
идея балансировки вероятностей выглядит интересно

Многие кандидаты предлагают наивное решение: выбрать случайное число
от 1 до 3. Если выбрано число 1, то возвращается текущий узел; если 2 —
происходит переход к левому поддереву, а если 3 — происходит переход
к правому поддереву. Такое решение не работает. Почему? Можно ли исправить его так, чтобы оно заработало?

Предыдущее решение (с выбором случайного числа от 1 до 3) не работает
из-за того, что вероятности узлов не равны. Например, корень будет возвращаться с вероятностью 1/3, даже если дерево содержит более 50 узлов.
Очевидно, все узлы не могут иметь вероятность 1/3. Проблему можно решить выбором случайного числа от 1 до размера дерева. Впрочем, проблема
будет решена только для корня. А как насчет остальных узлов?

Проблема предыдущего решения заключается в том, что на одной стороне
поддерева может быть больше узлов, чем на другой. Следовательно, необходимо взвесить вероятности выбора левого или правого поддерева в зависимости от количества узлов с каждой из сторон. Как будет работать такое
решение? Как узнать количество узлов?

4.12.
=====

Дано бинарное дерево, в котором каждый узел содержит целое число (положительное или отрицательное). Разработайте алгоритм для подсчета всех
путей, сумма значений которых соответствует заданной величине. Обратите
внимание, что путь не обязан начинаться или заканчиваться в корневом или
листовом узле, но он должен идти вниз (переходя только от родительских
узлов к дочерним).

Подсказки:
----------

Попробуйте упростить задачу. Что если путь должен начинаться от корня?

Не забудьте, что пути могут перекрываться. Например, при поиске суммы 6
действителен как путь 1->3->2, так и путь 1->3->2->4->-6->2

Если каждый путь должен начинаться от корня, мы можем обойти все возможные пути, начиная от корня. Сумма отслеживается в процессе обхода,
а счетчик totalPaths увеличивается при каждом нахождении пути с целевой суммой. Как распространить ситуацию на пути, который может начинаться с произвольного узла? Не забудьте: для начала стоит создать решение, работающее методом «грубой силы». Оптимизацией можно заняться
позднее

Чтобы обобщить решение для путей с произвольным началом, достаточно
повторить процесс для всех узлов

Если вы спроектировали алгоритм так, как описано выше, он выполняется
за время O(N log N) в сбалансированном дереве. Это объясняется тем, что
он содержит N узлов, каждый из которых в худшем случае находится на
глубине O(log N). Узел обрабатывается один раз для каждого узла, расположенного выше него. Следовательно, N узлов будут обработаны O(log N) раз.
Существует оптимизация, которая обеспечит время выполнения O(N)

Какая работа дублируется в текущем алгоритме «грубой силы»?

Рассмотрите каждый путь, начинающийся от корня (существуют N таких путей), как массив. По сути наш алгоритм «грубой силы» берет каждый массив
и ищет в нем все смежные подпоследовательности, имеющие заданную сумму. Для этого мы вычисляем все подмассивы и их суммы. Возможно, стоит
сосредоточиться на этой подзадаче. Как для заданного массива найти все
смежные подпоследовательности с конкретной суммой? Еще раз поразмыслите над повторением работы в алгоритме «грубой силы»

Мы ищем подмассивы с суммой targetSum. Заметим, что значение runningSumi
 (сумма элементов от 0 до i) может быть получено за постоянное
время. Чтобы подмассив элементов от i до j имел сумму targetSum, значение runningSumi-1 + targetSum должно быть равно runningSumj
 (попробуйте нарисовать элементы массива или цепочку чисел). Так как мы можем
отслеживать runningSum «на ходу», как быстро найти количество индексов
i, для которых выполняется это уравнение?

Попробуйте использовать хеш-таблицу, связывающую значение runningSum
с количеством элементов, обладающих этим значением runningSum

После того как вы проработаете алгоритм для поиска в массиве всех смежных подмассивов с заданной суммой, попробуйте применить его к дереву.
Помните, что в процессе обхода и модификации хеш-таблицы может возникнуть необходимость в «восстановлении» ее содержимого при возврате