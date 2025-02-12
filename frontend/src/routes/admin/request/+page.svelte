<script>
  import { onMount } from "svelte";
  import AnyRequest from "$lib/admin/any_request.svelte";
  import JwtHelper from "$lib/admin/jwt_helper.svelte";
  import Header from "$lib/header.svelte";

  const backend = 'http://localhost:8080';
  let routes = $state({});

  onMount(async function() {
    const response = await fetch(`${backend}/api/test/list`);
    let j = await response.json();
    //console.log(j);
    routes = j.data;
});
</script>

<Header />
<AnyRequest />
<JwtHelper />


<h2>Routes</h2>
{#each Object.entries(routes) as [route, methods]}
    <ul>
        {#if methods.includes('GET') }
            <a href={route}>{route}</a> - {methods}
        {:else}
            {route} - {methods}
        {/if}
    </ul>
{/each}