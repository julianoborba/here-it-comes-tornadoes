from unittest import TestCase
from worker.queue_consumer import consumer


class ConsumerTest(TestCase):

    def test_should_parse_queue_message(self):
        messages = {
            'Messages': [
                {
                    'MessageId': '77062e28-40e8-502f-9427-fb7f3abfddcc',
                    'ReceiptHandle': 'avptgxktxigerbawfbjlkiaenwzgignqdfj...',
                    'MD5OfBody': 'f27eca4f499f59e0328f3f4ae35a4a1b',
                    'Body': 'Screamming guy system',
                    'Attributes': {
                        'SenderId': 'AIDAIT2UOQQY3AUEKVGXU',
                        'SentTimestamp': '1645902870467',
                        'ApproximateReceiveCount': '2',
                        'ApproximateFirstReceiveTimestamp': '1645902888338'
                    },
                    'MD5OfMessageAttributes': 'e64461b4cb51a781f7d35414369a7bfc',
                    'MessageAttributes': {
                        'Channel': {
                            'StringValue': 'FOOBAR123',
                            'DataType': 'String'
                        },
                        'Message': {
                            'StringValue': 'Here it comes!',
                            'DataType': 'String'
                        }
                    }
                }]
        }
        expected_parsed_message = {
            'subject': 'Screamming guy system',
            'message': 'Here it comes!',
            'channel': 'FOOBAR123'
        }
        expected_receipt_handle = 'avptgxktxigerbawfbjlkiaenwzgignqdfj...'

        parsed_message, receipt_handle = consumer.parse_first_message(messages)

        self.assertEqual(expected_parsed_message, parsed_message)
        self.assertEqual(expected_receipt_handle, receipt_handle)
