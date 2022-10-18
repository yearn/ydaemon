import React, {ReactElement} from 'react';
import Document, {DocumentContext, DocumentInitialProps, Head, Html, Main, NextScript} from 'next/document';
import {SkipNavLink} from '@reach/skip-nav';

const modeScript = `
  let darkModeMediaQuery = window.matchMedia('(prefers-color-scheme: dark)')

  updateMode()
  darkModeMediaQuery.addEventListener('change', updateModeWithoutTransitions)
  window.addEventListener('storage', updateModeWithoutTransitions)

  function updateMode() {
    let isSystemDarkMode = darkModeMediaQuery.matches
    let isDarkMode = true

    if (isDarkMode) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }

    if (isDarkMode === isSystemDarkMode) {
      delete window.localStorage.isDarkMode
    }
  }

  function updateModeWithoutTransitions() {
    updateMode()
  }
`;

class MyDocument extends Document {
	static async getInitialProps(ctx: DocumentContext): Promise<DocumentInitialProps> {
		const initialProps = await Document.getInitialProps(ctx);
		return {...initialProps};
	}

	render(): ReactElement {
		return (
			<Html lang={'en'} className={'dark'}>
				<Head>
					<link
						rel={'preconnect'}
						href={'https://brand.yearn.finance'}
						crossOrigin={'true'} />
					<link href={'https://brand.yearn.finance/fonts/fonts.css'} rel={'stylesheet'} />
					<script dangerouslySetInnerHTML={{__html: modeScript}} />
				</Head>
				<body className={'bg-neutral-200 transition-colors duration-150'} data-theme={'dark'}>
					<SkipNavLink />
					<Main />
					<NextScript />
				</body>
			</Html>
		);
	}
}

export default MyDocument;
