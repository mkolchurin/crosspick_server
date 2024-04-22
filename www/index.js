;(() => {
  // expectingMessage is set to true
  // if the user has just submitted a message
  // and so we should scroll the next message into view when received.
  let expectingMessage = false
  function dial() {
    const conn = new WebSocket(`ws://${location.host}/subscribe`)

    conn.addEventListener('close', ev => {
      appendLog(`WebSocket Disconnected code: ${ev.code}, reason: ${ev.reason}`, true)
      if (ev.code !== 1001) {
        appendLog('Reconnecting in 1s', true)
        setTimeout(dial, 1000)
      }
    })
    conn.addEventListener('open', ev => {
      console.info('websocket connected')
    })

    // This is where we handle messages received.
    conn.addEventListener('message', ev => {
      if (typeof ev.data !== 'string') {
        console.error('unexpected message type', typeof ev.data)
        return
      }
      // id = Number.parseInt(ev.data)
      document.getElementById(ev.data).children[0].src = "http://localhost/api/v1/storage/maps/default.jpg"
      console.log(ev.data)
    })
  }
  dial()

  const publishForm = document.getElementById('publish-form')

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
    // messageInput.value = ''

    expectingMessage = true
    try {
      const resp = await fetch('/publish', {
        method: 'POST',
        body: msg,
      })
      if (resp.status !== 202) {
        throw new Error(`Unexpected HTTP Status ${resp.status} ${resp.statusText}`)
      }
    } catch (err) {
      appendLog(`Publish failed: ${err.message}`, true)
    }
  }
})()
