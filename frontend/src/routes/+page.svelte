<script lang="ts">
	import onwebsite from '$lib/detection/onwebsite';
	import { _ } from 'svelte-i18n';
	import { StartTutorial } from '$lib/tutorial/driver';
	import IsLinux from '$lib/detection/IsLinux.svelte';

</script>

<button
	onclick={() => {
		StartTutorial();
	}}
	class="btn btn-primary text-white"
>
	{$_('tutorial_btn')}
</button>

<h2 class="text-center text-white text-[clamp(2rem,6vw,4.2rem)] font-black leading-[1.1] xl:text-left">
	<span
		class="text-primary"
		>{$_('main_title_choose')}
	</span>
	{$_('main_title_your')}
	<span
		class="text-primary"
		>{$_('main_title_role')}</span
	>
</h2>
<div id="tutorial-play" class="flex gap-4 mt-4 md:flex-row flex-col">
	{#if !onwebsite}
		<div
			class="card md:w-96 md:h-52 bg-white rounded-lg shadow border-gray-200 border"
		>
			<div class="card-body">
				<h2 class="card-title text-gray-800">{$_('host_card_title')}</h2>
				<p class="text-gray-600">{$_('host_card_description')}</p>
				<a href="/mode/host" class="btn btn-primary text-white">{$_('host_card_cta')}</a>
			</div>
		</div>
	{/if}
	<div
		class="card md:w-96 md:h-52 bg-white border border-gray-200 rounded-lg shadow"
	>
		<div class="card-body">
			<h2 class="card-title text-gray-800">{$_('client_card_title')}</h2>
			<p class="text-gray-600">{$_('client_card_description')}</p>
			<IsLinux>
			    <button onclick={async () => {
							const { BrowserOpenURL } = await import('$lib/wailsjs/runtime/runtime');
							BrowserOpenURL("http://localhost:8080/mode/client")
				}} class="btn btn-primary text-white">{$_('client_card_cta')}</button>
			</IsLinux>
			<IsLinux not>
			<a href="/mode/client" class="btn btn-primary text-white">{$_('client_card_cta')}</a>
			</IsLinux>
		</div>
	</div>
</div>
