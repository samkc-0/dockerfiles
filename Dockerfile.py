FROM debian:stable-slim

RUN apt update
RUN apt upgrade -y

RUN apt install -y build-essential wget zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev libsqlite3-dev libbz2-dev

RUN wget "https://www.python.org/ftp/python/3.14.0/Python-3.14.0.tgz"
RUN tar -xf Python-3.14.0.tgz

RUN cd Python-3.14.0 && ./configure --enable-optimizations && make && make altinstall


COPY main.py main.py
COPY books/ books/
CMD ["python3.14", "main.py"]
