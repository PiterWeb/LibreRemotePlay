<script lang="ts">
	import { isDesktop } from "./detect_os";
	import type { Snippet } from 'svelte';
	let { children, not = false }: { children?: Snippet; not?: boolean } = $props();
</script>

{#if isDesktop}
	{#if !not}
		{@render children?.()}
	{/if}
{:else if not}
	{@render children?.()}
{/if}