docker_build:
	pipenv lock -r > requirements.txt
	docker build -t doofi/tornado-worker:latest .

install:
	pipenv install --dev

run_flake8:
	pipenv run flake8

run_autopep8:
	pipenv run autopep8 --in-place --recursive .

run_tests:
	pipenv run python3 -m unittest discover tests -f

run_tests_coverage:
	pipenv run coverage run --source worker -m unittest discover
	pipenv run coverage report --fail-under=80

clean:
	find . \( -name __pycache__ -o -name "*.pyc" \) -delete
	pipenv run coverage erase
	pipenv --rm
