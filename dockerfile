# syntax=docker/dockerfile:1

FROM python:3.8-slim-buster

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip3 install -r requirements.txt

COPY . .

# CMD ["gunicorn", "-w", "2", "-b", "0.0.0.0:8000", "--log-level", "info", "app:app"]
CMD [ "python3", "app.py" ]
