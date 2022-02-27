from unittest import TestCase
from worker.message_dispatcher import dispatcher


class DispatcherTest(TestCase):

    def test_should_dispatch(self):
        dispatcher.dispatch('foobar')
