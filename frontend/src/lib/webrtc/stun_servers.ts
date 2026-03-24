import { get, writable } from 'svelte/store';
import type { ServersConfig, ICEServer } from '$lib/webrtc/ice';
import log from '$lib/logger/logger';

const defaultStunServers = [
	'stun:stun.l.google.com:19302',
	'stun:stun.ipfire.org:3478',
	'stun:stun.l.google.com:19305'
];

const defaultStunConfig: Readonly<ServersConfig> = {
  default: {
    server: {
      urls: defaultStunServers
    },
    enable: true
	}
};

const stunServersStore = writable<ServersConfig>(
	JSON.parse(localStorage.getItem('stunServers') ?? 'false') || defaultStunConfig
);

stunServersStore.subscribe((stunServers) =>
	localStorage.setItem('stunServers', JSON.stringify(stunServers))
);

function removeServerFromGroup(group: string, url: string) {
	stunServersStore.update((stunServers) => {
		stunServers[group].server.urls = stunServers[group].server.urls.filter((server) => server !== url);
		return stunServers;
	});
}

function modifyGroup(name: string, enable?:boolean, newName?: string, username?: string, credential?: string) {

	log({name, enable, newName, username, credential});
	if (newName) {
		stunServersStore.update((stunServers) => {
			stunServers[newName] = stunServers[name];
			if (username) stunServers[newName].server.username = username;
      if (credential) stunServers[newName].server.credential = credential;
			if (enable !== undefined) stunServers[newName].enable = enable
			delete stunServers[name];
			return stunServers;
		});

		return;
	}

	stunServersStore.update((stunServers) => {
		stunServers[name].server.username = username;
    stunServers[name].server.credential = credential;
		if (enable !== undefined) stunServers[name].enable = enable
		return stunServers;
	});
}

function addServerToGroup(group: string, url: string) {
	stunServersStore.update((stunServers) => {
		stunServers[group].server.urls.push('stun:' + url);
		return stunServers;
	});
}

function createServerGroup(name: string, username?: string, credential?: string) {
	const newServer: ServersConfig = {
    [name]: {
      server: {    
   			urls: [],
   			username: username,
   			credential: credential
      },
      enable: true
		}
	};
	stunServersStore.update((stunServers) => {
		return {
			...stunServers,
			...newServer
		};
	});
}

function deleteServerGroup(name: string) {
	stunServersStore.update((stunServers) => {
		delete stunServers[name];
		return stunServers;
	});
}

function exportStunServers(): ICEServer[] {
	const servers = get(stunServersStore);
  const serversArray = Object.keys(servers).filter((key) => {
    return servers[key].enable
  }).map((key) => {
		return {
			urls: servers[key].server.urls,
			...(servers[key].server.username && { username: servers[key].server.username }),
			...(servers[key].server.credential && { credential: servers[key].server.credential })
		};
	});

	return serversArray;
}

export {
	stunServersStore,
	addServerToGroup,
	removeServerFromGroup,
	createServerGroup,
	deleteServerGroup,
	modifyGroup,
	exportStunServers,
	defaultStunConfig
};
