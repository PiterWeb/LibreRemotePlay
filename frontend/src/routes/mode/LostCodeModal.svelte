<script lang="ts">
	import Modal from "$lib/layout/Modal.svelte";
	
	interface Props {
	  lostCodeModalElement: Modal
	  copyCodeToClipboard(): Promise<void>
	}
	
	let { lostCodeModalElement = $bindable(), copyCodeToClipboard }: Props = $props()

</script>
<Modal bind:this={lostCodeModalElement}>
	<p class="text-xl">Did lost your code?</p>
	<div class="flex gap-2">
		<button onclick={copyCodeToClipboard} class="btn btn-primary">Copy to clipboard</button>
		<button class="btn btn-neutral">Close</button>
	</div>
</Modal>