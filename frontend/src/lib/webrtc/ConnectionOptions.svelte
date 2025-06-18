<script>
	import { _ } from 'svelte-i18n';
	import LANMode from '$lib/webrtc/lan_mode.svelte';
	import Tooltip from '$lib/layout/Tooltip.svelte';
</script>

<Tooltip ref="label-lan-mode-checkbox" placement="bottom">
    <p>{$_('lan-mode-explanation')}</p>
</Tooltip>

<section class="flex flex-row-reverse items-center gap-2">
	<label
		id="label-lan-mode-checkbox"
		for="lan-mode-checkbox"
		class="font-semibold text-gray-900 dark:text-white">{$_('lan-mode')}</label
	>
	<input
		id="lan-mode-checkbox"
		type="checkbox"
		class="checkbox checkbox-xs checkbox-primary"
		bind:checked={LANMode.enabled}
	/>
</section>
