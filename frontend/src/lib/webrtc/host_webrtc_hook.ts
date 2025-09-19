import {
	TryCreateHost as createHostFn,
	TryClosePeerConnection as closeConnectionFn
} from '$lib/wailsjs/go/bindings/App';

import { _ } from 'svelte-i18n';
import { get } from 'svelte/store';
import { showToast, ToastType } from '$lib/toast/toast_hook';
import { goto } from '$app/navigation';
import { toogleLoading, setLoadingMessage, setLoadingTitle } from '$lib/loading/loading_hook';
import { StopStreaming } from '$lib/webrtc/stream/host_stream_hook';
import type { ICEServer } from '$lib/webrtc/ice';
import { exportStunServers } from './stun_servers';
import { exportTurnServers } from './turn_servers';
import { isLinux } from '$lib/detection/detect_os';
import { IS_RUNNING_EXTERNAL } from '$lib/detection/onwebsite';
import log from '$lib/logger/logger';
import LANMode from './lan_mode.svelte';

const BROWSER_BASE_URL = 'http://localhost:8080/mode/host/connection';

let host: boolean = false;

enum ConnectionState {
	Connected = 'CONNECTED',
	Failed = 'FAILED',
	Disconnected = 'DISCONNECTED'
}

interface CreateHostOptions {
	easyConnect: boolean;
	clientCode: string;
}

export async function CreateHost(options: CreateHostOptions) {
	try {
		const { clientCode: client, easyConnect } = options;

		const ICEServers: ICEServer[] = LANMode.enabled ? [] : [...exportStunServers(), ...exportTurnServers()];

		const hostCode = await createHostFn(ICEServers, client);

		if (isError(hostCode)) {
      throw hostCode;
		}

		if (!easyConnect && navigator && navigator.clipboard && navigator.clipboard.writeText) {
			navigator.clipboard.writeText(hostCode).catch(() => {
				showToast(get(_)('error-copying-host-code-to-clipboard'), ToastType.ERROR);
			});
			showToast(get(_)('host-code-copied-to-clipboard'), ToastType.SUCCESS);
		} else if (!easyConnect) {
			showToast(get(_)('error-copying-host-code-to-clipboard'), ToastType.ERROR);
			throw hostCode
		}

		toogleLoading();
		setLoadingMessage(get(_)('waiting-for-client-to-connect'));
		setLoadingTitle(get(_)('make-sure-to-pass-the-code-to-the-client'));

		const { EventsOnce } = await import('$lib/wailsjs/runtime/runtime');

		EventsOnce('connection_state', async (state: ConnectionState) => {
			toogleLoading();

			try {
  			switch (state.toUpperCase()) {
  				case ConnectionState.Connected:
  					showToast(get(_)('connected'), ToastType.SUCCESS);
  					host = true;
  					if (await isLinux()) {
              log("Connection stablished and isLinux")
  						const { BrowserOpenURL } = await import('$lib/wailsjs/runtime/runtime');
  						BrowserOpenURL(BROWSER_BASE_URL);
  					} else log("Connection stablished and not isLinux")
  					goto('/mode/host/connection');
  					break;
  				case ConnectionState.Failed:
  					showToast(get(_)('connection-failed'), ToastType.ERROR);
  					goto('/');
  					break;
  				default:
  					showToast(get(_)('unknown-connection-state'), ToastType.ERROR);
  					break;
  			}
			
			} catch(e) {
			  log(e, {err: true})
			}
		});
		
		return hostCode
	} catch (e) {
    log(e, {err: true})
		showToast(get(_)('error-creating-host'), ToastType.ERROR);
	}
}

function isError(err: string) {
	return err.toUpperCase().includes('ERROR');
}

//@ts-expect-error Override window object property to use in debug mode from console
window.CloseHostConnection = CloseHostConnection

export function CloseHostConnection(fn?: () => void) {
	if (!host) return;
	closeConnectionFn();
	if (fn) fn();
	host = false;
	StopStreaming();
}

export async function ListenForConnectionChanges() {
	if (IS_RUNNING_EXTERNAL) return;

	const { EventsOn, EventsOff } = await import('$lib/wailsjs/runtime/runtime');

	const connectionStateCancelEventListener = EventsOn(
		'connection_state',
		(state: ConnectionState) => {
			switch (state.toUpperCase()) {
				case ConnectionState.Connected:
					showToast(get(_)('connected'), ToastType.SUCCESS);
					host = true;
					goto('/mode/host/connection');
					break;
				case ConnectionState.Failed:
					showToast(get(_)('connection-failed'), ToastType.ERROR);
					goto('/');
					break;
				case ConnectionState.Disconnected:
					showToast(get(_)('connection-lost'), ToastType.ERROR);
					host = false;
					goto('/');
					connectionStateCancelEventListener();
					EventsOff('connection_state');
					break;
			}
		}
	);
}
