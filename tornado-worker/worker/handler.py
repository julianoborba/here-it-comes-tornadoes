from worker.sqs import consumer
from worker.notice import dispatcher


def run():
    messages = consumer.receive_message()
    notices = consumer.parse_message(messages)
    dispatcher.dispatch(notices)
    receipt_handles = [notice['receipt_handle'] for notice in notices]
    consumer.delete_received_message(receipt_handles)
