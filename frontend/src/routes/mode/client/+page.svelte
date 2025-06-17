<script>
	import { CreateClientWeb, ConnectToHostWeb } from '$lib/webrtc/client_webrtc_hook';

	import {
		easyConnectServerIpDomain,
		easyConnectID,
		handleEasyConnectClient,
		DEFAULT_EASY_CONNECT_ID,
		DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN
	} from '$lib/easy_connect/easy_connect.svelte';
	import { showToast, ToastType } from '$lib/toast/toast_hook';
	import { _ } from 'svelte-i18n';
	import { onMount } from 'svelte';
	import LANMode from '$lib/webrtc/lan_mode.svelte';

	let code = $state('');
	let clientCreated = $state(false);

	onMount(() => {
		easyConnectID.value = DEFAULT_EASY_CONNECT_ID;
		easyConnectServerIpDomain.value = DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN;
	});

	function handleConnectToHost() {
		if (code.length < 1) {
			showToast($_('code-is-empty'), ToastType.ERROR);
			return;
		}

		ConnectToHostWeb(code);
	}

	async function handleCreateClient() {
		await CreateClientWeb({ easyConnect: false });
		clientCreated = true;
	}
</script>

<h2 class="text-center text-[clamp(2rem,6vw,4.2rem)] font-black leading-[1.1] xl:text-left">
	<span
		class="[&amp;::selection]:text-base-content text-transparent relative col-start-1 row-start-1 bg-clip-text bg-gradient-to-r from-blue-700 via-blue-800 to-gray-900"
		>{$_('client_card_title')}
	</span>
</h2>

<div
	class="mt-12 md:mt-4 card bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"
>
	<div class="card-body">
		<section class="flex flex-row-reverse items-center gap-2">
			<label for="lan-mode-checkbox" class="font-semibold text-gray-900 dark:text-white"
				>{$_('lan-mode')}</label
			>
			<input
				id="lan-mode-checkbox"
				type="checkbox"
				class="checkbox checkbox-xs checkbox-primary"
				bind:checked={LANMode.enabled}
			/>
		</section>
		<section class="flex flex-col md:flex-row [&>ol]:px-4">
			<ol>
				<li class="mb-10 me-4">
					<h3 class="text-2xl font-semibold text-gray-900 dark:text-white">
						{$_('easy-connect')}
					</h3>
					<p class="text-sm text-gray-400 dark:text-gray-500 mb-4">
						{$_('easy-connect-description')}
					</p>

					<section class="flex flex-col gap-4">
						<div>
							<label
								for="ip-domain-easy-connect"
								class="mb-1 text-md font-normal leading-none text-gray-400 dark:text-gray-500"
								>{$_('ip-domain-easy-connect')}</label
							>

							<input
								id="ip-domain-easy-connect"
								type="text"
								placeholder={$_('ip-domain-easy-connect')}
								class="input input-bordered w-full"
								bind:value={easyConnectServerIpDomain.value}
							/>
						</div>

						<div class="flex flex-col">
							<label
								for="id-easy-connect"
								class="mb-1 text-md font-normal leading-none text-gray-400 dark:text-gray-500"
								>{$_('id-easy-connect')}</label
							>

							<input
								id="id-easy-connect"
								type="number"
								placeholder={$_('id-easy-connect')}
								class="input input-bordered w-28"
								min="0000"
								max="9999"
								bind:value={easyConnectID.value}
							/>
						</div>

						<button id="connect-to-host" onclick={handleEasyConnectClient} class="btn btn-primary"
							>{$_('connect-to-host')}</button
						>
					</section>
				</li>
				<li class="block md:hidden">
					<hr class="border-gray-200 dark:border-gray-700" />
				</li>
			</ol>
			<ol>
				<li>
					<h3 class="text-2xl font-semibold text-gray-900 dark:text-white">
						{$_('manual-connection')}
					</h3>
					<p class="text-sm text-gray-400 dark:text-gray-500 mb-4">
						{$_('manual-connection-description')}
					</p>
				</li>
				<li>
					<ol class="relative border-s border-gray-200 dark:border-gray-700">
						<li class="mb-10 ms-4">
							{#if !clientCreated}
								<div
									class="absolute w-3 h-3 bg-gray-200 rounded-full mt-1.5 -start-1.5 border border-white dark:border-gray-900 dark:bg-gray-700"
								></div>
							{:else}
								<span
									class="absolute flex items-center justify-center rounded-full mt-1.5 w-3 h-3 -start-1.5 border border-white dark:border-gray-900"
								>
									<svg
										class="w-2.5 h-2.5 text-green-500 dark:text-green-400 flex-shrink-0"
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
								class="mb-1 text-sm font-normal leading-none text-gray-400 dark:text-gray-500"
								>{$_('first-step')}</label
							>
							<h3 class="text-lg font-semibold text-gray-900 dark:text-white">
								{$_('share-the-code-with-your-host')}
							</h3>
							{#if clientCreated}
								<p class="text-gray-400 dark:text-gray-500 mb-4">
									{$_('if-your-code-is-missing-from-your-clipboard-you-must-restart-the-process')}
								</p>
							{/if}

							{#if !clientCreated}
								<button id="create-client" onclick={handleCreateClient} class="btn btn-primary"
									>{$_('client_card_cta')}</button
								>
							{/if}
						</li>
						<li class="mb-10 ms-4">
							<div
								class="absolute w-3 h-3 bg-gray-200 rounded-full mt-1.5 -start-1.5 border border-white dark:border-gray-900 dark:bg-gray-700"
							></div>
							<label
								for="connect-to-host"
								class="mb-1 text-sm font-normal leading-none text-gray-400 dark:text-gray-500"
								>{$_('second-step')}</label
							>
							<h3 class="text-lg font-semibold text-gray-900 dark:text-white">
								{$_('get-the-code-from-your-host')}
							</h3>
							<div class="join">
								<input
									type="text"
									placeholder="Paste here code"
									class="input input-bordered w-full max-w-xs"
									bind:value={code}
								/>
								<button id="connect-to-host" onclick={handleConnectToHost} class="btn btn-primary"
									>{$_('connect-to-host')}</button
								>
							</div>
						</li>
					</ol>
				</li>
			</ol>
		</section>
	</div>
</div>
