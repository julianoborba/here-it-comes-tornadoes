from distutils.core import setup

setup(
    name='tornado-worker',
    version='v1.0.0',
    packages=['worker', 'worker.sqs', 'worker.notice'],
    entry_points={
        'console_scripts': ['worker=handler:handler']
    },
    url='',
    license='',
    author='',
    author_email='',
    description=''
)
