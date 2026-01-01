<script lang="ts">
	import { IS_RUNNING_EXTERNAL} from '$lib/detection/onwebsite';
	import type { Snippet } from 'svelte';
	let { children, not = false }: { children?: Snippet; not?: boolean } = $props();
</script>

{#if IS_RUNNING_EXTERNAL}
	{#if !not}
		{@render children?.()}
	{/if}
{:else if not}
	{@render children?.()}
{/if}
