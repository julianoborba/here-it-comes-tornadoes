from worker.sqs import consumer
from worker.notice import dispatcher


def run():
    messages = consumer.receive_message()
    notices = consumer.parse_message(messages)
    dispatcher.dispatch(notices)

