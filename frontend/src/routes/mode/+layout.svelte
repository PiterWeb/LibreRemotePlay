<script lang="ts">
	import { beforeNavigate, goto, onNavigate } from '$app/navigation';
	import BackwardButton from '$lib/layout/BackwardButton.svelte';
	import { showToast, ToastType } from '$lib/toast/toast_hook';
	import { _ } from 'svelte-i18n'


	import { CloseClientConnection } from '$lib/webrtc/client_webrtc_hook';
	import { CloseHostConnection } from '$lib/webrtc/host_webrtc_hook';
	import Modal from '$lib/layout/Modal.svelte';
	/** @type {{children?: import('svelte').Snippet}} */
	let { children } = $props();
	
	let modalElement: Modal
	
	let toUrl = $state<string | null>()
	
	function handleToast() {
		showToast($_('you-are-now-disconnected'), ToastType.INFO);
	}

	function closeConnection() {
		try {
			CloseClientConnection(handleToast);
			CloseHostConnection(handleToast);
		} catch {}
	}

	beforeNavigate(async (navigator) => {
		const nextPathname = navigator.to?.url.pathname ?? '';

		const actualPathname = navigator.from?.url.pathname ?? '';

		// If the user is navigating to the same page or one level up, we don't want to close the connection
		// but if goes one level down, we want to close the connection
		if (actualPathname === nextPathname) return;
		if (actualPathname.includes('/mode/client') && nextPathname.includes(actualPathname)) return;
		if (actualPathname.includes('/mode/host') && nextPathname.includes(actualPathname)) return;
		if (actualPathname.includes('/mode/config')) return;
		
		// If the user is redirected by the app should continue
		if (navigator.type === "goto") return
		
		// If the user tries to leave the page, we will show the browser's dialog to confirm the action, in other cases it is show a custom dialog
		navigator.cancel() 
		
		if (navigator.type === "leave") return
		
		toUrl = navigator.to?.route.id
		
		modalElement.openModal()
		
	});
	
	function onLeave() {
        closeConnection()
        goto(toUrl ?? "/")
	}
	
</script>

<Modal bind:this={modalElement}>
    
    <p>{$_('are-you-sure-you-want-to-leave')}</p>
    
    <button class="btn btn-primary">{$_('no')}</button>
    
    <button onclick={onLeave} class="btn btn-neutral">{$_('yes')}</button>
    
</Modal>

<BackwardButton path="/" />

{@render children?.()}
