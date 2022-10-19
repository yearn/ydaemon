import React, {ReactElement, ReactNode} from 'react';

export default function Dropdown({title, children}: {title: string, children: ReactNode}): ReactElement {
	return (
		<details>
			<summary>{title}</summary>
			<div>
				{children}
			</div>
		</details>
	);
}
