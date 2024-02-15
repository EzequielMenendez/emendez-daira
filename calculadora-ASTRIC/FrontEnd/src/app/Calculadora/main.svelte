<script lang="ts">
    import Router from 'svelte-spa-router';
    import routes from '../Calculadora/routes';
    export let prefix: string;
</script>

<main><Router {prefix} {routes} /></main>