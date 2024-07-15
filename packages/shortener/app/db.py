import datetime
import os
from peewee import *

database = PostgresqlDatabase(
    os.getenv("POSTGRES_DB"),
    user=os.getenv("POSTGRES_USER"),
    password=os.getenv("POSTGRES_PASSWORD"),
    host=os.getenv("POSTGRES_HOST"),
    port=os.getenv("POSTGRES_PORT")
)

class BaseModel(Model):
    class Meta:
        database = database


class User(BaseModel):
    username = CharField(max_length=255, unique=True)
    email = CharField(max_length=255, unique=True)
    password = CharField(max_length=1024)
    created_at = DateTimeField(default=datetime.datetime.now)
    updated_at = DateTimeField(default=datetime.datetime.now)


class Url(BaseModel):
    # Optional foreign key to user by username
    user_username = CharField(max_length=225)
    short = CharField(max_length=255, unique=True)
    long = CharField(max_length=1024)
