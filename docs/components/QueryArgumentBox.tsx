import React, {ReactElement} from 'react';

type TQueryArgumentBoxProps = {
	title: string;
	description: string;
	children: ReactElement;
}
function QueryArgumentBox({title, description, children}: TQueryArgumentBoxProps): ReactElement {
	return (
		<ul>
			<li>
				<p>
					<b>{title}</b>{': '}{description}
				</p>
				<div style={{marginTop: -16}}>
					{children}
				</div>
			</li>
		</ul>
	);
}

export default QueryArgumentBox;