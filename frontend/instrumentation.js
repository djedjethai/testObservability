export async function register() {
	
	if (process.env.NEXT_RUNTIME === 'nodejs') {
		console.log("is registerrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr")
    		await import('./instrumentation.node.js')
  	}
}
