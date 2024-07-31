import os
import jwt
from flask import Flask, request, jsonify, make_response
from flask_restful import Api
from .resources import UrlResource, RedirectResource
from .constants import AUTH_TOKEN_NAME

def create_app():
    app = Flask(__name__)
    api = Api(app)
    jwt_secret = os.getenv("API_JWT_SECRET")
    
    @app.before_request
    def validteGatewayToken():
        auth_token = request.headers.get(AUTH_TOKEN_NAME)
        try:
            jwt.decode(auth_token, jwt_secret, algorithms=['HS256'])
        except:
            return make_response(jsonify(), 401)
    
    api.add_resource(UrlResource, '/v1')
    api.add_resource(RedirectResource, '/v1/<string:short>')
    
    return app

if __name__ == "__main__":
    app = create_app()
    app.run(port=80, debug=True)