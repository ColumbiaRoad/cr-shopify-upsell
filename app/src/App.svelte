<script>
	import Onboarding from "./components/Onboarding.svelte";
	import Dashboard from "./components/Dashboard.svelte";

	export let subscribed;
</script>

<main>
	{#if subscribed}
		<Dashboard />
	{:else}
		<Onboarding />
	{/if}
</main>


