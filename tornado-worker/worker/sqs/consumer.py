from os import getenv
from boto3 import client

QUEUE_URL = getenv('QUEUE_URL')
QUEUE_REGION = getenv('QUEUE_REGION')

SQS_CLIENT = client(
    'sqs',
    region_name=QUEUE_REGION,
    endpoint_url=QUEUE_URL
)


def receive_message(max_number_of_messages=1):
    return SQS_CLIENT.receive_message(
        QueueUrl=QUEUE_URL,
        AttributeNames=[
            'SentTimestamp'
        ],
        MaxNumberOfMessages=max_number_of_messages,
        MessageAttributeNames=[
            'All'
        ],
        VisibilityTimeout=0,
        WaitTimeSeconds=0
    )


def delete_received_message(receipt_handles):
    for receipt_handle in receipt_handles:
        SQS_CLIENT.delete_message(
            QueueUrl=QUEUE_URL,
            ReceiptHandle=receipt_handle
        )


def parse_message(messages):
    if not messages.get('Messages'):
        print('no messages to parse')
        return []

    notices = []
    for message in messages['Messages']:
        notices.append({
            'subject': message['Body'],
            'finding': message['MessageAttributes']['Finding']['StringValue'],
            'channel': message['MessageAttributes']['Channel']['StringValue'],
            'receipt_handle': message['ReceiptHandle']
        })
    print(
        f'{len(messages["Messages"])} message(s) were '
        f'parsed into {len(notices)} notice(s)'
    )

    return notices
