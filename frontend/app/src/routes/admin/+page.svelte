
<script>
    const backend = 'http://localhost:8080'
    const ls = (typeof window !== 'undefined') ? window.localStorage : null;

    // Create Domain
    let domainDisplayName = $state('');
    let domainDescription = $state('');
    let domainNotes = $state('');

    // List domain
    let domains = $state([]);
    let domainListQueryParams = $state('');

    // User
    let userUsername = $state('');
    let userDisplayName = $state('');
    let users = $state([]);
    let userListQueryParams = $state('');



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
            const response = await fetch(`${backend}/api/protected/domain`, {
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

    <span>
        <h3>List</h3>
        <button onclick={async () => {
            const response = await fetch(`${backend}/api/domain${domainListQueryParams}`, {
                headers: {'Content-Type': 'application/json'},
            });
            const j = await response.json();
            console.log(j);
            domains = j.data;

        }}>List</button>
        <input bind:value={domainListQueryParams} type="text" placeholder="Query Params"/>
        <table>
            <thead>
                <tr>
                    <th>Display Name</th>
                    <th>Description</th>
                    <th>Notes</th>
                    <th>ID</th>
                </tr>
            </thead>
            <tbody>
                {#each domains as domain}
                    <tr>
                        <td>{domain.display_name}</td>
                        <td>{domain.description}</td>
                        <td>{domain.notes}</td>
                        <td>{domain.id}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </span>
</div>


<div>
    <h2>User</h2>
    <span>
        <h3>Create</h3>
        <input bind:value={userUsername} type="text" placeholder="Username"/>
        <input bind:value={userDisplayName} type="text" placeholder="Display Name"/>
        <button onclick={async () => {
            const response = await fetch(`${backend}/api/protected/user`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${ls?.getItem('jwt')}`,
                },
                body: JSON.stringify({
                    'username': userUsername,
                    'display_name': userDisplayName,
                }),
            });
            const j = await response.json();
            console.log(j);
        }}>Create</button>
    </span>

    <span>
        <h3>List</h3>
        <button onclick={async () => {
            const response = await fetch(`${backend}/api/user${userListQueryParams}`, {
                headers: {'Content-Type': 'application/json'},
            });
            const j = await response.json();
            console.log(j);
            users = j.data;
        }}>List</button>
        <input bind:value={userListQueryParams} type="text" placeholder="Query Params"/>
        <table>
            <thead>
                <tr>
                    <th>Username</th>
                    <th>Display Name</th>
                    <th>ID</th>
                </tr>
            </thead>
            <tbody>
                {#each users as user}
                    <tr>
                        <td>{user.username}</td>
                        <td>{user.display_name}</td>
                        <td>{user.id}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </span>
</div>