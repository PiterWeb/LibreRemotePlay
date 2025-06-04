export default async function log(info: string) {

    try {
        // const { LogPrint } = await import("$lib/wailsjs/runtime/runtime") 
        // LogPrint(info + '\n')        
    } finally {
        console.log(info)
    }

}