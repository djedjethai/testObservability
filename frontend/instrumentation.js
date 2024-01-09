export async function register() {
	if (process.env.NEXT_RUNTIME === 'nodejs') {
		console.log("is process.env.NEXT_RUNTIME")
    		await import('./instrumentation.node.js')
  	}
}

