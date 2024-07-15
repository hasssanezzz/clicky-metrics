FROM python:3.9-slim
WORKDIR /app
COPY . .

RUN pip3 install --no-cache-dir -r requirements.txt

EXPOSE 80
ENV FLASK_APP=/app/app/__init__.py
ENV FLASK_RUN_HOST=0.0.0.0
ENV FLASK_RUN_PORT=80
ENV FLASK_ENV=production

CMD ["flask", "run"]