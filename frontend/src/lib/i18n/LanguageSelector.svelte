<script lang="ts">
	import {locale, locales, _ } from 'svelte-i18n';

	async function setLocale(event: Event) {
		const target = event.target as HTMLButtonElement;
		await locale.set(target.dataset.value);
	}

	function getLocaleName(locale: string) {
		return new Intl.DisplayNames(locale, { type: 'language' }).of(locale) ?? '';
	}
</script>

<section
	id="tutorial-language"
	class="
    min-w-full p-4 bg-white border border-gray-200 rounded-lg shadow sm:p-6 md:p-8"
>
	<label for="helper-text" class="text-xl font-medium text-gray-900"
		>{$_('language-selector-title')}</label
	>

	<ul class="grid grid-cols-2 justify-between w-full">
		{#each $locales as locale}
			<li class="w-full my-4 flex justify-center">
				<button onclick={setLocale} data-value={locale} class="btn btn-primary text-white"
					>{getLocaleName(locale)}</button
				>
			</li>
		{/each}
	</ul>
</section>
