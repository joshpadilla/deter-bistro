#!/usr/bin/env python3

from bottle import request, route, run, template
import requests
import time

@route('/orderPie')
def orderPie():
    cnt = request.query.count
    print ("Taking order for " + cnt + " pies")
    time.sleep(1)
    r = requests.get('http://pizzachef.deterlab.net:8081/pieRequest?count='+cnt)
    pcnt = int(r.text)
    print ("Got " + str(pcnt) + " pizzas from the pizza chef")
    return str(r.text)

print ("Waiter is waiting for customer")
run(host='localhost', port=8080)
