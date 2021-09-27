# accounting-of-jobs

#Задача:
Написать простое REST-приложение, позволяющее вести учет рабочих мест.
Поддерживаются, по крайней мере следующие операции: добавление, удаление и обновление рабочего места.
Рабочее место содержит следующие сведения:
- имя компьютера
- сетевой адрес
- имя текущего пользователя
Клиентская часть работает на windows. Запуск клиента без параметров автоматически добавляет/обновляет сведения о текущей системе. Групповые операции приветствуются.
Серверная часть хранит данные в JSON формате в БД PostgreSQL и работает на linux.
бонус1:
Клиент умеет работать на linux и windows системах.
бонус2:
Серверная часть поставляется deb-пакетом и регистрируется как systemd сервис.

#Реализация

Сервер находится в директории cmd/aoj/
Клиент находится в директории cmd/aoj-client/

Для запуска сервера нужно:
1. Инициализировать файл config.yaml. Предпочтительно, сдетать его копию с названием config.local.yaml
2. Создать БД postgresql. При запуске сервера, нужные таблицы сами появятся в базе

Клиент запускается с параметрами
При запуске без параметров создается новое рабочее место с параметрами текущей системы
-m --method используемый метод взаимодействия с сервером GET, POST(при запуске без параметров - по умолчанию), PUT, DELETE. Не зависит от регистра
--id id записи в бд. Обязателен для PUT и DELETE, опционален для GET
-u --username имя пользователя. Обязателен для PUT, для POST обязателен при указании -h иили --ip
-h --hostname имя компьютера. Обязателен для PUT, для POST обязателен при указании -u или --ip
--ip ip адрес компьютера. Обязателен для PUT, для POST обязателен при указании -h и -u