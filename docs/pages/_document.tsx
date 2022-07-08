import React, {ReactElement} from 'react';
import Document, {Html, Head, Main, NextScript, DocumentContext} from 'next/document';
import {SkipNavLink} from '@reach/skip-nav';

type TInitialProps = {
    html: string;
    head?: (JSX.Element | null)[] | undefined;
    styles?: React.ReactElement[] | React.ReactFragment | undefined;
}

class MyDocument extends Document {
	static async getInitialProps(ctx: DocumentContext): Promise<TInitialProps> {
		const initialProps = await Document.getInitialProps(ctx);
		return {...initialProps} as any; // eslint-disable-line
	}
	render(): ReactElement {
		return (
			<Html lang={'en'} className={'dark'}>
				<Head>
					<link rel={'preconnect'} href={'https://fonts.googleapis.com'} />
					<link rel={'preconnect'} href={'https://fonts.gstatic.com'} crossOrigin={'true'} />
					<link href={'https://fonts.googleapis.com/css2?family=Roboto+Mono:wght@400;700&family=Roboto:wght@400;700&display=swap'} rel={'stylesheet'} />
				</Head>
				<body className={'bg-neutral-200 transition-colors duration-150'} data-theme={'light'}>
					<SkipNavLink />
					<Main />
					<NextScript />
				</body>
			</Html>
		);
	}
}

export default MyDocument;
