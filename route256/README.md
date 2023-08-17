# Route 256: Junior
Песочница (Go)

#### Оповещение

Для решений на Go считывать данные рекомендуется, используя bufio.NewReader(os.Stdin). 

Иначе при больших входных данных могут быть проблемы с каким-то внутренним буфером в Go и данные не будут вычитаны полностью. 

Например, зачастую можно делать вот так: in := bufio.NewReader(os.Stdin) и потом fmt.Fscan(in, &T).
````
in := bufio.NewReader(os.Stdin)
fmt.Fscan(in, &T)
````
Аналогично с выводом. Следует использовать out := bufio.NewWriter(os.Stdout) (не забывайте про defer out.Flush()).
````
out := bufio.NewWriter(os.Stdout)
defer out.Flush()
````
А вот если делать просто fmt.Scan(&t) или fmt.Println(...), то при больших входных данных могут быть проблемы.