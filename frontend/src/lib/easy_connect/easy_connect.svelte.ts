import { ConnectToHostWeb, CreateClientWeb } from "$lib/webrtc/client_webrtc_hook"
import {SendClientCode, server} from "@libreremoteplay/signals"
import { _ } from 'svelte-i18n'

export const easyConnectServerIpDomain = $state({value: "localhost:80"})
export const easyConnectID = $state({value: 0})

export async function handleEasyConnectClient() {

    const easyConnectServerUrl = (() => { 
        if (easyConnectServerIpDomain.value.length < 1) {
            return "";
        }
        return `ws://${easyConnectServerIpDomain.value}/ws`;
    })()

    console.log("Easy Connect Server URL:", easyConnectServerUrl);

    if (!URL.canParse(easyConnectServerUrl)) {
        throw new Error("Easy Connect Server URL is not set");
    }

    const serverInstance = server(easyConnectServerUrl)

    const clientCode = await CreateClientWeb()

    console.log("Client Code:", clientCode);

    if (!clientCode) {
        throw new Error("Failed to create client code");
    }

    const hostCode = await SendClientCode(serverInstance, {data: clientCode}, easyConnectID.value)

    console.log("Host Code:", hostCode);

    if (hostCode.data.length < 1) {
		throw new Error("Failed to send client code to host");
	}

    await ConnectToHostWeb(hostCode.data)

} 