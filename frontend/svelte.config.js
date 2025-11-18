import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const config = {
    preprocess: vitePreprocess(),
    kit: {
        adapter: adapter({
            // Generate a static site that matches Docker's /frontend/dist
            pages: 'dist',
            assets: 'dist',
            fallback: 'index.html'
        })
    }
};

export default config;