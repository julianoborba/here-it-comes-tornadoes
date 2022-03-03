FROM python:3.9-slim AS build-env
WORKDIR /app
COPY requirements.txt .
RUN pip3 install -r requirements.txt
COPY worker/ worker/
COPY setup.py .
RUN pip3 install .

FROM gcr.io/distroless/python3 AS runtime-env
COPY --from=build-env /usr/local/lib/python3.9/site-packages /app
WORKDIR /app
CMD [ "-m", "worker.main" ]
