FROM python:3

RUN pip3 install confluent-kafka numpy matplotlib wordcloud pandas
RUN mkdir /www
ADD wc.py /

CMD [ "python3", "/wc.py" ]
