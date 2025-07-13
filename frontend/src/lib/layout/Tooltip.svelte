<script lang="ts">
  import {
  	FloatingArrow,
  	arrow,
  	autoUpdate,
  	flip,
  	offset,
  	useDismiss,
  	useFloating,
  	useHover,
  	useInteractions,
  	useRole,
		type UseFloatingOptions,
  } from "@skeletonlabs/floating-ui-svelte";
	import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  
  interface TooltipProps {
    children: import('svelte').Snippet,
    ref: string,
    placement?: UseFloatingOptions["placement"]
    dark?: boolean
  }
  
  let { children, ref, placement, dark = true }: TooltipProps = $props();
  
  // State
  let open = $state(false);
  let elemArrow: HTMLElement | null = $state(null);
  
  // Use Floating
  const floating = useFloating({
  	whileElementsMounted: autoUpdate,
  	get open() {
  		return open;
  	},
  	onOpenChange: (v) => {
  		open = v;
  	},
  	placement: placement ?? "top",
  	get middleware() {
  		return [offset(10), flip(), elemArrow && arrow({ element: elemArrow })];
  	},
  });
  
  onMount(() => {
    const elementForTooltip = document.getElementById(ref)
    const referenceProps = interactions.getReferenceProps()
    
    for (const [key, value] of Object.entries(referenceProps)) {
      if (key.startsWith('on') && typeof value === 'function') {
        // Events like onClick, onMouseEnter...
        const eventName = key.slice(2).toLowerCase();
        elementForTooltip?.addEventListener(eventName as keyof HTMLElementEventMap, value as EventListener);
      } else {
        // Atributes like aria-label, id, role, etc.
        elementForTooltip?.setAttribute(key, value as string);
      }
    }
    
    floating.elements.reference = document.getElementById(ref)
  })
  
  // Interactions
  const role = useRole(floating.context, { role: "tooltip" });
  const hover = useHover(floating.context, { move: false });
  const dismiss = useDismiss(floating.context);
  const interactions = useInteractions([role, hover, dismiss]);
</script>

<!-- Floating Element -->
{#if open}
	<div
		bind:this={floating.elements.floating}
		style={floating.floatingStyles}
		{...interactions.getFloatingProps()}
		class="floating popover-neutral"
		transition:fade={{ duration: 200 }}
	>
        <div class:bg-white={dark} class:text-gray-800={dark} class:bg-gray-800={!dark} class:text-white={!dark} class="border-2 rounded-lg p-4">
            {@render children?.()}
        </div>
		<FloatingArrow bind:ref={elemArrow} context={floating.context} fill="#575969" />
	</div>
{/if}