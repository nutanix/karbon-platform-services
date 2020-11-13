FROM python:3

WORKDIR /usr/src/app

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY loadgen.py /
COPY run.sh /

CMD [ "/run.sh" ]
