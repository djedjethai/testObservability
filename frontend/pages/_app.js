// import 'bootstrap/dist/css/bootstrap.css'
// import { Buffer } from 'buffer';
// import axios from 'axios';
import Layout from '../components/Layout'
import Header from '../components/header'
// import buildClient from '../services/build-client'
// import { authRoutes } from '../services/config'

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

	// const client = buildClient(appContext.ctx)
 	
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



	// // get the datas of the token and return it within the token
	// // NOTE that normaly should be done at the gateway level
	// // but I thing the gateway can valideta the jwt but not returning the data...
	// try {

	// 	// let { data } = await client.get('/api/v1/auth/jwtgetdata') 
	// 	let { data } = await client.get(authRoutes.jwtgetdata) 
	// 	// say that needed for cookie to be passed in but get block with cors
	// 	// let { data } = await client.get('/api/v1/auth/jwtgetdata', { 
	// 	// 	withCredentials: true, 
	// 	// })

	// 	currentUser = data

	// console.log("at _app seee the datas: ", currentUser)

	// } catch(e){
	// 	if (e.response.status === 400) {
	// 	 	// no jwtoken
	// 	 	console.log("/api/v1/auth/jwtgetdata err === 400: ", e.response.status)
	// 	} else if(e.response.status === 401) {
	// 		console.log("see errrrr from jwtgetdata: ", e.response.status)
	// 		console.log("see errrrr from jwtgetdata: ", e.response.data)
	// 		// invalid/expired jwtoken target endpoint to refresh
	// 		try{
	// 			// let { data } = await client.get('/api/v1/auth/refreshopenid')
	// 			let { data } = await client.get(authRoutes.refreshopenid)
	// 			console.log("seee the data: ", data)
	// 			
	// 			// document.cookie = `jwt_token=${JSON.stringify(data)}`;
	// 			currentUser = data

	// 		} catch(e){
	// 			console.log("e error after refeshingToken: ", e)
	// 			// TODO remove token from cookie, not allow
	// 			// throw e // put that all section into a try then throw
	// 		}
	// 		
	// 	} else if(e.response.status === 403) {
	// 		console.log("Innn 403 no pbbbbbbbbbbb")
	// 		// any other error
	// 		// TODO remove token from cookie, not allow
	// 	} else {
	// 	//
	// 		console.log("In _app in else.....: ", e)
	// 	}
	// }
}

