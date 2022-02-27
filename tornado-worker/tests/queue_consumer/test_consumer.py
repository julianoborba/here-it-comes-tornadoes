from unittest import TestCase
from worker.queue_consumer import consumer


class ConsumerTest(TestCase):

    def test_should_consume(self):
        consumer.consume()
