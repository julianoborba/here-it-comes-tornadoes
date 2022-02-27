from worker.queue_consumer import consumer
from worker.message_dispatcher import dispatcher


def run():
    queue_message = consumer.receive_one_message()
    parsed_message, receipt_handle = consumer.parse_first_message(queue_message)
    consumer.delete_received_message(receipt_handle)
    dispatcher.dispatch(parsed_message)


if __name__ == '__main__':
    run()
