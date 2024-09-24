# Auth service

Для запуска необходимо иметь установленный **minikube**, **kubectl**, **helm**. Также для удобства используйте утилиту **make**

Порядок запуска:

1. Склонировать репозиторий
```shell
    git clone https://github.com/POMBNK/otusauth && cd otusauth
``` 
2. Добавить в файл hosts:
    - 127.0.0.1 arch.homework
   
затем для старта приложения выполнить команду ```make start``` 


3. Запустить http тесты. Убедитесь, что установлена утилита newman. 
В противном случае импортируйте коллекцию в Postman и запустите ее вручную
```shell
make test
```

4. По окончанию работы, очистить helm list + джобы
```shell
make stop
```

