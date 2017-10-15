var exampleSocket = new window.WebSocket('ws://localhost:8080/socket')
exampleSocket.onopen = function () {
  exampleSocket.send('registerplr:uwe')
  exampleSocket.send('gettiles:-10,10,-10,10')
}
exampleSocket.onmessage = function (msg) {
  console.log(msg.data)
}
