FROM python:3.9-slim AS build-env
WORKDIR /app

RUN adduser worker
USER worker
WORKDIR /home/worker

COPY --chown=worker:worker requirements.txt .
RUN pip3 install --user -r requirements.txt

COPY --chown=worker:worker worker/ worker/
COPY --chown=worker:worker setup.py .
RUN pip3 install --user .

CMD [ "python", "-m", "worker.main" ]
