package main

import (
	"fmt"
	"lesson8/envconf"
	"lesson8/fileConfig"
	"os"
)

//К приложению из практической части предыдущего урока добавьте возможность читать данные из файлов.
//Конфигурация может быть задана в форматах yaml или json. Также по желанию вы можете добавить и другие форматы.
//Помимо чтения конфигурации приложение также должно валидировать её - например, все URL’ы должны соответствовать ожидаемым форматам.
//Работу с конфигурацией необходимо вынести в отдельный пакет (не в пакет main).
//Предложить тему для консультации

//Темы
//Дженерики, так как их нет в нашем курсе
//Что выносить в env а что в файл?
//Marshal сложной структуры json

//Как правильно заполнить структуру?
//type Bimbo struct {
//		Messages []struct {
//			Attachment struct {
//				Type    string `json:"type"`
//				Payload struct {
//					TemplateType string `json:"template_type"`
//					Text         string `json:"text"`
//					Buttons      []struct {
//						Type                string `json:"type"`
//						URL                 string `json:"url"`
//						Title               string `json:"title"`
//						WebviewHeightRatio  string `json:"webview_height_ratio"`
//						MessengerExtensions bool   `json:"messenger_extensions"`
//					} `json:"buttons"`
//				} `json:"payload"`
//			} `json:"attachment"`
//		} `json:"messages"`
//	}

const confName = "config.yaml"

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

	err = os.Setenv("SENTRY_URL", "http://sentry:9000")
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
	envConf, err := envconf.New()
	if err != nil {
		fmt.Println("Ошибка получения конфигурации", err)
		return
	}
	fmt.Printf("Конфигурация env %+v\n", *envConf)

	//Получаем конфигурацию из файла
	fileConf, err := fileConfig.New(confName)
	if err != nil {
		fmt.Println("Ошибка получения конфигурации", err)
		return
	}

	fmt.Printf("Конфигурация из файла %+v\n", *fileConf)
}
