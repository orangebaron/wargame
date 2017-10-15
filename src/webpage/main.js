var mapCenter = [0, 0]
const squareSize = 50
var map = {} // map of "x,y" to unit at that space
const unitColors = {
  water: '#008080'
}
const canvas = document.getElementById('maincanvas')
const context = canvas.getContext('2d')
function fixCanvasDimensions () {
  canvas.width = canvas.offsetWidth
  canvas.height = canvas.offsetHeight
}
function drawUnit (x, y, u) {
  context.fillStyle = unitColors[u]
  context.fillRect(((x - mapCenter[0] - 0.5) * squareSize) + (canvas.width / 2),
    ((y - mapCenter[1] - 0.5) * squareSize) + (canvas.height / 2), squareSize, squareSize)
}

var exampleSocket = new window.WebSocket('ws://localhost:8080/socket')
exampleSocket.onopen = startGame
exampleSocket.onmessage = function (msg) {
  var data = msg.data
  console.log(data)
  if (data.substring(0, 6) === 'block:') { // server sends a block's data
    data = data.substring(6).split(',') // format: ['x','y','unit']
    drawUnit(parseInt(data[0]), parseInt(data[1]), data[2])
    map[data[0] + ',' + data[1]] = data[2]
  }
}

function startGame () {
  exampleSocket.send('registerplr:uwe')
  exampleSocket.send('gettiles:-5,5,-5,5')
  fixCanvasDimensions()
}
