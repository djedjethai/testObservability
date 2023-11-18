import Link from 'next/link'

export default ({ currentUser }) => {

	const isCurrentUser = Object.keys(currentUser).length > 0

	const links = [
		// tricks if the first condition is true the second will be render
		// if first condition is false return false
		!isCurrentUser && {labels:'Sign Up', href:'/auth/signup'},
		!isCurrentUser && {labels:'Sign In', href:'/auth/signin'},
		isCurrentUser && {labels:'Create order', href:'/tickets/new'},
		isCurrentUser && {labels:'Get Order', href:'/orders'},
		isCurrentUser && {labels:'Sign Out', href:'/auth/signout'}
	]
		.filter(linkConfig => linkConfig) // to return only true elemt
		.map(({ labels, href }) => {	
			return <li key={href} className="nav-item">
				<Link href={href} legacyBehavior>
					<a className="nav-link">{labels}</a>
				</Link>
			</li>
		})

	return <nav className="navbar navbar-light bg-light">
		<Link href="/" legacyBehavior>
			<a className="navbar-brand">Apprendre ASR</a>
		</Link>

		<div className="d-flex justify-content-end">
			<ul className="nav d-flex align-items-center">
				{links}
			</ul>
		</div>
	</nav>
}

