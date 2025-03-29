type keyHandler = (keycode: string) => void

export function handleKeyDown(callback: keyHandler) {
	
	const handler = (event: KeyboardEvent) => {
		return callback(event.code + '_1');
	}

	document.addEventListener('keydown', handler);

	return handler
}

export function unhandleKeyDown(callback: ReturnType<typeof handleKeyDown>) {
	document.removeEventListener("keydown", callback)
}

export function handleKeyUp(callback: keyHandler) {

	const handler = (event: KeyboardEvent) => {
		return callback(event.code + '_0');
	}

	document.addEventListener('keyup', handler);

	return handler
}

export function unhandleKeyUp(callback: ReturnType<typeof handleKeyUp>) {
	document.removeEventListener("keyup", callback)
}

const specialKeys = new Set([
	'SHIFTLEFT',
	'SHIFTRIGHT',
	'CONTROLLEFT',
	'CONTROLRIGHT',
	'ALTLEFT',
	'ALTRIGHT'
]);
