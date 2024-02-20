# 03-url-collection

## Создание списка url

Напишите программу, которая будет хранить ваши url. На основании созданного шаблона допишите код, который позволяет 
добавлять новые ссылки, удалять и выводить список.

Для решения задачи используйте структуры. Обязательные поля структуры должны быть дата добавления, имя ссылки, 
теги для ссылки через запятую и сам url.

Например

```go
type Item struct {
	Name string
	Date time.Time
	Tags string
	Link string
}
```

Для хранение даты вы можете использовать пакет time. Чтобы создать дату можно использовать следующую функцию time.Now()
Чтобы вывести дату как строку используйте функцию Format, например
```go
t := time.Now() // создание текущей даты
t.Format(time.DateTime) // как строка
```
Наш проект использует библиотеку github.com/eiannone/keyboard для отслеживания нажатия клавиш, более подробно про 
подключение пакетов мы поговорим на отдельной лекции. В шаблоне вам нужно дописать код действий под соотвествующими 
нажатиями клавиш.

```shell
go build -o urlcollection
./urlcollection
```

** Задача со звездочкой. Реализовать хранение списка url в файле, подгрузка списка урлов из файла и удаление соответственно. Можно передалать в этом плане формат и вместо отслеживания нажатий клавиатуры сделать через классические аругменты.

** Задача с двумя звездочками, реализовать полностекстовый поиск по имени ссылки и по тегам. Можно реализовать сначала только по тегам так как это проще и после этого попробовать взяться за полнотекстовый поиск. Желательно посмотреть подходы в интернете как можно подобное реализовать.

## Решение задачи

Сначала задачу решил в соответствие с приведенным шаблоном. Но потом сделал рефакторинг для добавления возможности хранения в файле, а также исключил пакет github.com/eiannone/keyboard (субъективно мне не понравилось его поведение под Windows).  
Также добавил возможность полнотекстового поиска по всем полям всех объектов в хранилище (но тут в очень примитивном виде, ElasticSearch для данной задачи точно избыточный). В качестве score для ранжирования результатов использую счетчик наличия искомой подстроки в каком либо поле умноженный на коэффициент для поля (коэффициент использовал исходя из собственных предпочтений, т.к. в ТЗ не было указано).  
Также добавил тесты и Makefile.