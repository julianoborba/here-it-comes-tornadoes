from os import getenv
from requests import post

SLACK_WEBHOOK_URL = getenv('SLACK_WEBHOOK_URL', 'https://hooks.slack.com/<your>/<slack>/<webhook>')


def build_slack_notice(notice):
    data = {
        'text': notice['subject'],
        'username': 'Guy',
        'channel': notice['channel'],
        'attachments': [
            {
                'text': notice['message'],
                'attachment_type': 'default',
                'color': '#ad1721',
            }
        ]
    }
    return data


def dispatch(notices):
    for notice in notices:
        post(
            SLACK_WEBHOOK_URL,
            json=build_slack_notice(notice)
        )
