<script lang="ts" module>
	let modal: HTMLDialogElement;

	export function openModal() {
		modal?.showModal?.();
	}
</script>
<script lang="ts">
	interface Props {
		children?: import('svelte').Snippet;
	}

	const { children }: Props = $props();
</script>

<!-- Open the modal using ID.showModal() method -->
<dialog bind:this={modal} class="modal modal-bottom sm:modal-middle">
	<div class="modal-box">
		<form method="dialog" class="w-full mx-auto flex flex-col justify-center gap-4">
            <!-- if there is a button in form, it will close the modal -->
            {@render children?.()}
		</form>
	</div>
</dialog>
