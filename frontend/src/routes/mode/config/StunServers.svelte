<script lang="ts">
	import { _ } from 'svelte-i18n';
</script>

<section
	id="tutorial-stun-card"
	class="min-w-full p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8"
>
	<form class="max-w-lg mx-auto space-y-6" method="GET" action="/mode/config/advanced/stun">
		<label for="helper-text" class="text-xl font-medium text-gray-900"
			>{$_('stun-servers-title')}</label
		>
		<span class="text-sm text-gray-600">
			{$_('advance')}
		</span>
		<br />
		<a
			href="https://raw.githubusercontent.com/pradt2/always-online-stun/master/valid_hosts.txt"
			target="_blank"
			class="text-sm font-medium hover:underline text-blue-500"
			>{$_('server-list-link')}</a
		>

		<div class="max-w-lg space-y-4 flex justify-center h-full align-middle">
			<button type="submit" class="btn btn-primary text-white w-full">{$_('config')}</button>
		</div>
	</form>
</section>
