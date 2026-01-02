export namespace audio {
	
	export class AudioProcess {
	    Name: string;
	    Pid: number;
	
	    static createFrom(source: any = {}) {
	        return new AudioProcess(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Pid = source["Pid"];
	    }
	}

}

export namespace webrtc {
	
	export class ICEServer {
	    urls: string[];
	    username?: string;
	    credential?: any;
	    credentialType?: number;
	
	    static createFrom(source: any = {}) {
	        return new ICEServer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.urls = source["urls"];
	        this.username = source["username"];
	        this.credential = source["credential"];
	        this.credentialType = source["credentialType"];
	    }
	}

}

