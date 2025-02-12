<script>
    import Header from "../../lib/header.svelte";
    import { onMount } from 'svelte';

    let type = $state('');
    let text= $state('');
    onMount(() => {
        const params = new URLSearchParams(window.location.search);
        text = params.get('text');
        if (params.has('type')) {
            type = params.get('type');
        }
    });
</script>

<svelte:head>
	<title>Search </title>
	<!--<meta name="description" content="Known socially" />-->
</svelte:head>
<Header />


<h1>Search</h1>
Search for '{type}' with text '{text}'.
