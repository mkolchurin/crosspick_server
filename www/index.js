;(() => {
  // expectingMessage is set to true
  // if the user has just submitted a message
  // and so we should scroll the next message into view when received.
  let expectingMessage = false
  var conn = null

  function update() {
    conn.send(JSON.stringify({ op: 'getMaps', body: {} }))
  }

  function dial() {
    setInterval(() => {
      update()
    }, 1000);
    conn = new WebSocket(`ws://${location.host}/api/v1/decider/0`)

    conn.addEventListener('close', ev => {
      appendLog(`WebSocket Disconnected code: ${ev.code}, reason: ${ev.reason}`, true)
      if (ev.code !== 1001) {
        appendLog('Reconnecting in 1s', true)
        setTimeout(dial, 1000)
      }
    })
    conn.addEventListener('open', ev => {
      console.info('websocket connected')
      a = JSON.stringify({ op: 'register', body: {username: "User1", uid: "1"} })
      console.info(a)
      conn.send(a)
      conn.send(JSON.stringify({ op: 'getMaps', body: {} }))
    })

    // This is where we handle messages received.
    conn.addEventListener('message', ev => {
      if (typeof ev.data !== 'string') {
        console.error('unexpected message type', typeof ev.data)
        return
      }
      console.log(ev.data)
      publishForm.innerHTML = ""
      JSON.parse(ev.data).forEach(element => {
        const publishForm = document.getElementById('mapContainer')
        
        var button = document.createElement("button")
        button.id = element.name
        button.type = "submit"
        button.className = "map"
        publishForm.appendChild(button)
        if(element.checked) {
          button.style.backgroundColor = "green"
        }

        img = document.createElement("img")
        img.src = "http://1.mkolchurin.ru:9988/api/v1/storage/maps/" + element.img
        img.width = 200
        img.height = 200
        button.appendChild(img)
        
        
      })
    })

    
  }
  dial()

  const publishForm = document.getElementById('mapContainer')
  const clearButton = document.getElementById('clear-button')

  clearButton.onclick = ev => {
    conn.send(JSON.stringify({ op: 'clear', body: {} }))
    conn.send(JSON.stringify({ op: 'getMaps', body: {} }))
  }

  // appendLog appends the passed text to messageLog.
  function appendLog(text, error) {
    console.log(text)
  }

  // onsubmit publishes the message from the user when the form is submitted.
  publishForm.onsubmit = async ev => {
    ev.preventDefault()
    const msg = ev.submitter.id
    if (msg === '') {
      return
    }
    a = JSON.stringify({ op: 'selectMap', body: {uid: "1", mapName: ev.submitter.id.toString(), checked: true} })
    conn.send(a)
    conn.send(JSON.stringify({ op: 'getMaps', body: {} }))
  }
})()
