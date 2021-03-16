import http.client
import time

while True:
    
    conn = http.client.HTTPConnection("localhost", 8080)
    payload = ''
    headers = {}
    conn.request("GET", "/log?n=20", payload, headers)
    res = conn.getresponse()
    data = res.read()
    print(data.decode("utf-8"))

    time.sleep(1)