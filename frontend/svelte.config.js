import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

/** @type {import('@sveltejs/kit').Config} */
const dev = process.env.NODE_ENV === 'development';

const config = {
    preprocess: vitePreprocess(),
    kit: {
        adapter: adapter({
            // Generate a static site that matches Docker's /frontend/dist
            pages: 'dist',
            assets: 'dist',
            fallback: 'index.html'
        }),
        paths: {
            // Serve the built SPA under /app while keeping dev at root
            base: dev ? '' : '/app'
        }
    }
};

export default config;
