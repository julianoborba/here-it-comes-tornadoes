from queue_consumer import consumer
from message_dispatcher import dispatcher


def run():
    message = consumer.consume()
    dispatcher.dispatch(message)


if __name__ == '__main__':
    run()
