import App from './App.svelte';

const app = new App({
	target: document.body,
	props: {
		subscribed: window['subscribed'],
		extensionEnabled: window['extensionEnabled']
	}
});

export default app;