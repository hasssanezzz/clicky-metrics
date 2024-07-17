import re
from .constants import URL_REGEX, SHORT_URL_SIZE
from nanoid import generate


def api_error(data):
    return {
        'error': data,
        'meta': None
    }


def api_response(data):
    return {
        'data': data,
        'meta': None
    }


def generate_short():
    return generate(size=SHORT_URL_SIZE)


def read_url_or_return_error(jsonBody):
    err = api_error({'url': ['a valid URL must be provided']}), 400
    if 'url' not in jsonBody or not re.match(URL_REGEX, jsonBody['url']):
        return err, False
    return jsonBody['url'], True
