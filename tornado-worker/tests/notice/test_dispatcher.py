from unittest import TestCase
from worker.notice import dispatcher


class DispatcherTest(TestCase):

    def test_should_dispatch_a_notice(self):
        notice = {
            'subject': 'Screamming guy system',
            'message': 'Here it comes!',
            'channel': 'FOOBAR123',
            'receipt_handle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...'
        }
        expected_slack_notice = {
            'attachments': [
                {
                    'attachment_type': 'default',
                    'color': '#ad1721',
                    'text': 'Here it comes!'
                }
            ],
            'channel': 'FOOBAR123',
            'text': 'Screamming guy system',
            'username': 'Guy'
        }

        slack_notice = dispatcher.build_slack_notice(notice)

        self.assertEqual(expected_slack_notice, slack_notice)
