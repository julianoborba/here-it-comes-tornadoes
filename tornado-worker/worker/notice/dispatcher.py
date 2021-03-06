from worker.sqs import consumer
from os import getenv
from requests import post

SLACK_WEBHOOK_URL = getenv('SLACK_WEBHOOK_URL')


def build_slack_notice(notice):
    data = {
        'username': 'Guy',
        'channel': notice['channel'],
        'text': notice['subject'],
        'attachments': [
            {
                'text': notice['finding'],
                'attachment_type': 'default',
                'color': '#ad1721',
            }
        ]
    }
    return data


def dispatch(notices):
    counter = 0
    for notice in notices:
        post(
            SLACK_WEBHOOK_URL,
            json=build_slack_notice(notice)
        ).raise_for_status()

        consumer.delete_received_message(
            [notice['receipt_handle'] for notice in notices]
        )

        counter += 1

    print(
        f'{counter} notices were dispatched '
        f'out of {len(notices)} notices in total'
    )
