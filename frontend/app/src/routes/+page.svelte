<script>
    import { onMount } from "svelte";
    import Header from "../lib/header.svelte";

    const backend = 'http://localhost:8080';
    let routes = $state({});

    onMount(async function() {
        const response = await fetch(`${backend}/api`);
        let j = await response.json();
        //console.log(j);
        routes = j.data;
    });
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Known socially" />
</svelte:head>
<Header />


<h2>Routes</h2>
{#each Object.entries(routes) as [method, rs]}
    <p>{method}</p>
    <ul>
        {#each rs as route}
            <li><a href='{route}'>{route}</a></li>
        {/each}
    </ul>
{/each}

<h2>Welcome to SvelteKit</h2>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
