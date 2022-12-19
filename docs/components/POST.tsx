import React, {ReactElement} from 'react';

export default function POST({path}: {path: string}): ReactElement {
	return (
		<div className={'getTag inline-flex flex-row items-center'}>
			<div
				className={'rounded-md px-2 font-bold'}
				style={{
					backgroundImage: 'linear-gradient(to bottom right, #34A14F, #63C532)',
					color: '#FFFFFF'
				}}>
				{'POST'}
			</div>
			<b className={'pl-2'}>
				{path}
			</b>
		</div>
	);
}
