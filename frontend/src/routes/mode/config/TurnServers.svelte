<script lang="ts">
	import { _ } from 'svelte-i18n';
</script>

<section
	id="tutorial-turn-card"
	class="min-w-full p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8"
>
	<form class="max-w-lg mx-auto space-y-6" method="get" action="/mode/config/advanced/turn">
		<label for="helper-text" class="text-xl font-medium text-gray-900"
			>{$_('turn-servers-title')}</label
		>
		<span class="text-sm text-gray-600">
			{$_('advance')}
		</span>

		<br />

		<a
			href="https://github.com/coturn/coturn"
			target="_blank"
			class="text-sm font-medium hover:underline text-blue-500"
			>{$_('selfhost-turn-link')}</a
		>

		<div class="max-w-lg space-y-4 flex justify-center h-full align-middle">
			<button type="submit" class="btn btn-primary w-full text-white">{$_('config')}</button>
		</div>
	</form>
</section>
