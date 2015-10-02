#!/usr/bin/env ruby

require 'webrick'
require 'net/http'
require 'uri'
require 'mysql'

server = WEBrick::HTTPServer.new :Port => 8081

trap 'INT' do server.shutdown end

server.mount_proc '/pieRequest' do |req, res|
  #assess the request
  cnt = req.query()['count'].to_i
  puts "%d pies requested" % cnt
  sleep(1)

  #get the ingredients
  con = Mysql.new 'bistrofridge.deterlab.net', 'root', 'pizza', 'bistrofridge'
  rs = con.query "SELECT count FROM ingredient_packs"
  record = rs.fetch_hash
  remaining = record["count"].to_i
  puts "the bistro fridge has #{(remaining)} ingredient packs remaining"

  if remaining == 0
    req.status = 500
    res.body = "we are out of ingredients"
    return
  end
  cmd = con.prepare "UPDATE ingredient_packs SET count = ?"
  cmd.execute (remaining - cnt)
  
  rs = con.query "SELECT count FROM doughballs"
  record = rs.fetch_hash
  remaining = record["count"].to_i
  puts "the bistro fridge has #{(remaining)} doughballs remaining"

  if remaining == 0
    req.status = 500
    res.body = "we are out of doughballs"
    return
  end
  cmd = con.prepare "UPDATE doughballs SET count = ?"
  cmd.execute (remaining - cnt)

  
  #bake the pies
  uri = URI.parse("http://woodbrick.deterlab.net:8082")
  http = Net::HTTP.new(uri.host, uri.port)
  request = Net::HTTP::Get.new("/cook")
  response = http.request(request)
  puts "The oven says %s" % response.body
  sleep(1)

  #return the result
  res.body = cnt.to_s
end

server.start
