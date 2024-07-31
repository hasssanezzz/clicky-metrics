from flask import request, redirect
from flask_restful import Resource

from .db import UrlRepo
from .constants import USERNAME_HEADER_NAME
from .utils import api_error, api_response, read_url_or_return_error
    

class UrlResource(Resource):
    def get(self):
        username = request.headers.get(USERNAME_HEADER_NAME)
        if not username:
            return api_error({'username': ['username header not found']}), 422
        
        results = UrlRepo.getByUsername(username)
        if results is not None:
            return api_response(results)
        return api_error({ 'root': ['internal server error']}), 500
    
    def post(self):
        result, ok = read_url_or_return_error(request.get_json())
        if not ok:
            return result
        
        username = request.headers.get(USERNAME_HEADER_NAME)
        if not username:
            username = None
            
        # TODO make sure username exists
        newUrl = UrlRepo.create(username=username, long=result)
        print('USERNAME:', username)
        if not newUrl:
            return api_error({ 'root': ['internal server error']}), 500
        return api_response(newUrl.to_dict())
    
    
class RedirectResource(Resource):
    def get(self, short):
        try:
            url = UrlRepo.getByShort(short)
            if not url:
                raise
            return url['long']
        except:
            return api_error({'short': ['short not found']}), 404
    
    def put(self, short):
        result, ok = read_url_or_return_error(request.get_json())
        if not ok:
            return result

        url = UrlRepo.getByShort(short)
        if not url:
            return api_error({'short': ['short not found']}), 404
            
        try:
            UrlRepo.update(short, result)
        except Exception as e:
            print(f'error trying to update url with short {short}: {e}')
            return api_error({ 'root': ['internal server error']}), 500
        else:
            url['long'] = result
            return api_response(url)
        
    
    def delete(self, short):
        url = UrlRepo.getByShort(short)
        if not url:
            return api_error({'short': ['short not found']}), 404
        
        try:
            UrlRepo.delete(short)
        except Exception as e:
            print(f'error trying to deleting url with short {short}: {e}')
            return api_error({ 'root': ['internal server error']}), 500
        else:
            return api_response(None)