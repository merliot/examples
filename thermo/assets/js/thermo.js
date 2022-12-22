var conn

var all = document.getElementById("all")
var furnace = document.getElementById("furnace")
var aircond = document.getElementById("aircond")
var sensor = document.getElementById("sensor")
var furnace_link = document.getElementById("furnace_link")
var aircond_link = document.getElementById("aircond_link")
var sensor_link = document.getElementById("sensor_link")
var slider = document.getElementById("slider")
var sp = document.getElementById("sp")
var temp = document.getElementById("temp")
var relay0 = document.getElementById("relay0")
var relay1 = document.getElementById("relay1")

function clear() {
	all.style.display = "none"
}

function refresh(msg) {
	if (msg.Relays.States[0]) {
		relay0.textContent = "ON"
		relay0.style.backgroundColor = "lightgreen"
		relay0.style.color = "black"
	} else {
		relay0.textContent = "OFF"
		relay0.style.backgroundColor = "red"
		relay0.style.color = "white"
	}
	if (msg.Relays.States[1]) {
		relay1.textContent = "ON"
		relay1.style.backgroundColor = "lightgreen"
		relay1.style.color = "black"
	} else {
		relay1.textContent = "OFF"
		relay1.style.backgroundColor = "red"
		relay1.style.color = "white"
	}
	if (msg.Relays.Online) {
		furnace.style.backgroundColor = "lightblue"
		aircond.style.backgroundColor = "lightblue"
		furnace_link.href = "/" + encodeURIComponent(msg.Relays.Id)
		aircond_link.href = "/" + encodeURIComponent(msg.Relays.Id)
	} else {
		furnace.style.backgroundColor = "lightgrey"
		aircond.style.backgroundColor = "lightgrey"
		relay0.style.backgroundColor = "lightgrey"
		relay1.style.backgroundColor = "lightgrey"
		relay0.style.color = "grey"
		relay1.style.color = "grey"
		furnace_link.href = ""
		aircond_link.href = ""
	}
	if (msg.Sensors.Online) {
		sensor.style.backgroundColor = "lightblue"
		temp.style.backgroundColor = "lightgreen"
		sensor_link.href = "/" + encodeURIComponent(msg.Sensors.Id)
	} else {
		sensor.style.backgroundColor = "lightgrey"
		temp.style.backgroundColor = "lightgrey"
		sensor_link.href = ""
	}

	slider.value = msg.SetPoint
	sp.innerHTML = msg.SetPoint + "&deg;F"
	temp.innerHTML = msg.Sensors.Temp + "&deg;F"
	all.style.display = "flex"
}

function setpoint(val) {
	conn.send(JSON.stringify({Msg: "SetPoint", Val: parseInt(val)}))
}

function Run(ws) {

	function connect() {
		conn = new WebSocket(ws)

		conn.onopen = function(evt) {
			clear()
			conn.send(JSON.stringify({Msg: "_GetState"}))
		}

		conn.onclose = function(evt) {
			clear()
			setTimeout(connect, 1000)
		}

		conn.onerror = function(err) {
			conn.close()
		}

		conn.onmessage = function(evt) {
			var msg = JSON.parse(evt.data)

			console.log('thermo', msg)

			switch(msg.Msg) {
			case "_ReplyState":
				refresh(msg)
				break
			}
		}
	}

	connect()
}
