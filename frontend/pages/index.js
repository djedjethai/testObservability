import Router from 'next/router'
import { useState, useEffect } from 'react'

import { trace, context, propagation } from '@opentelemetry/api'


const Home =  ({pics}) => {

	const [numbers, setNumbers] = useState('')
	const [method, setMethod] = useState('')
	const [errors, setErrors] = useState(null)

		
	const rendPics =  pics =>{
		console.log("The pics: ", pics)
	}

	const onSubmit = async(event) => {
		event.preventDefault()


		const operands = numbers.split(",").map(Number);
		console.log("sned request, operands: ", operands)
		console.log("sned request, method: ", method)

		// NOTE is the headers are really needed ??
		const headers = {};
		propagation.inject(context.active(), headers);

		await trace
			.getTracer("nextJs example")
			.startActiveSpan('calculate', async(span) => {
				try{
					const response = await fetch('http://127.0.0.1:4000/calculate', {
        					method: 'POST',
        					headers: {
        						'Content-Type': 'application/json',
							...headers, // Add the injected headers
        					},
						body: JSON.stringify({ method , operands }),
      					});

					// console.log("See the response: ", response)
					const responseData = await response.json();
  					console.log('Response data:', responseData);
				} catch(e) {
					console.log(e)
				} finally {
					span.end()
				}
			})
	}




  	return (
	  		<div>
				{rendPics(pics)}
	  			the main page uu............
				<form onSubmit={onSubmit}>
					<h1>Sign in</h1>
					<div className="form-group">
						<label>Numbers</label>
						<input 
							value={numbers} 
							onChange={e => setNumbers(e.target.value)}
							className="form-control" 
						/>
					</div>
					<div className="form-group">
						<label>Operation</label>
						<input 
							value={method}
							onChange={e => setMethod(e.target.value)}
							className="form-control" 
						/>
					</div>
					{/*noneed to check if err or not as it default to null
					see the useRequest hook*/}
					{errors}
					<button className="btn btn-primary">Sign In</button>
				</form>

	  		</div>
        )
}

// Home.getInitialProps = async (appContext, client, currentUser) => {
Home.getInitialProps = async (appContext, currentUser) => {
	// const client = buildClient(appContext.ctx)
	
	// const response = await fetch('http://localhost:9096');

 	// domain "apiclient:4000" as apiclient and nextjs are in a bridge network
	// localhost from a container belong to the container, can t be use between container
	
	console.log("the index getInitialProps, currentUser: ", currentUser)

	// const { setUrlAndCreateCodeChallenge } = useOauth2()

	// setUrlAndCreateCodeChallenge();

	// need to create an arr to then destruct
 	// otherwise the datas get parsed
 	return { pics: "grrrrrr" }
}


export default Home

// // ================
// const Home =  ({pics}) => {
// 
// 	const [numbers, setNumbers] = useState('')
// 	const [method, setMethod] = useState('')
// 	const [errors, setErrors] = useState(null)
// 
// 		
// 	const rendPics =  pics =>{
// 		console.log("The pics: ", pics)
// 	}
// 
// 	const onSubmit = async(event) => {
// 		event.preventDefault()
// 
// 
// 		const operands = numbers.split(",").map(Number);
// 		console.log("sned request, operands: ", operands)
// 		console.log("sned request, method: ", method)
// 
// 		try{
// 			const response = await fetch('http://127.0.0.1:4000/calculate', {
//         			method: 'POST',
//         			headers: {
//         				'Content-Type': 'application/json',
//         			},
// 				body: JSON.stringify({ method , operands }),
//       			});
// 
// 			// console.log("See the response: ", response)
// 			const responseData = await response.json();
//   			console.log('Response data:', responseData);
// 		} catch(e) {
// 			console.log(e)
// 		}
// 
// 
// 
// 		// const e = await runOauthOpenID(email, password, 'signin')
// 		// if(e){
// 		// 	// TODO refactor that
// 		// 	setErrors(<div className="alert alert-danger">
// 		// 		<h4>Oooops ...</h4>
// 		// 		<ul className="my-0">
// 		// 			<li key={e.message}>{e.message}</li>
// 		// 		</ul>
// 		// 	</div>)
// 		// } else {
// 		// 	Router.push('/')
// 		// }
// 	}
// 
// 
// 
// 
//   	return (
// 	  		<div>
// 				{rendPics(pics)}
// 	  			the main page uu............
// 				<form onSubmit={onSubmit}>
// 					<h1>Sign in</h1>
// 					<div className="form-group">
// 						<label>Numbers</label>
// 						<input 
// 							value={numbers} 
// 							onChange={e => setNumbers(e.target.value)}
// 							className="form-control" 
// 						/>
// 					</div>
// 					<div className="form-group">
// 						<label>Operation</label>
// 						<input 
// 							value={method}
// 							onChange={e => setMethod(e.target.value)}
// 							className="form-control" 
// 						/>
// 					</div>
// 					{/*noneed to check if err or not as it default to null
// 					see the useRequest hook*/}
// 					{errors}
// 					<button className="btn btn-primary">Sign In</button>
// 				</form>
// 
// 	  		</div>
//         )
// }
// 
// // Home.getInitialProps = async (appContext, client, currentUser) => {
// Home.getInitialProps = async (appContext, currentUser) => {
// 	// const client = buildClient(appContext.ctx)
// 	
// 	// const response = await fetch('http://localhost:9096');
// 
//  	// domain "apiclient:4000" as apiclient and nextjs are in a bridge network
// 	// localhost from a container belong to the container, can t be use between container
// 	
// 	console.log("the index getInitialProps, currentUser: ", currentUser)
// 
// 	// const { setUrlAndCreateCodeChallenge } = useOauth2()
// 
// 	// setUrlAndCreateCodeChallenge();
// 
// 	// need to create an arr to then destruct
//  	// otherwise the datas get parsed
//  	return { pics: "grrrrrr" }
// }
// 
// 
// export default Home
// 
