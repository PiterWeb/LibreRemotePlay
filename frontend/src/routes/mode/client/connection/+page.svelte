<script lang="ts">
	import { CloseStreamClientConnection } from '$lib/webrtc/stream/client_stream_hook';
	import { videoSpeedOptimizationEnabled } from '$lib/webrtc/stream/stream_config.svelte';
	import {consumingStream } from '$lib/webrtc/stream/stream_signal_hook.svelte';
	import { _ } from 'svelte-i18n';
	import { gamepadLatency } from '$lib/devices/gamepad/gamepad_hook.svelte';
	import { keyboardLatency } from '$lib/devices/keyboard/keyboard_hook.svelte';
	import { mouseLatency } from '$lib/devices/mouse/mouse_hook.svelte';
	
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

	<div class="flex flex-row gap-10 w-full mt-6 mx-auto md:col-span-4">
		<div class="w-full flex flex-col">
		    <h4 class="text-lg text-gray-300">{$_('gamepad-latency-ms')}</h4>
            <input
           	type="number"
           	class="input w-16 px-2 h-10 bg-neutral border border-gray-400 text-white text-center"
           	bind:value={gamepadLatency.value}
           	pattern="0|[1-9]\d*"
           	min={0}
           	step="5"
           	/>
		</div>
		<div class="w-full flex flex-col">
		    <h4 class="text-lg text-gray-300">{$_('mouse-latency-ms')}</h4>
            <input
           	type="number"
           	class="input w-16 px-2 h-10 bg-neutral border border-gray-400 text-white text-center"
           	bind:value={mouseLatency.value}
           	pattern="0|[1-9]\d*"
           	min={0}
           	step="5"
           	/>
		</div>
		<div class="w-full flex flex-col">
		    <h4 class="text-lg text-gray-300">{$_('keyboard-latency-ms')}</h4>
            <input
           	type="number"
           	class="input w-16 px-2 h-10 bg-neutral border border-gray-400 text-white text-center"
           	bind:value={keyboardLatency.value}
           	pattern="0|[1-9]\d*"
           	min={0}
           	step="5"
           	/>
		</div>
	</div>
	
	{:else}

	<section class="flex flex-row-reverse items-center gap-2">
		<label
			id="label-external-whip-checkbox"
			for="lan-mode-checkbox"
			class="font-semibold text-white">{$_('config-video-speed-optimization')}</label
		>
		<input
			id="lan-mode-checkbox"
			type="checkbox"
			class="checkbox checkbox-xs checkbox-primary"
			checked={videoSpeedOptimizationEnabled.value}
			onchange={() => videoSpeedOptimizationEnabled.value = !videoSpeedOptimizationEnabled.value}
		/>
	</section>
	
	<button class="md:col-span-2 btn btn-primary" onclick={connectToStream}>
		{$_('connect-to-stream')}</button>
		
	{/if}
</div>

