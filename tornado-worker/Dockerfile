FROM python:3.9-slim AS build-env
WORKDIR /app
COPY Pipfile /app
COPY Pipfile.lock /app
RUN pip install pipenv
RUN PIPENV_VENV_IN_PROJECT=1 pipenv install --deploy

FROM gcr.io/distroless/python3 as runtime-env
WORKDIR /app
COPY --from=build-env /app/.venv /app/.venv
ENV PATH="/app/.venv/bin:$PATH"
COPY worker .
CMD ["main.py"]
