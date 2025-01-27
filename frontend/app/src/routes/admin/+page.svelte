
<script>
    const backend = 'http://localhost:8080'
    const ls = (typeof window !== 'undefined') ? window.localStorage : null;

    let domainDisplayName = $state('');
    let domainDescription = $state('');
    let domainNotes = $state('');
</script>

<h1>Testing page</h1>

<a href="../">Home</a>


<div>
    <h2>Domain</h2>
    <span>
        <h3>Create</h3>
        <input bind:value={domainDisplayName} type="text" placeholder="Display Name"/>
        <input bind:value={domainDescription} type="text" placeholder="Description"/>
        <input bind:value={domainNotes} type="text" placeholder="Notes"/>
        <button onclick={async () => {
            const response = await fetch(`${backend}/domain`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${ls?.getItem('jwt')}`,
                },
                body: JSON.stringify({
                    'display_name': domainDisplayName,
                    'description': domainDescription,
                    'notes': domainNotes,
                }),
            });
            const j = await response.json();
            console.log(j);
        }}>Create</button>
    </span>

</div>
