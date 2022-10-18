import React, {ReactElement} from 'react';

export default function GET({path}: {path: string}): ReactElement {
	return (
		<div className={'getTag inline-flex flex-row items-center'}>
			<div
				className={'rounded-md px-2 font-bold'}
				style={{backgroundColor: '#0675F9', color: '#FFFFFF'}}>
				{'GET'}
			</div>
			<b className={'pl-2'}>
				{path}
			</b>
		</div>
	);
}
