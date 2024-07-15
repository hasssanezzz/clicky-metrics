import os
import requests

AUTH_HOST = os.getenv("AUTH_SERVICE_HOST")
AUTH_PORT = os.getenv("AUTH_SERVICE_PORT")
AUTH_URL = '/v1/validate'


class AuthService:
    @staticmethod
    def validate(token: str) -> str | None:
        if not token:
            return None
        try:
            headers = {'Authorization': token}
            res = requests.get(
                f'http://{AUTH_HOST}:{AUTH_PORT + AUTH_URL}', headers=headers)
            if res.status_code == 200:
                return res.json()['data']
            return None
        except:
            return None
