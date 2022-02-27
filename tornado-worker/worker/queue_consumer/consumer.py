from os import getenv
from boto3 import client
from botocore import UNSIGNED
from botocore.config import Config

QUEUE_URL = getenv('QUEUE_URL', 'http://localhost:4566/000000000000/notices')

SQS_CLIENT = client(
    'sqs',
    config=Config(signature_version=UNSIGNED),
    endpoint_url=QUEUE_URL
)


def receive_one_message():
    return SQS_CLIENT.receive_message(
        QueueUrl=QUEUE_URL,
        AttributeNames=[
            'SentTimestamp'
        ],
        MaxNumberOfMessages=1,
        MessageAttributeNames=[
            'All'
        ],
        VisibilityTimeout=0,
        WaitTimeSeconds=0
    )


def delete_received_message(message_receipt_handle):
    if not message_receipt_handle:
        return

    SQS_CLIENT.delete_message(
        QueueUrl=QUEUE_URL,
        ReceiptHandle=message_receipt_handle
    )


def parse_first_message(queue_message):
    if not queue_message.get('Messages'):
        return None, None

    message = queue_message['Messages'][0]

    receipt_handle = message['ReceiptHandle']

    message = {
        'subject': message['Body'],
        'message': message['MessageAttributes']['Message']['StringValue'],
        'channel': message['MessageAttributes']['Channel']['StringValue']
    }

    return message, receipt_handle
