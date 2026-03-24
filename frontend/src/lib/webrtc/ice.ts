
export interface ServersConfig {
  [group: string]: {
    enable: boolean;
    server: ICEServer;
  };
}

export interface ICEServer {
	urls: string[];
	username?: string;
	credential?: string;
}