<script lang="ts">
	import {
		easyConnectServerIpDomain,
		easyConnectID,
		handleEasyConnectClient,
		handleEasyConnectHost,
		DEFAULT_EASY_CONNECT_ID,
		DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN
	} from '$lib/easy_connect/easy_connect.svelte';
	
	import { _ } from 'svelte-i18n';
	import { onMount } from 'svelte';

	const { role }: {role: "CLIENT" | "HOST"} = $props()
	
	onMount(() => {
		easyConnectID.value = DEFAULT_EASY_CONNECT_ID;
		easyConnectServerIpDomain.value = DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN;
	});
	
</script>

    <section
    	class="md:col-span-3 md:pt-8 card border rounded-lg shadow-xl bg-white border-gray-700"
    >
    	<div class="card-body">
    		<section class="flex flex-col md:flex-row [&>ol]:px-4">
    			<ol>
    				<li class="mb-10">
    					<h3 class="text-2xl font-semibold text-gray-800">
    						{$_('easy-connect')}
    					</h3>
    					<p class="text-sm text-gray-500 mb-4">
    						{$_('easy-connect-description')}
    					</p>
    
    					<section class="flex flex-col gap-4">
    						<div>
    							<label
    								for="ip-domain-easy-connect"
    								class="mb-1 text-lg font-normal leading-none text-gray-600"
    								>{$_('ip-domain-easy-connect')}</label
    							>
    
    							<input
    								id="ip-domain-easy-connect"
    								type="text"
    								placeholder={$_('ip-domain-easy-connect')}
    								class="input input-primary w-full"
    								bind:value={easyConnectServerIpDomain.value}
    							/>
    						</div>
    
    						<div class="flex flex-col">
    							<label
    								for="id-easy-connect"
    								class="mb-1 text-lg font-normal leading-none text-gray-600"
    								>{$_('id-easy-connect')}</label
    							>
    
    							<input
    								id="id-easy-connect"
    								type="number"
    								placeholder={$_('id-easy-connect')}
    								class="input input-primary w-28"
    								min="0000"
    								max="9999"
    								bind:value={easyConnectID.value}
    							/>
    						</div>
    
                            {#if role === "CLIENT"}
                               	<button onclick={handleEasyConnectClient} class="btn btn-primary"
                          							>{$_('connect-to-host')}</button>
                            {:else}
                                <button onclick={handleEasyConnectHost} class="btn btn-primary"
                         							>{$_('connect-to-client')}</button>
                            {/if}
          
    					
    					</section>
    				</li>
    			</ol>
    			
    		</section>
    	</div>
    </section>
