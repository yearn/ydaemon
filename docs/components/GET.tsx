import React, {ReactElement} from 'react';

export default function GET({path}: {path: string}): ReactElement {
	return (
		<div className={'getTag inline-flex flex-row items-center'}>
			<div
				className={'rounded-md px-2 font-bold'}
				style={{
					backgroundImage: 'linear-gradient(to bottom right, #0675F9, #00B4FF)',
					color: '#FFFFFF'
				}}>
				{'GET'}
			</div>
			<b className={'pl-2'}>
				{path}
			</b>
		</div>
	);
}
