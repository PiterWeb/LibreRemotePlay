import { LogPrintln } from "$lib/bindings/App";


export default async function log(info: unknown, { err } = { err: false }) {
	try {
		await LogPrintln(`Browser - ${JSON.stringify(info)}\n`);
  } catch {/* */} finally {
		if (err) console.error(err);
		console.trace(info);
	}
}
