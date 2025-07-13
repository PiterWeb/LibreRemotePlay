import * as MockRTC from 'mockrtc';
import * as fs from 'fs';

const adminServer = MockRTC.getAdminServer();
adminServer.start().then(() => console.log('WebRTC Admin server started'));

const mockRTC = MockRTC.getRemote({ recordMessages: true, debug: true });

const signalWasmBuffer = fs.readFileSync('../../../static/wasm/signal.wasm');

describe('Basic connection', async () => {
	const wasmModule = await WebAssembly.instantiate(signalWasmBuffer);

	const signalEncode = wasmModule.instance.exports.signalEncode as <T>(signal: T) => string;
	const signalDecode = wasmModule.instance.exports.signalDecode as <T>(signal: string) => T;

	beforeEach(() => mockRTC.start());
    afterEach(() => mockRTC.stop());

	it('load', async () => {
		cy.visit('http://localhost:34115/').wait(1000).log('hello');

		const mockPeer = await mockRTC.buildPeer().thenEcho();

		const { offer: mockOffer, setAnswer } = await mockPeer.createOffer();

		const mockOfferEncoded = signalEncode(mockOffer)

		// const answerEncoded = ""

		// const answer: RTCSessionDescriptionInit = signalDecode(answerEncoded)

		// await setAnswer(answer)

		// Start WebRTC connection from the CypressUI
	});
});
