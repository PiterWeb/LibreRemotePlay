<script lang="ts">
	import type { BeforeNavigate } from '@sveltejs/kit';

	interface Props {
		children?: import('svelte').Snippet;
	}

	const { children }: Props = $props();
	
	let modal: HTMLDialogElement
	let modalNavigator = $state<BeforeNavigate>()
	
	export function openModal(navigator ?:  BeforeNavigate) {
      modal?.showModal?.();
      modalNavigator = navigator
    }
	
    export function getNavigator() {
      return modalNavigator
    }
    
</script>

<!-- Open the modal using ID.showModal() method -->
<dialog bind:this={modal} class="modal modal-bottom sm:modal-middle">
	<div class="modal-box bg-white text-gray-900">
		<form method="dialog" class="w-full mx-auto flex flex-col justify-center gap-4">
            <!-- if there is a button in form, it will close the modal -->
            {@render children?.()}
		</form>
	</div>
</dialog>
