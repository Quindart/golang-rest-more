# Shared RabbitMQ

RabbitMQ is infrastructure shared by `blog-service` and `sale-service`; neither
application compose file starts its own broker.

Start the broker once from this directory:

```sh
docker compose -f docker-compose.messaging.yaml up -d
```

For services running on the host, configure both environments with:

```dotenv
RABBITMQ_URL=amqp://app:app_password@localhost:5672/
```

The broker exposes AMQP on port 5672 and the management UI on port 15672.
The durable queue `blog-sale.queue` carries persistent `user.created` events.
In `blog-service`, the user application service writes the user and outbox row
in one database transaction. An application outbox worker passes pending events
to the RabbitMQ publisher adapter. In `sale-service`, the RabbitMQ adapter
decodes the delivery and calls the application handler; that handler calls
`CustomerRepository.Create`. The repository stores the event ID and customer in
one database transaction, then the adapter acknowledges the message. This makes
redelivery safe by event ID and unique user ID.
