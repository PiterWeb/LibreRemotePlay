export enum Events {
  connection_state = "connection_state",
  error = "ERROR", 
  streaming_signal_client = "streaming-signal-client",
  streaming_signal_server = "streaming-signal-server"
}

const baseUrl = () => location.port === "5173" ? `http://${location.hostname}:8090/events` : `${location.origin}/events`

function createWS(topic: Events, once = false) {
  console.log(`${baseUrl()}/listen/${topic}`)
  if (once) return new WebSocket(`${baseUrl()}/listen/${topic}?once=true`)
  else return new WebSocket(`${baseUrl()}/listen/${topic}`)
}

type Clbk = (data: any) => void

function createListener(callback: Clbk, onceCancel = () => {}) {
  return ({ data }: { data: any }) => {
    callback(data)
    onceCancel()    
  }
}

const eventsListenersMap = new Map<Events, WebSocket[]>()

function addEventSourceToMap(e: Events, eventSource: WebSocket) {
  eventsListenersMap.set(e, [...eventsListenersMap.get(e) ?? [], eventSource])
}

function removeWSEventFromMap(e: Events, eventSource: WebSocket) {
  eventsListenersMap.set(e, eventsListenersMap.get(e)?.filter(es => es !== eventSource) ?? [])
}

export function EventsOn(e: Events, callback: Clbk): () => void {  
  try {

    let closedByUser = false
    
    const ws = createWS(e)
    addEventSourceToMap(e, ws)
      
    const listener = createListener((data) => {
      if (closedByUser) return
      callback(data)
    })

    let unlistener = () => {
      closedByUser = true
      removeWSEventFromMap(e, ws)
      ws.send("cancel")
      ws.removeEventListener("message", listener)
      ws.close()
    }
    
    ws.addEventListener("message", listener)
    ws.addEventListener("close", (ev) => {
      if (closedByUser) return

      unlistener = EventsOn(e, callback)
    })
    
    return unlistener
    
  } finally {
    return () => {}
  }
}

export function EventsOnce(e: Events, callback: Clbk): () => void {
  try {

    let closedByUser = false
    
    const ws = createWS(e, true)
    addEventSourceToMap(e, ws)
    
    let unlistener = () => {
      closedByUser = true
      removeWSEventFromMap(e, ws)
      ws.send("cancel")
      ws.removeEventListener("message", listener)
      ws.close()
    }
    
    const listener = createListener((data) => {
      if (closedByUser) return
      callback(data)
    }, unlistener)
    
    ws.addEventListener("message", listener)
    ws.addEventListener("close", (ev) => {
      if (closedByUser) return

      unlistener = EventsOn(e, callback)
    })
    
    return unlistener
    
  } finally {
    return () => {}
  }
}

export function EventsOff(e: Events, ...additionalEventNames: Events[]): void {
  eventsListenersMap.get(e)?.forEach((ws) => {
    ws.send("cancel")
    ws.close()
  })

  eventsListenersMap.set(e, [])

  additionalEventNames.forEach((aEvent) => {
    eventsListenersMap.get(aEvent)?.forEach((ws) => {
      ws.send("cancel")
      ws.close()
    })

    eventsListenersMap.set(aEvent, [])
  })
}

export async function EventsEmit(topic: Events, data: any) {
  console.log(`${baseUrl()}/emmit/${topic}`)
  const res = await fetch(`${baseUrl()}/emmit/${topic}`, {
    method: "POST",
    body: JSON.stringify(data)
  })
    
  return res.ok
}