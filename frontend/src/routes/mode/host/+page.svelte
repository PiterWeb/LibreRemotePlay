<script>
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

	let code = $state('');
	let generatedCode = $state(false);

	onMount(() => {
		easyConnectID.value = DEFAULT_EASY_CONNECT_ID;
		easyConnectServerIpDomain.value = DEFAULT_EASY_CONNECT_SERVER_IP_DOMAIN;
	});

	async function handleConnectToClient() {
		if (code.length < 1) {
			showToast($_('code-is-empty'), ToastType.ERROR);
			return;
		}

		try {
			await CreateHost({ clientCode: code, easyConnect: false });
			generatedCode = true;
		} catch {
			generatedCode = false;
		}
	}
</script>

<h2 class="text-center text-[clamp(2rem,6vw,4.2rem)] font-black leading-[1.1] xl:text-left">
	<span
		class="[&amp;::selection]:text-base-content text-transparent relative col-start-1 row-start-1 bg-clip-text bg-gradient-to-r from-blue-700 via-blue-800 to-gray-900"
		>{$_('host_card_title')}
	</span>
</h2>

<div
	class="mt-12 md:mt-4 card bg-white border border-gray-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700"
>
	<div class="card-body">
		<ConnectionOptions/>
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

						<button id="connect-to-client" onclick={handleEasyConnectHost} class="btn btn-primary"
							>{$_('connect-to-client')}</button
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
							{#if !generatedCode}
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
								{$_('get-your-host-code')}
							</h3>
							<div class="join">
								<input
									disabled={generatedCode}
									type="text"
									id="first_name"
									class="max-w-xs bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-l-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
									placeholder={$_('paste-here-code')}
									required
									bind:value={code}
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
								class="absolute w-3 h-3 bg-gray-200 rounded-full mt-1.5 -start-1.5 border border-white dark:border-gray-900 dark:bg-gray-700"
							></div>
							<label
								for="connect-to-host"
								class="mb-1 text-sm font-normal leading-none text-gray-400 dark:text-gray-500"
								>{$_('second-step')}</label
							>
							<h3 class="text-lg font-semibold text-gray-900 dark:text-white">
								{$_('share-the-code-with-your-client')}
							</h3>
						</li>
						<div class="card-actions gap-4 justify-end items-center flex-col w-full"></div>
					</ol>
				</li>
			</ol>
		</section>
	</div>
</div>
