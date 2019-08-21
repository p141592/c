# local_settings
Настройки командной оболочки MacOS / Linux

## Установка настроек:

* Сделает список приложений с версиями, которые были установлены
* Сохранит ваши текущие настройки в папку .local_settings/backup/
* Cкачает настройки в домашнюю дирректорию вашего пользователя 
* Cоздать ссылку на скрипт local_settings.sh в /usr/local/bin/, через который можно управлять настройками
* Выполнит установку всех необходимых приложений

`cd ~ && git init && git remote add origin git@github.com:p141592/local_settings.git && git pull origin master && ln -s local_settings.sh /usr/local/bin/local_settings.sh && local_settings.sh init`

## Загрузить новые настройки:

* Выполнит git pull актуальных настроек
* Запустит команду установки

`local_settings.sh upgrade` 


## Удалить настройки

* Удалит приложения
* Удалит логи
* Удалит настройки
* Вернет файлы настроек
* Установит приложения, которые были раньше

`local_settings.sh remove`
