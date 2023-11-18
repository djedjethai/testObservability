// import Head from 'next/head'
// import withAuth from '../hoc/withAuth'


// export default function Layout({ children }) {
const Layout = ({ children, currentUser }) => {
 	return (
		<>
      		<main>{children}</main>
		</>
  	)
}

// export default withAuth(Layout)
export default Layout

