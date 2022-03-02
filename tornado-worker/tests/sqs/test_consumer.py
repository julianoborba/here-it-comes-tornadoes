from unittest import TestCase
from worker.sqs import consumer


class ConsumerTest(TestCase):

    def test_should_parse_queue_message_into_notice(self):
        messages = {
            'Messages': [
                {
                    'MessageId': '77062e28-40e8-502f-9427-fb7f3abfddcc',
                    'ReceiptHandle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
                    'MD5OfBody': 'f27eca4f499f59e0328f3f4ae35a4a1b',
                    'Body': 'Screamming guy system detect a tornado',
                    'Attributes': {
                        'SenderId': 'AIDAIT2UOQQY3AUEKVGXU',
                        'SentTimestamp': '1645902870467',
                        'ApproximateReceiveCount': '2',
                        'ApproximateFirstReceiveTimestamp': '1645902888338'
                    },
                    'MD5OfMessageAttributes': 'e64461b4cb51a781f7d3541436...',
                    'MessageAttributes': {
                        'Channel': {
                            'StringValue': 'FOOBAR123',
                            'DataType': 'String'
                        },
                        'Finding': {
                            'StringValue': 'Here it comes!',
                            'DataType': 'String'
                        }
                    }
                },
                {
                    'MessageId': '77062e28-40e8-502f-9427-fb7f3abfddcc',
                    'ReceiptHandle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
                    'MD5OfBody': 'f27eca4f499f59e0328f3f4ae35a4a1b',
                    'Body': 'Screamming guy system detect a tornado',
                    'Attributes': {
                        'SenderId': 'AIDAIT2UOQQY3AUEKVGXU',
                        'SentTimestamp': '1645902870467',
                        'ApproximateReceiveCount': '2',
                        'ApproximateFirstReceiveTimestamp': '1645902888338'
                    },
                    'MD5OfMessageAttributes': 'e64461b4cb51a781f7d3541436...',
                    'MessageAttributes': {
                        'Channel': {
                            'StringValue': 'FOOBAR123',
                            'DataType': 'String'
                        },
                        'Finding': {
                            'StringValue': 'Here it comes!',
                            'DataType': 'String'
                        }
                    }
                }
            ]
        }
        expected_notices = [
            {
                'subject': 'Screamming guy system detect a tornado',
                'finding': 'Here it comes!',
                'channel': 'FOOBAR123',
                'receipt_handle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
            },
            {
                'subject': 'Screamming guy system detect a tornado',
                'finding': 'Here it comes!',
                'channel': 'FOOBAR123',
                'receipt_handle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
            }
        ]
        expected_receipt_handles = [
            'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
            'avptgxktxigerbawfbjlkiaenwzgignqdfj...'
        ]

        notices = consumer.parse_message(messages)
        receipt_handles = [notice['receipt_handle'] for notice in notices]

        self.assertEqual(expected_notices, notices)
        self.assertEqual(expected_receipt_handles, receipt_handles)
