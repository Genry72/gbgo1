package main

import (
	"fmt"
	"lesson8/envconf"
	"os"
)

//Разработайте пакет для чтения конфигурации типичного веб-приложения через флаги или переменные окружения.
//Пример конфигурации можно посмотреть здесь. По желанию вы можете задать другие имена полям, сгруппировать их или добавить собственные поля.
//Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
//Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).

func main() {
	//Для демонстрации пишем все значения из примера в переменные окружения.
	os.Clearenv()
	err := os.Setenv("PORT", "8082")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Setenv("DB_URL", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Setenv("JAEGER_URL", "http://jaeger:16686")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Здесь будет ошибка при валидации
	err = os.Setenv("SENTRY_URL", "http:/sentry:9000")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Setenv("KAFKA_BROKER", "kafka:9092")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Setenv("SOME_APP_ID", "testid")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.Setenv("SOME_APP_KEY", "testkey")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Получаем конфигурацию из переменных окружение
	conf, err := envconf.New()
	if err != nil {
		fmt.Println("Ошибка получения конфигурации", err)
		return
	}
	fmt.Printf("%+v\n", conf)
}
