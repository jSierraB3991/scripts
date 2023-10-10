FROM python:3.12.0rc2-alpine3.18

WORKDIR /usr/src/app

COPY ./requirements.txt .
RUN apk add curl jq
RUN pip install --no-cache-dir -r requirements.txt

COPY ./src .

EXPOSE 5000

CMD [ "python3", "-m" , "flask", "run", "--host=0.0.0.0"]