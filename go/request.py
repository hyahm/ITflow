import requests
import json
import os

url = "http://127.0.0.1:10001"



login_data = {
	"username": "admin@qq.com",
	"password": "admin"
}

r = requests.post(url + "/user/login", data=json.dumps(login_data) )

token = ""
if r.json()["code"] == 0:
	token = r.json()["token"]
else:
	print(r.json())
	os.exit(1)


headers = {
	"X-Token": token
}


status_update_data = {
	"id": 6,
	"name": "mytest"
}


r = requests.post(url + "/status/update", data=json.dumps(status_update_data), headers=headers )

print(r.json())



