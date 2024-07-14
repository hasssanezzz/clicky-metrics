from flask import jsonify
from flask_restful import Resource

from db import Url


class UrlResource(Resource):
    def get(self):
        urls = Url.select()
        return jsonify([{'id': url.id, 'username' : url.user_username , 'short': url.short, 'long': url.long } for url in urls])
