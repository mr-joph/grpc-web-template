import { defineConfig } from 'vite';
import { svelte } from '@sveltejs/vite-plugin-svelte';

export default defineConfig({
	plugins: [svelte()],

	server: {
		port: 1234,
		open: true,
	},

	base: "projects/core/website/dist/"
});
