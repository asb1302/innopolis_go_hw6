# hw6

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