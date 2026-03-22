<script lang="ts">
	import { _ } from 'svelte-i18n';
</script>

<section
	id="tutorial-mouse-card"
	class="min-w-full p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8"
>
	<form class="max-w-lg mx-auto space-y-6" method="GET" action="/mode/config/mouse">
		<label for="helper-text" class="text-xl font-medium text-gray-900"
			>{$_('mouse-title')}</label
		>
		<br />

		<div class="max-w-lg space-y-4 flex justify-center h-full align-middle">
			<button type="submit" class="btn btn-primary text-white w-full">{$_('config')}</button>
		</div>
	</form>
</section>
