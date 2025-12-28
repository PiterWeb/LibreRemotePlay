<script lang="ts">
	import media_css from "$lib/css/media-video.css?raw"
	import 'player.style/microvideo';
	import { consumingStream } from '$lib/webrtc/stream/stream_signal_hook.svelte';
	import log from "$lib/logger/logger";

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
	
	let videoElement: HTMLVideoElement | null = $state(null)
	
	const onContextMenu = (ev: PointerEvent) => {
	  ev.preventDefault()
	  return false
	}
	
	$effect(() => {
	    if (!videoElement) return
				
		try {
    		// Disable pause on video
    		// TODO: Look for better aproaches to prevent video pause
    		videoElement.pause = () => {}
		} catch(e) {
		    log(e, {err: true})
		}
		
		// Disable right-click context menu
        videoElement.addEventListener("contextmenu", onContextMenu)
		
		return () => {
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