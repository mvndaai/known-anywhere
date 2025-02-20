//import adapter from '@sveltejs/adapter-auto';
import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter-auto only supports some environments, see https://svelte.dev/docs/kit/adapter-auto for a list.
		// If your environment is not supported, or you settled on a specific environment, switch out the adapter.
		// See https://svelte.dev/docs/kit/adapters for more information about adapters.
		adapter: adapter({
			pages: '../bin/frontend',
			assets: '../bin/frontend',
			//fallback: '/api' // may differ from host to host
		}),
		appDir: 'app', // Change the directory for build files. Github didn't like /_app/
		//prerender: {
		//	//crawl: true,
		//	entries: ['*']
		//},
	}
};

export default config;
