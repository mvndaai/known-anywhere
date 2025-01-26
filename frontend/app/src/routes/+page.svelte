<script>
    // import { json } from "@sveltejs/kit";
    import { onMount } from "svelte";

    const backend = 'http://localhost:8080'

    let routes = $state({});

    onMount(async function() {
        const response = await fetch(`${backend}/api`);
        let j = await response.json();
        //console.log(j);
        routes = j.data;
    });

    const ls = (typeof window !== 'undefined') ? window.localStorage : null;


    let subject = $state(ls?.getItem('subject') || '');
    let username = $state(ls?.getItem('username') || '');
    let generateJWT = async () => {
        const response = await fetch(`${backend}/test/jwt`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body:  JSON.stringify({
                'sub': subject,
                'username': username,
                'exp': Math.floor(new Date().setDate(new Date().getDate() + 5) / 1000),
            }),
        });
        let j = await response.json();
        ls?.setItem('subject', subject);
        ls?.setItem('username', username);
        ls?.setItem('jwt', j.data);
        //console.log(j);
    }
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


<div>
    <span>Create JWT</span>
    <input bind:value={subject} type="text" placeholder="Subject"/>
    <input bind:value={username} type="text" placeholder="Username"/>
    <button onclick={generateJWT}>Generate</button>
</div>