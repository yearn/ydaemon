import React from 'react';
import {useTheme} from 'next-themes';
import {useClientEffect} from '@yearn-finance/web-lib';
import type {AppProps} from 'next/app';
import	'../style.css';

function	WrappedApp({Component, pageProps}: AppProps): React.ReactElement {
	const {theme} = useTheme();

	useClientEffect((): void => {
		document.body.dataset.theme = theme;
	}, [theme]);

	return (
		<Component {...pageProps} />
	);
}

function	App(props: AppProps): React.ReactElement {
	const getLayout = (props.Component as any).getLayout || ((page: React.ReactElement): React.ReactElement => page); // eslint-disable-line

	return (getLayout(<WrappedApp {...props} />));
}

export default App;
