<p align="center">
  <img alt="tornadoes-are-cool" src="https://user-images.githubusercontent.com/6361839/155834015-9f1beb9b-1076-431b-9592-f34ba8e955ca.png" height="200" />
</p>

# here-it-comes-tornadoes

A basic system to enqueue sensors findings and forward to some Slack channel, hopefully.

### Local usage

First, build images for API and Worker aplications
```
make -C tornado-api docker_build
make -C tornado-worker docker_build
```

Run containers with compose
```
docker-compose up
```

Setup a local AWS SQS
```
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name notices
```

Now you can submit a dummy notice through producer container
```
curl --request POST \
  --url http://localhost:8080/notice \
  --header 'Content-Type: application/json' \
  --data '{"origin":"Screamming guy system","message":"Here it comes!","channel": "C05002EAE"}'
```

Consume the queue and dispatch to Slack through worker container
```
docker run -i --env-file ./tornado-worker/.env --net=host --rm "doofi/tornado-worker:latest"
```
