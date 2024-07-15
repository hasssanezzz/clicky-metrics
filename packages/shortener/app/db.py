import os
from peewee import *

from .utils import generate_short

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


class Url(BaseModel):
    user_username = CharField(max_length=225)
    short = CharField(max_length=255, unique=True)
    long = CharField(max_length=1024)

    def to_dict(self):
        return {
            'id': self.id,
            'user_username': self.user_username,
            'short': self.short,
            'long': self.long
        }
        
class UrlRepo:
    @staticmethod
    def getByUsername(username):
        try:
            query = Url.select(Url).where(Url.user_username == username)
            return [url.to_dict() for url in query]
        except:
            return None
    
    @staticmethod
    def getByShort(short):
        try:
            url = Url.select(Url).where(Url.short == short).get()
            return url.to_dict()
        except:
            return None
    
    @staticmethod
    def create(username, long):
        try:
            url = Url(username=username, short=generate_short(), long=long)
            url.save()
            return url
        except:
            return None
        
    @staticmethod
    def update(short, long):
        Url.update(long=long).where(Url.short == short).execute()
        
    @staticmethod
    def delete(short):
        Url.delete().where(Url.short == short).execute()