# coding:utf-8

import requests
import json
import time


start_time = time.time()
session = requests.Session()

base_url = 'http://127.0.0.1:8080/'

print(json.dumps({
	'mobile': '17792518420',
	'pwd': 'yu226520'
	}))

r = session.post(base_url + 'login', data=json.dumps({
	'mobile': '17792518420',
	'pwd': 'yu226520'
	}))

print(r.ok, r.json())


if r.ok:
	print('1---------->', time.time() - start_time)
	token = r.json()['msg']['token']
	print(token)
	rs = session.get(base_url + 'torch/loads', headers={
		'token': token
		}, params={'plantid': 65})
	print('2---------->', time.time() - start_time)
	print(rs.ok, rs.json())



# r = session.post(base_url + 'register', data=json.dumps({
# 	'name': 'admin4',
# 	'mobile': '17792518417',
# 	'pwd': '123456',
# 	'permission': '0'
# 	}))
# print(r.ok, r.text)