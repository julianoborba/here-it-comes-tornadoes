<p align="center">
  <img alt="tornadoes-are-cool" src="https://user-images.githubusercontent.com/6361839/155834015-9f1beb9b-1076-431b-9592-f34ba8e955ca.png" height="200" />
</p>

# here-it-comes-tornadoes

A basic system to enqueue sensors findings and forward to some Slack channel, hopefully.

### Don't mind if i put random annotations. M'kay?

```
https://github.com/localstack/localstack
localstack start -d


awslocal sqs create-queue --queue-name tornados
- http://localhost:4566/000000000000/tornados


go run tornado-route.go
- localhost:8080/health-check
- localhost:8080/notice


awslocal sqs receive-message --queue-url http://localhost:4566/000000000000/tornados --attribute-names All --message-attribute-names All --max-number-of-messages 10

```
