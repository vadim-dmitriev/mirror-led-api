# docker build -t mirror-led-api .
# docker run -d --restart unless-stopped -p 8090:8090 --privileged mirror-led-api:latest

FROM python:3.7-slim

WORKDIR /app

COPY requirements.txt requirements.txt

RUN apt-get update \
	&& apt-get install -y gcc \
	libc-dev \
    && rm -rf /var/lib/apt/lists/*

RUN pip3 install -r requirements.txt

COPY . .

EXPOSE 8090

CMD [ "python3", "main.py"]
