from flask import Flask
from flask_restful import Api
from resources import UrlResource

app = Flask(__name__)
api = Api(app)

api.add_resource(UrlResource, "/url")

if __name__ == "__main__":
    app.run(port=80, debug=True)