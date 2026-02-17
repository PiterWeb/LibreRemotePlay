<script lang="ts">
	import { CreateHost } from '$lib/webrtc/host_webrtc_hook';
	import { showToast, ToastType } from '$lib/toast/toast_hook';
	import { _ } from 'svelte-i18n';
	import {
		DEFAULT_EASY_CONNECT_ID,
		DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN,
		easyConnectID,
		easyConnectServerIpDomain,
		handleEasyConnectHost
	} from '$lib/easy_connect/easy_connect.svelte';
	import { onMount } from 'svelte';
	import ConnectionOptions from '$lib/webrtc/ConnectionOptions.svelte';
	import Modal from '$lib/layout/Modal.svelte';
	import EasyConnect from '../EasyConnect.svelte';

	let modalElement: Modal
	
	let hostCode = $state('')
	let clientCode = $state('');
	let generatedCode = $state(false);

	onMount(() => {
		easyConnectID.value = DEFAULT_EASY_CONNECT_ID;
		easyConnectServerIpDomain.value = DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN;
	});

	async function handleConnectToClient() {
		if (clientCode.length < 1) {
			showToast($_('code-is-empty'), ToastType.ERROR);
			return;
		}

		try {
			const code = await CreateHost({ clientCode, easyConnect: false });
			hostCode = code ?? ""
			generatedCode = true;
			modalElement.openModal()
		} catch {
			generatedCode = false;
		}
	}

	async function copyCodeToClipboard() {
		try {
			await navigator.clipboard.writeText(hostCode)
			showToast($_('client-code-copied-to-clipboard'), ToastType.SUCCESS);
		} catch {
			showToast($_('error-copying-client-code-to-clipboard'), ToastType.ERROR);
		}
	}

</script>

<Modal bind:this={modalElement}>
	<p class="text-xl">Did lost your code?</p>
	<div class="flex gap-2">
		<button onclick={copyCodeToClipboard} class="btn btn-primary">Copy to clipboard</button>
		<button class="btn btn-neutral">Close</button>
	</div>
</Modal>

<h2 class="text-center text-white text-[clamp(2rem,6vw,4.2rem)] font-black leading-[1.1] xl:text-left">
		{$_('host_card_title')}
</h2>

<div class="mt-8 md:mt-0 card rounded-lg shadow-xl border border-gray-800 w-full p-2 mb-0">
    <section
    	class="card-body rounded-lg shadow-xl bg-white p-4"
    >           
          		<ConnectionOptions/>
    </section>
</div>

<div class="mt-0 grid md:grid-cols-5 gap-4 card rounded-lg shadow-xl border border-gray-800 p-2">
    <EasyConnect role="HOST"/>
    <section class="md:col-span-2 md:pt-8 card border rounded-lg shadow-xl bg-white border-gray-700">
        <div class="card-body">
            <section class="flex flex-col md:flex-row [&>ol]:px-4">
                <ol>
				<li>
					<h3 class="text-2xl font-semibold text-gray-800">
						{$_('manual-connection')}
					</h3>
					<p class="text-sm text-gray-500 mb-4">
						{$_('manual-connection-description')}
					</p>
				</li>
				<li>
					<ol class="relative border-s border-gray-700">
						<li class="mb-10 ms-4">
							{#if !generatedCode}
								<div
									class="absolute w-3 h-3 rounded-full mt-1.5 -start-1.5 border border-gray-900 bg-gray-700"
								></div>
							{:else}
								<span
									class="absolute flex items-center justify-center rounded-full mt-1.5 w-3 h-3 -start-1.5 border border-gray-900"
								>
									<svg
										class="w-2.5 h-2.5 text-green-400 flex-shrink-0"
										aria-hidden="true"
										xmlns="http://www.w3.org/2000/svg"
										fill="currentColor"
										viewBox="0 0 20 20"
									>
										<path
											d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5Zm3.707 8.207-4 4a1 1 0 0 1-1.414 0l-2-2a1 1 0 0 1 1.414-1.414L9 10.586l3.293-3.293a1 1 0 0 1 1.414 1.414Z"
										/>
									</svg>
								</span>
							{/if}
							<label
								for="create-client"
								class="mb-1 text-sm font-normal leading-none text-gray-500"
								>{$_('first-step')}</label
							>
             
							<h3 class="text-lg font-semibold text-gray-600">
								{$_('get-your-host-code')}
							</h3>
							<div class="join">
								<input
									disabled={generatedCode}
									type="text"
									id="first_step"
									class="input input-bordered w-full max-w-xs"
									placeholder={$_('paste-here-code')}
									required
									bind:value={clientCode}
								/>
								<button
									disabled={generatedCode}
									onclick={handleConnectToClient}
									class="btn btn-primary">{$_('connect-to-client')}</button
								>
							</div>
						</li>
						<li class="mb-10 ms-4">
							<div
								class="absolute w-3 h-3 rounded-full mt-1.5 -start-1.5 border border-gray-900 bg-gray-700"
							></div>
							<label
								for="connect-to-host"
								class="mb-1 text-sm font-normal leading-none text-gray-500"
								>{$_('second-step')}</label
							>
							<h3 class="text-lg font-semibold text-gray-600">
								{$_('share-the-code-with-your-client')}
							</h3>
						</li>
						<div class="card-actions gap-4 justify-end items-center flex-col w-full"></div>
					</ol>
				</li>
			</ol>
            </section>
    	</div>
    </section>
</div>