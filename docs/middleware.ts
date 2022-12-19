// eslint-disable-next-line @next/next/no-server-import-in-page
import {NextRequest, NextResponse} from 'next/server';

export const config = {
	matcher: [
		'/docs',
		'/docs/install/:path*',
		'/docs/hooks/:path*'
	]
};

export async function middleware(request: NextRequest): Promise<NextResponse> {
	if (request.nextUrl.pathname === ('/docs')) {
		const	url = request.nextUrl.clone();
		url.pathname = url.pathname.replace('docs', '');
		return NextResponse.redirect(url);
	}
	if (request.nextUrl.pathname.startsWith('/docs/install')) {
		const	url = request.nextUrl.clone();
		url.pathname = url.pathname.replace('install', '1-install');
		return NextResponse.redirect(url);
	}
	if (request.nextUrl.pathname.startsWith('/docs/hooks')) {
		const	url = request.nextUrl.clone();
		url.pathname = url.pathname.replace('hooks', '3-web3Hooks');
		return NextResponse.redirect(url);
	}
	return NextResponse.next();
}
