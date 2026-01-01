<script lang="ts">
	import onwebsite from '$lib/detection/onwebsite';
	import type { Snippet } from 'svelte';
	let { children, not = false }: { children?: Snippet; not?: boolean } = $props();
</script>

{#if onwebsite}
	{#if !not}
		{@render children?.()}
	{/if}
{:else if not}
	{@render children?.()}
{/if}
