<script lang="ts">
	import media_css from "$lib/css/media-video.css?raw"
	import 'player.style/microvideo';
	import { consumingStream } from '$lib/webrtc/stream/stream_signal_hook.svelte';

    /** @type {{children?: import('svelte').Snippet}} */
    let {children} = $props()

	let media: HTMLElement | null = $state(null)

	// Apply custom styles to the media element
	$effect(() => {
		if (!media) return

		const shadowRoot = media.shadowRoot
		const styles = document.createElement('style')
		styles.textContent = media_css

		shadowRoot?.appendChild(styles)
	})

	// const events: (keyof HTMLVideoElementEventMap)[] = ["click", "auxclick"]
	
	let videoElement: HTMLVideoElement | null = $state(null)
	
	$effect(() => {
	    if (!videoElement) return
				
		const onContextMenu = (ev: Event) => {
		  ev.preventDefault()
		  return false
		}
		
		// Disable right-click context menu
		videoElement.addEventListener("contextmenu", onContextMenu)
				
		// const onClick = (ev: Event) => {
		//     if (!(ev instanceof PointerEvent)) return
		//     ev.preventDefault()
		//     console.log(`Click ${ev.type}: ${ev.button}`)
		// }
		
		// Intercept clicks
		// events.forEach(ev_name => videoElement?.addEventListener(ev_name, onClick))
		
		// Disable pause on video
		// TODO: Look for better aproaches to prevent video pause
		videoElement.pause = () => {}
		
		return () => {
          // events.forEach(_ => videoElement?.removeEventListener("click", onClick))
          videoElement?.removeEventListener("contextmenu", onContextMenu)
		}
	})
	
</script>

{@render children?.()}

<media-theme-microvideo bind:this={media} class:hidden={!consumingStream.value}>
<!-- svelte-ignore a11y_media_has_caption -->
<video
    bind:this={videoElement}
	slot="media"
	id="stream-video"
	playsinline
	>
</video>
</media-theme-microvideo>

<!-- <audio id="stream-audio" class:hidden={!consumingStream.value} muted={!consumingStream.value} controls playsinline></audio> -->