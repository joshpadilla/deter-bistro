
request = require('request')

console.log("Asking waiter for 47 pizzas")

request "http://waiter.deterlab.net:8080/orderPie?count=47", (error, response, body) =>
  if !error and response.statusCode == 200
    console.log("Received " + body + " pies from our waiter")
  else
    console.log(
      "Something terrible has happened to our waiter (" + response.statusCode + ")"
    )
