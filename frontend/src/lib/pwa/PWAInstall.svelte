<script lang="ts">
	import { _ } from "svelte-i18n";
	
	interface InstallEvent extends Event {
	  prompt(): Promise<{outcome: "accepted" | "dismissed"}>
	}
	
    let canInstall = $state(false)
    let installPrompt = $state<InstallEvent>();
    
    window.addEventListener("beforeinstallprompt", (event) => {
      event.preventDefault();
      installPrompt = event as InstallEvent;
      console.log("Can install")
      canInstall = true
    });
	
    async function installPWA() {
      try {
        
        if (!installPrompt) return
        
        const result = await installPrompt.prompt();
        console.log(`Install prompt was: ${result.outcome}`);
        installPrompt = undefined;
        if (result.outcome === "accepted") {
          canInstall = false
        }
      
      } catch {
        canInstall = true
      }
    }
    
    $inspect(canInstall)
</script>

{#if canInstall}
    <button aria-label="install pwa" class="btn btn-primary" onclick={installPWA}>{$_('install_pwa')}</button>
{/if}