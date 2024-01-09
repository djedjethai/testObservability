// import Head from 'next/head'
// import withAuth from '../hoc/withAuth'
import OpenTelemetryRegistration from './OpenTelemetryRegistration'


// export default function Layout({ children }) {
const Layout = ({ children, currentUser }) => {
 	return (
		<>
      			<main>
				<OpenTelemetryRegistration />
				{children}
			</main>
		</>
  	)
}

// export default withAuth(Layout)
export default Layout

