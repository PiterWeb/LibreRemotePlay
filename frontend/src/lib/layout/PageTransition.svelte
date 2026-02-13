<script lang="ts">
	import { playAudio } from '$lib/audio/audio_player';

	interface Props {
		key: string;
		children?: import('svelte').Snippet;
	}

	let { key, children }: Props = $props();

	$effect(() => {
		key && playAudio('page_transition');
	});
</script>

{#key key}
		{@render children?.()}
{/key}
