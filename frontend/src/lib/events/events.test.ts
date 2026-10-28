import { Events, EventsEmit, EventsOn, EventsOnce } from "$lib/events/events"
import { afterEach, beforeEach, test, vi } from "vitest"


function delay(t: number) {
  return new Promise(resolve => setTimeout(resolve, t));
}

//@ts-ignore
globalThis.location = {
  port: "8090",
  hostname: "localhost",
  origin: "http://localhost:8090"
}

// beforeEach(() => {
//   vi.useFakeTimers()
// })
// 
// afterEach(() => {
//   vi.clearAllMocks()
// })

test('event listen, emit and then cancel', async (t) => {

  let i = 0
  
  const unlisten = EventsOn(Events.error, () => {
    i++
  })

  // Handled events
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))

  await delay(500)
  
  unlisten()
  
  
  // Unhandled events
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))

  await delay(500)
  t.expect(i).toBe(2)
  

})

test('once listener', async (t) => {
  let i = 0

  EventsOnce(Events.error, () => i++)
  
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))
  EventsEmit(Events.error, JSON.stringify({}))

  await delay(500)
  
  t.expect(i).toBe(1)

})

test('once listener cancel before emit event', async (t) => {
  let i = 0

  const unlistener = EventsOnce(Events.error, () => i++)

  unlistener()
  
  EventsEmit(Events.error, JSON.stringify({}))

  await delay(500)
  
  t.expect(i).toBe(0)

})

test('multiple listener, one emit', async (t) => {

  let i = 0;

  EventsOnce(Events.error, () => i++)
  EventsOnce(Events.error, () => i++)

  EventsEmit(Events.error, JSON.stringify({}))

  await delay(500)
  
  t.expect(i).toBe(2)

})

test("multiple on listener", async (t) => {

  let i = 0

  EventsOn(Events.streaming_signal_client, () => i++)
  EventsOn(Events.streaming_signal_client, () => i++)

  EventsEmit(Events.streaming_signal_client, JSON.stringify({}))
  EventsEmit(Events.streaming_signal_client, JSON.stringify({}))
  EventsEmit(Events.streaming_signal_client, JSON.stringify({}))
  EventsEmit(Events.streaming_signal_client, JSON.stringify({}))

  await delay(500)
  t.expect(i).toBe(8)

})

test('streaming_signal_server', async (t) => {

  let data: unknown;
  
  EventsOnce(Events.streaming_signal_server, (d) => {
    data = d
  })

  EventsEmit(Events.streaming_signal_server, {test: true})

  await delay(500)
  
  t.expect(data).toBe(JSON.stringify({ test: true }))
  t.annotate(JSON.stringify(data))

})