<script lang="ts">
	import StunServers from './StunServers.svelte';
	import IsWindows from '$lib/detection/IsWindows.svelte';
	import ViGEmDownload from './ViGEmDownload.svelte';
	import { _ } from 'svelte-i18n';
	import LanguageSelector from '$lib/i18n/LanguageSelector.svelte';
	import TurnServers from './TurnServers.svelte';
</script>

<h2 class="text-center text-white text-[clamp(2rem,6vw,4.2rem)] font-black leading-[1.1] xl:text-left">
		{$_('config_title')}
</h2>

<div class="grid gap-28 w-full md:grid-cols-3 py-12 xl:[&>*]:w-80 lg:[&>*]:w-72 md:[&>*]:w-60 [&>*]:w-72">
	<LanguageSelector />
	<StunServers />
	<TurnServers />

	<IsWindows>
		<ViGEmDownload />
	</IsWindows>
</div>
