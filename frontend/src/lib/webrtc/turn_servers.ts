import { get, writable } from 'svelte/store';
import type { ServersConfig, ICEServer } from '$lib/webrtc/ice';

const defaultTurnConfig: Readonly<ServersConfig> = {

};

const turnServersStore = writable<ServersConfig>(
	JSON.parse(localStorage.getItem('turnServers') ?? 'false') || defaultTurnConfig
);

turnServersStore.subscribe((turnServers) =>
	localStorage.setItem('turnServers', JSON.stringify(turnServers))
);

function removeServerFromGroup(group: string, url: string) {
	turnServersStore.update((turnServers) => {
    turnServers[group].server.urls = turnServers[group].server.urls.filter((server) => server !== url);
		return turnServers;
	});
}

function modifyGroup(name: string, enable?: boolean, newName?: string, username?: string, credential?: string) {
	if (newName) {
		turnServersStore.update((turnServers) => {
			turnServers[newName] = turnServers[name];
			turnServers[newName].server.username = username;
			turnServers[newName].server.credential = credential;
			if (enable !== undefined) turnServers[newName].enable = enable
			delete turnServers[name];
			return turnServers;
		});

		return;
	}

	turnServersStore.update((turnServers) => {
		turnServers[name].server.username = username;
    turnServers[name].server.credential = credential;
    if (enable !== undefined) turnServers[name].enable = enable
		return turnServers;
	});
}

function addServerToGroup(group: string, url: string) {
	turnServersStore.update((turnServers) => {
		turnServers[group].server.urls.push('turn:' + url);
		return turnServers;
	});
}

function createServerGroup(name: string, username?: string, credential?: string) {
  const newServer: ServersConfig = {
    [name]: { 
      server :{
        urls: [],
        username: username,
        credential: credential
      },
      enable: true
    }
  };
  turnServersStore.update((turnServers) => {
    return {
      ...turnServers,
			...newServer
		};
	});
}

function deleteServerGroup(name: string) {
	turnServersStore.update((turnServers) => {
		delete turnServers[name];
		return turnServers;
	});
}

function exportTurnServers(): ICEServer[] {
	const servers = get(turnServersStore);
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
	turnServersStore,
	addServerToGroup,
	removeServerFromGroup,
	createServerGroup,
	deleteServerGroup,
	modifyGroup,
	exportTurnServers,
	defaultTurnConfig
};
