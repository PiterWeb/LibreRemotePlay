<script lang="ts">
  import {  usePreferedCodecsOrderedStorage,restoreDefaultCodecs, preferedCodecsOrdered } from './stream_config.svelte';
	import { _ } from 'svelte-i18n';
	import { useSortable, reorder } from '$lib/layout/useSortable.svelte';

  let sortable = $state<HTMLElement | null>(null);

  usePreferedCodecsOrderedStorage()

  useSortable(() => sortable, {
        group: "codec-list",
        animation: 200,
        ghostClass: 'opacity-0',
        onEnd(evt) {
           preferedCodecsOrdered.value = reorder(preferedCodecsOrdered.value, evt);
        }
  });

</script>

<section class="md:col-span-2 flex flex-col items-center gap-4 w-full">
  
  <div class="flex flex-col border w-full shadow-2xs rounded-xl bg-neutral-900 border-neutral-700 shadow-neutral-700/70">
    <header class="p-4 md:p-5">
      <h3 class="text-lg font-bold text-white">
        {$_("codec-list")}
      </h3>
      <p class="mt-2 text-neutral-400">
        {$_("codec-list-preference")}
      </p>
      
    </header>
    <footer class="border-t rounded-b-xl py-3 px-4 md:py-4 md:px-5 bg-neutral-900 border-neutral-700">
      <ul bind:this={sortable} class="max-w-xs flex flex-col">
        {#each preferedCodecsOrdered.value as codec}
          {#key codec}
            <li data-id={codec} class="inline-flex items-center gap-x-3 py-3 px-4 cursor-grab text-sm font-medium border -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg bg-neutral-900 border-neutral-700 text-neutral-200">
              {codec}
              <svg class="shrink-0 size-4 ms-auto text-neutral-500" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="9" cy="12" r="1"></circle>
                <circle cx="9" cy="5" r="1"></circle>
                <circle cx="9" cy="19" r="1"></circle>
                <circle cx="15" cy="12" r="1"></circle>
                <circle cx="15" cy="5" r="1"></circle>
                <circle cx="15" cy="19" r="1"></circle>
              </svg>
            </li>
          {/key}
        {/each}
      </ul>
    </footer>
  
  </div>
  
  <button onclick={restoreDefaultCodecs} class="btn">
    {$_('restore-codecs')}
  </button>
</section>

<!-- This make tailwind compile opacity-0 if it is not used already -->
<div class="hidden opacity-0"></div>