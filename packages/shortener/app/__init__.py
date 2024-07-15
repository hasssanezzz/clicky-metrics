from flask import Flask
from flask_restful import Api
from .resources import UrlResource, RedirectResource

def create_app():
    app = Flask(__name__)
    api = Api(app)
    
    api.add_resource(UrlResource, '/v1')
    api.add_resource(RedirectResource, '/v1/<string:short>')
    
    return app

if __name__ == "__main__":
    app = create_app()
    app.run(port=80, debug=True)