export enum Events {
  connection_state = "connection_state",
  error = "ERROR", 
  streaming_signal_client = "streaming-signal-client",
  streaming_signal_server = "streaming-signal-server"
}

const baseUrl = () => location.port === "5173" ? `http://${location.hostname}:8090/events` : `${location.origin}/events`

function createWS(topic: Events) {
  console.log(`${baseUrl()}/listen/${topic}`)
  return new WebSocket(`${baseUrl()}/listen/${topic}`)
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

function removeWSEventToMap(e: Events, eventSource: WebSocket) {
  eventsListenersMap.set(e, eventsListenersMap.get(e)?.filter(es => es !== eventSource) ?? [])
}

export function EventsOn(e: Events, callback: Clbk): () => void {  
  try {

    let closedByUser = false
    
    const ws = createWS(e)
    addEventSourceToMap(e, ws)
      
    const listener = createListener(callback)

    let unlistener = () => {
      closedByUser = true
      removeWSEventToMap(e, ws)
      ws.send("")
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
    
    const ws = createWS(e)
    addEventSourceToMap(e, ws)
    
    let unlistener = () => {
      closedByUser = true
      removeWSEventToMap(e, ws)
      ws.send("")
      ws.removeEventListener("message", listener)
      ws.close()
    }
    
    const listener = createListener(callback, unlistener)
    
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
    ws.send("")
    ws.close()
  })

  eventsListenersMap.set(e, [])

  additionalEventNames.forEach((aEvent) => {
    eventsListenersMap.get(aEvent)?.forEach((ws) => {
      ws.send("")
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