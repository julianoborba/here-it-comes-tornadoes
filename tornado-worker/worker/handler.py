from worker.sqs import consumer
from worker.notice import dispatcher


def handler(event=None, context=None):
    messages = consumer.receive_message()
    notices = consumer.parse_message(messages)
    dispatcher.dispatch(notices)
