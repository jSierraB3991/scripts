FROM python:3.12.0rc2-alpine3.18

WORKDIR /usr/src/app

COPY ./src/requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

ENV DJANGO_SUPERUSER_USERNAME=django
ENV DJANGO_SUPERUSER_PASSWORD=123456
ENV DJANGO_SUPERUSER_EMAIL=admin@example.com

COPY ./src .

EXPOSE 8000