// import 'bootstrap/dist/css/bootstrap.css'
// import { Buffer } from 'buffer';
// import axios from 'axios';
import Layout from '../components/Layout'
import Header from '../components/header'
// import buildClient from '../services/build-client'
// import { authRoutes } from '../services/config'

import { register } from '../instrumentation';
import { trace, context, propagation } from '@opentelemetry/api';

const AppComponent = ({ Component, pageProps, currentUser }) => {

	return (
		<Layout currentUser={currentUser}>
			<Header currentUser={ currentUser } />
			<div className="container">
				<Component currentUser={ currentUser } {...pageProps} />
			</div>
		</Layout>
	)
}

export default AppComponent


AppComponent.getInitialProps = async (appContext) => {

	// test OpenTelemetry from the server side

	// NOTE is the headers are really needed ??
	const headers = {};
	propagation.inject(context.active(), headers);

	await trace
		.getTracer("nextJs example")
		.startActiveSpan('calculate', async(span) => {
			try{
				const response = await fetch('http://127.0.0.1:4000/', {
        				method: 'GET',
        				headers: {
        					'Content-Type': 'application/json',
						...headers, // Add the injected headers
        				},
      				});

				const data = await response.text();
    				console.log('Raw response:', data);

				// console.log("See the response: ", response)
				// const responseData = await response.json();
  				// console.log('Response data:', responseData);
			} catch(e) {
				console.log(e)
				span.end()
			} finally {
				console.log("span end........... ")
				span.end()
			}
		})

	let currentUser = {}
	let pageProps = {}

	if(appContext.Component.getInitialProps){
        	 // pageProps = await appContext.Component.getInitialProps(appContext.ctx, client, currentUser)
        	 pageProps = await appContext.Component.getInitialProps(appContext.ctx, currentUser)
        }

	return {
        	pageProps,
		currentUser,
		// ...currentUser
        }
}

