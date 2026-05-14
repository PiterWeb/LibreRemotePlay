enum MouseType {
	Click = 0,
	Move = 1
}

enum MouseValue {
	Left = 0,
	Central = 1,
	Right = 2
}

enum MouseState {
  MouseDown = 0,
  MouseUp = 1
}

type mouseHandler = (output: ArrayBuffer) => void;


const click_events = ['mousedown', 'mouseup'] as const;

export const mouseLatency = $state({ value: 0 })

export function handleClick(callback: mouseHandler) {
  const mouseElement = document.getElementById("mouse-element")
  
	const handler = (event: MouseEvent) => {
		// console.log(`Click ${event.type}: ${event.button}`);

		const btnCLicked = event.button as MouseValue;
		
		const buf = new ArrayBuffer(3)
		const view = new Uint8Array(buf);
		
		const stateBtn = event.type === "mousedown" ? MouseState.MouseDown : MouseState.MouseUp
		
		view[0] = MouseType.Click
		view[1] = btnCLicked 
		view[2] = stateBtn
		
		if (mouseLatency.value === 0) return callback(buf)
		
		return setTimeout(() => callback(buf), mouseLatency.value) 
	};

	click_events.forEach((event_name) => mouseElement?.addEventListener(event_name, handler, true));

	return handler;
}

export function unhandleClick(callback: ReturnType<typeof handleClick>) {
  const mouseElement = document.getElementById("mouse-element")
	click_events.forEach((event_name) => mouseElement?.removeEventListener(event_name, callback, true));
}

export function handleMove(callback: mouseHandler) {
  
  const mouseElement = document.getElementById("mouse-element")
  
  
  const handler = (event: MouseEvent) => {
    if (!mouseElement) return
    
    const xAxis = event.offsetX;
    const yAxis = event.offsetY;
		console.log(`Move x:${xAxis}, y:${yAxis}`);
		
		const buf = new ArrayBuffer(1 + (2 * 4))
		const view = new DataView(buf);
		
		view.setUint8(0, MouseType.Move)
		view.setUint16(1, xAxis)
    view.setUint16(1 + 2, yAxis)
    view.setUint16(3 + 2, mouseElement.offsetWidth)
    view.setUint16(5 + 2, mouseElement.offsetHeight)
		
		if (mouseLatency.value === 0) return callback(buf)
		
		return setTimeout(() => callback(buf), mouseLatency.value)
	};

	mouseElement?.addEventListener("mousemove", handler, true);

	return handler;
}

export function unhandleMove(callback: ReturnType<typeof handleMove>) {
  const mouseElement = document.getElementById("mouse-element")
  
  mouseElement?.removeEventListener("mousemove", callback, true);
}

