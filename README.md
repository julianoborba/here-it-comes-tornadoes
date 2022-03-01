<p align="center">
  <img alt="tornadoes-are-cool" src="https://user-images.githubusercontent.com/6361839/155834015-9f1beb9b-1076-431b-9592-f34ba8e955ca.png" height="200" />
</p>

# here-it-comes-tornadoes

A basic system to enqueue sensors findings and forward to some Slack channel, hopefully.

### Usage

```
make -C tornado-api docker_build
make -C tornado-worker docker_build
docker-compose up

export AWS_ACCESS_KEY_ID=foo
export AWS_SECRET_ACCESS_KEY=bar
aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name notices

curl --request POST \
  --url http://localhost:8080/notice \
  --header 'Content-Type: application/json' \
  --data '{"origin":"Screamming guy system","message":"Here it comes!","channel": "C05002EAE"}'

docker run -i --net=host --rm "doofi/tornado-worker:latest"
```
