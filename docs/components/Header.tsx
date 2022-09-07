import React, {ReactElement} from 'react';

export function Header(): ReactElement {
	return (
		<header className={'flex flex-col items-center mb-10'}>
			<div className={'mt-8 w-auto h-12 md:h-14'}>
				<p className={'sr-only'}>{'yDaemon'}</p>
				<svg width={'64'} height={'64'} viewBox={'0 0 1024 1024'} fill={'none'} xmlns={'http://www.w3.org/2000/svg'}>
					<circle cx={'512'} cy={'512'} r={'512'} fill={'#6BB941'}/>
					<path d={'M645.095 167.388L512.731 345.535L501.031 330.202L378.173 168.118L295.537 229.447L460.078 449.941V609.105H564.653V449.21L728.463 229.447L645.095 167.388Z'} fill={'white'}/>
					<path d={'M688.242 392.992L623.888 474.764C654.602 503.969 673.616 545.585 673.616 590.852C673.616 679.925 601.218 752.206 512 752.206C422.782 752.206 350.384 679.925 350.384 590.852C350.384 544.855 370.129 503.239 400.844 474.034L337.221 391.532C280.912 439.719 245.81 511.27 245.81 590.852C245.81 737.604 365.01 856.612 512 856.612C658.99 856.612 778.191 737.604 778.191 590.852C778.191 512 743.089 441.179 688.242 392.992Z'} fill={'white'}/>
				</svg>
			</div>

			<h1 className={'mb-6 text-lg font-bold text-center text-neutral-700 md:text-4xl'}>
				{'yDaemon'}
			</h1>

			<div className={'flex flex-wrap gap-2 justify-center max-w-[28rem] min-h-[3rem]'}>
				<a aria-label={'Version'} href={'https://www.npmjs.com/package/yearn/ydaemon'} className={'h-5'}>
					<img
						alt={''}
						src={'https://img.shields.io/github/workflow/status/yearn/ydaemon/Go?colorA=2B323B&colorB=1e2329&style=flat&label=Build'}
					/>
				</a>

				<a aria-label={'Downloads'} href={'https://www.npmjs.com/package/yearn/ydaemon'} className={'h-5'}>
					<img
						alt={''}
						src={'https://img.shields.io/github/go-mod/go-version/yearn/ydaemon?colorA=2B323B&colorB=1e2329&style=flat&label=Version'}
					/>
				</a>

				<a aria-label={'Stars'} href={'https://github.com/yearn/ydaemon'} className={'h-5'}>
					<img
						alt={''}
						src={'https://img.shields.io/github/stars/yearn/ydaemon?colorA=2B323B&colorB=1e2329&style=flat&label=Stars'}
					/>
				</a>
			</div>
		</header>
	);
}
