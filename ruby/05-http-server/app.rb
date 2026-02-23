require 'socket'

server = TCPServer.new("localhost", 8080)

loop do
  client = server.accept
  request_line = client.gets
  verb, path = request_line.split

  if verb == 'GET' && path == '/'
    response = "HTTP/1.1 200 OK\r\n" +
               "Content-Type: text/html\r\n" +
               "Content-Length: 13\r\n" +
               "\r\n" +
               "<h1>Hello</h1>\r\n\r\n"
    client.print response
  else
    client.print "HTTP/1.1 404 Not Found\r\n\r\n"
  end

  client.close
end
