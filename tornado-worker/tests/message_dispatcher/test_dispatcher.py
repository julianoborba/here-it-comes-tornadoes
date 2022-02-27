from unittest import TestCase
from worker.message_dispatcher import dispatcher


class DispatcherTest(TestCase):

    def test_should_dispatch(self):
        parsed_message = {
            'subject': 'Screamming guy system',
            'message': 'Here it comes!',
            'channel': 'FOOBAR123'
        }
        expected_result = True

        result = dispatcher.dispatch(parsed_message)

        self.assertEqual(expected_result, result)
