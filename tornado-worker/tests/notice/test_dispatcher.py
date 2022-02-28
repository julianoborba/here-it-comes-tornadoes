from unittest import TestCase
from worker.notice import dispatcher


class DispatcherTest(TestCase):

    def test_should_dispatch_a_notice(self):
        notices = [
            {
                'subject': 'Screamming guy system',
                'message': 'Here it comes!',
                'channel': 'FOOBAR123',
                'receipt_handle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...'
            },
            {
                'subject': 'Screamming guy system',
                'message': 'Here it comes!',
                'channel': 'FOOBAR123',
                'receipt_handle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...'
            }
        ]
        expected_result = True

        result = dispatcher.dispatch(notices)

        self.assertEqual(expected_result, result)
