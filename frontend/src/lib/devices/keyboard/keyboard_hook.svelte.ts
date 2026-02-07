type keyHandler = (keycode: string) => void

export const keyboardLatency = $state({ value: 0 })

export function handleKeyDown(callback: keyHandler) {
	
	const handler = (event: KeyboardEvent) => {
		event.preventDefault()
    event.stopPropagation()
		
    if (keyboardLatency.value === 0) return callback(event.key + '_1')
    
		return setTimeout(() => callback(event.key + '_1'), keyboardLatency.value);
	}

	document.addEventListener('keydown', handler, true);

	return handler
}

export function unhandleKeyDown(callback: ReturnType<typeof handleKeyDown>) {
	document.removeEventListener("keydown", callback)
}

export function handleKeyUp(callback: keyHandler) {

	const handler = (event: KeyboardEvent) => {
		event.preventDefault()
    event.stopPropagation()
		
    if (keyboardLatency.value === 0) return callback(event.key + '_0')
    
		return setTimeout(() => callback(event.key + '_0'), keyboardLatency.value);
	}

	document.addEventListener('keyup', handler, true);

	return handler
}

export function unhandleKeyUp(callback: ReturnType<typeof handleKeyUp>) {
	document.removeEventListener("keyup", callback)
}