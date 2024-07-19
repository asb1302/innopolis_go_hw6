# hw6

### Задание дословно:

* Начинающий программист пытался поиграть с многопоточность и хотел сделать функцию, которая записывает в мапу цену и
  обновляет эту цену
* Разработчик хотел сделать обработчик, который будет изменять мапу с ценами
* Впоследствии был план, что мы запускаем несколько обработчиков и у нас есть один writer, который нужно сначала
  починить
* Когда починим writer, можем приступить ко второй части: есть несколько багов, связанных с конкурентным доступом и
  правильной расстановкой необходимых инструментов синхронизации.
* Т.е. если мы это запустим, то сейчас версия будет багованная
* Задача починить - версия не должна быть багованной
* Например, сейчас результат отличается от ожидаемого: мы почему-то получаем одни и те же значения, хотя есть тот же
  инкремент (31 строка), а в последствии есть еще один обработчик, который то же добавляет к каждому по единице
* Если сейчас всё раскомментируем, то вообще упадём


---

### Вывод до правок RunWriter:

```
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
```


### Вывод после правок RunWriter:

```
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
```

### Вывод после удаления комментариев:

```
map[inst1:2.1 inst2:3.1 inst3:4.1 inst4:5.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:3.1 inst2:4.1 inst3:5.1 inst4:6.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:6.1 inst2:7.1 inst3:8.1 inst4:9.1]
map[inst1:4.1 inst2:5.1 inst3:6.1 inst4:7.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:6.1 inst2:7.1 inst3:8.1 inst4:9.1]
map[inst1:7.1 inst2:8.1 inst3:9.1 inst4:10.1]
map[inst1:5.1 inst2:6.1 inst3:7.1 inst4:8.1]
map[inst1:6.1 inst2:7.1 inst3:8.1 inst4:9.1]
map[inst1:7.1 inst2:8.1 inst3:9.1 inst4:10.1]
map[inst1:8.1 inst2:9.1 inst3:10.1 inst4:11.1]
```
