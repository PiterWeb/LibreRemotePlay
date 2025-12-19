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

const click_events: (keyof HTMLVideoElementEventMap)[] = ['mousedown', 'mouseup'];

export function handleClick(callback: mouseHandler) {
	const handler = (event: Event) => {
		if (!(event instanceof MouseEvent)) return;
		console.log(`Click ${event.type}: ${event.button}`);

		const btnCLicked = event.button as MouseValue;
		
		const buf = new ArrayBuffer(3)
		const view = new Uint8Array(buf);
		
		const stateBtn = event.type === "mousedown" ? MouseState.MouseDown : MouseState.MouseUp
		
		view[0] = MouseType.Click
		view[1] = btnCLicked 
		view[2] = stateBtn
		
		return callback(buf);
	};

	click_events.forEach((event_name) => document?.addEventListener(event_name, handler, true));

	return handler;
}

export function unhandleClick(callback: ReturnType<typeof handleClick>) {
	click_events.forEach((event_name) => document?.removeEventListener(event_name, callback, true));
}

export function handleMove(callback: mouseHandler) {
	const handler = (event: Event) => {
		if (!(event instanceof MouseEvent)) return;

    const xAxis = event.pageX;
    const yAxis = event.pageY;
		console.log(`Move x:${xAxis}, y:${yAxis}`);
		
		const buf = new ArrayBuffer(1 + (8 * 8) * 2)
		const view = new DataView(buf);
		
		view.setUint8(0, MouseType.Move)
		view.setBigUint64(1, BigInt(xAxis))
		view.setBigUint64(2 + 8 * 8, BigInt(yAxis))
		
		return callback(buf);
	};

	document?.addEventListener("mousemove", handler, true);

	return handler;
}

export function unhandleMove(callback: ReturnType<typeof handleClick>) {
  document?.removeEventListener("mousemove", callback, true);
}

