<script lang="ts">
	import { CloseStreamClientConnection } from '$lib/webrtc/stream/client_stream_hook';
	import {consumingStream } from '$lib/webrtc/stream/stream_signal_hook.svelte';
	import { _ } from 'svelte-i18n';

	function connectToStream() {
		consumingStream.value = true;
	}

	function closeStream() {
		CloseStreamClientConnection()
	}

</script>

<div class="grid md:grid-cols-4 grid-cols-1 gap-10">
	{#if consumingStream.value}
	<button class="md:col-span-4 btn btn-primary btn-outline" onclick={closeStream}>
		{$_('close-stream')}</button>

	{:else}

	<button class="md:col-span-2 btn btn-primary" onclick={connectToStream}>
		{$_('connect-to-stream')}</button>
		
	{/if}
</div>

