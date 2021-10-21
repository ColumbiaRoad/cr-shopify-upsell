<script>
	import Onboarding from "./components/Onboarding.svelte";
	import Dashboard from "./components/Dashboard.svelte";

	export let subscribed;
	export let extensionEnabled;
</script>

{#if subscribed}
	<Dashboard extensionEnabled={extensionEnabled} />
{:else}
	<Onboarding />
{/if}
