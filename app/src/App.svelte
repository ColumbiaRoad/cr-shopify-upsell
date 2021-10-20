<script>
	import Onboarding from "./components/Onboarding.svelte";
	import Dashboard from "./components/Dashboard.svelte";

	export let subscribed;
	export let extentionEnabled;
</script>

{#if subscribed}
	<Dashboard extentionEnabled={extentionEnabled} />
{:else}
	<Onboarding />
{/if}
