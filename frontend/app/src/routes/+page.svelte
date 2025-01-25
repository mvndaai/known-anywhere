<script>
    import { json } from "@sveltejs/kit";
    import { onMount } from "svelte";

    let routes = {};

    onMount(async function() {
        const response = await fetch('http://localhost:8080/api');
        let j = await response.json();
        console.log(j);
        routes = j.data;
    });
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

<h1>Routes</h1>
{#each Object.entries(routes) as [method, rs]}
    <p>{method}</p>
    <ul>
        {#each rs as route}
            <li><a href='{route}'>{route}</a></li>
        {/each}
    </ul>
{/each}
