#!/bin/sh
forever start /usr/local/deter-bistro/Mixer/mixer;
forever start /usr/local/deter-bistro/SousChef/go run souschef.go;
forever start /usr/local/deter-bistro/WoodbrickOven/java WoodbrickOven;
forever start /usr/local/deter-bistro/PizzaChef/pizzachef.rb;
forever start /usr/local/deter-bistro/Waiter/waiter.py;
forever start /usr/local/deter-bistro/Customer/node customer.js;