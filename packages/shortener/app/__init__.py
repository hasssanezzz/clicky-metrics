from flask import Flask
from flask_restful import Api
from .resources import UrlResource

def create_app():
    app = Flask(__name__)
    api = Api(app)
    
    api.add_resource(UrlResource, '/v1/')
    
    return app

if __name__ == "__main__":
    app = create_app()
    app.run(port=80, debug=True)