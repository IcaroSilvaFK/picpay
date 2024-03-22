package main

import (
	"log"

	"github.com/IcaroSilvaFK/picpay/infra/configs"
	httpclient "github.com/IcaroSilvaFK/picpay/pkg/http"
)

func main() {

	forever := make(chan int)

	AmqpConsumer()

	<-forever
}

func AmqpConsumer() {
	client := httpclient.NewHttpClient("https://run.mocky.io/v3")
	con := configs.GetRabbitMQChannel()

	del, err := con.Consume(
		"picpay",
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	go func() {

		for d := range del {
			var res struct {
				Message bool `json:"message"`
			}

			log.Println(string(d.Body))

			if err := client.Get("/54dc2cf1-3add-45b5-b5a9-6bf7e7f1f4a6", &res); err != nil {
				d.Ack(false)
			} else {
				d.Ack(true)
			}

		}

	}()

}
