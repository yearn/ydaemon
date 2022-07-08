/* eslint-disable react/react-in-jsx-scope */
/* eslint-disable @typescript-eslint/explicit-function-return-type */

const GITHUB_URI = 'https://github.com/Majorfi/ydaemon';

export default {
	docsRepositoryBase: GITHUB_URI,
	feedbackLabels: 'feedback',
	feedbackLink: 'Question? Give us feedback →',
	floatTOC: true,
	defaultMenuCollapsed: false,
	footerEditLink: 'Edit this page on GitHub',
	footerText: () => (
		<div className={'text-sm text-current'}>
			{'Yearn Finance '}{new Date().getFullYear()}
		</div>
	),
	GITHUB_URI,
	head: ({meta, title}) => {
		// eslint-disable-next-line react-hooks/rules-of-hooks
		const description = meta.description || 'yDaemon is the next-gen API for Yearn. Based on the one from the [exporter](https://github.com/yearn/yearn-exporter), it brings a lot of new features and benefits without breaking the existing system.';
		const appTitle = title && !title.startsWith('Yearn') ? title + ' – Yearn' : 'yDaemon';
		return (
			<>
				<title>{appTitle}</title>
				<meta httpEquiv={'X-UA-Compatible'} content={'IE=edge'} />
				<meta name={'viewport'} content={'width=device-width, initial-scale=1'} />
				<meta name={'description'} content={description} />
				<meta name={'msapplication-TileColor'} content={'#62688F'} />
				<meta name={'theme-color'} content={'#ffffff'} />

				<link rel={'shortcut icon'} type={'image/x-icon'} href={'/favicons/favicon.ico'} />
				<link rel={'apple-touch-icon'} sizes={'180x180'} href={'/favicons/apple-touch-icon.png'} />
				<link rel={'icon'} type={'image/png'} sizes={'32x32'} href={'/favicons/favicon-32x32.png'} />
				<link rel={'icon'} type={'image/png'} sizes={'16x16'} href={'/favicons/favicon-16x16.png'} />
				<link rel={'icon'} type={'image/png'} sizes={'192x192'} href={'/favicons/android-chrome-192x192.png'} />
				<link rel={'icon'} type={'image/png'} sizes={'512x512'} href={'/favicons/android-chrome-512x512.png'} />

				<meta name={'robots'} content={'index,nofollow'} />
				<meta name={'googlebot'} content={'index,nofollow'} />
				<meta charSet={'utf-8'} />
			</>
		);
	},
	logo: () => {
		// eslint-disable-next-line react-hooks/rules-of-hooks
		return (
			<>
				<div style={{marginRight: 16}}>
					<svg width={'40'} height={'40'} viewBox={'0 0 1024 1024'} fill={'none'} xmlns={'http://www.w3.org/2000/svg'}>
						<circle cx={'512'} cy={'512'} r={'512'} fill={'#6BB941'}/>
						<path d={'M645.095 167.388L512.731 345.535L501.031 330.202L378.173 168.118L295.537 229.447L460.078 449.941V609.105H564.653V449.21L728.463 229.447L645.095 167.388Z'} fill={'white'}/>
						<path d={'M688.242 392.992L623.888 474.764C654.602 503.969 673.616 545.585 673.616 590.852C673.616 679.925 601.218 752.206 512 752.206C422.782 752.206 350.384 679.925 350.384 590.852C350.384 544.855 370.129 503.239 400.844 474.034L337.221 391.532C280.912 439.719 245.81 511.27 245.81 590.852C245.81 737.604 365.01 856.612 512 856.612C658.99 856.612 778.191 737.604 778.191 590.852C778.191 512 743.089 441.179 688.242 392.992Z'} fill={'white'}/>
					</svg>
				</div>
				<span className={'hidden text-xl font-bold md:inline text-neutral-0 font-roboto'}>
					{'yDaemon'}
				</span>
			</>
		);
	},
	nextLinks: true,
	darkMode: true,
	prevLinks: true,
	projectLink: GITHUB_URI,
	search: true,
	titleSuffix: ' – Yearn',
	unstable_flexsearch: true
};
