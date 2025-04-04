type keyHandler = (keycode: string) => void

export function handleKeyDown(callback: keyHandler) {
	
	const handler = (event: KeyboardEvent) => {
		return callback(event.key + '_1');
	}

	document.addEventListener('keydown', handler);

	return handler
}

export function unhandleKeyDown(callback: ReturnType<typeof handleKeyDown>) {
	document.removeEventListener("keydown", callback)
}

export function handleKeyUp(callback: keyHandler) {

	const handler = (event: KeyboardEvent) => {
		return callback(event.key + '_0');
	}

	document.addEventListener('keyup', handler);

	return handler
}

export function unhandleKeyUp(callback: ReturnType<typeof handleKeyUp>) {
	document.removeEventListener("keyup", callback)
}