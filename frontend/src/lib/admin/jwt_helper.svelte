<script>
    const backend = 'http://localhost:8080';


    let subject = $state(localStorage?.getItem('subject') || '');
    let username = $state(localStorage?.getItem('username') || '');
    let days = $state(localStorage?.getItem('days') || 5);
    const generateJWT = async () => {
        localStorage?.setItem('subject', subject);
        localStorage?.setItem('username', username);
        const response = await fetch(`${backend}/api/test/jwt`, {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body:  JSON.stringify({
                'sub': subject,
                'username': username,
                'exp': Math.floor(new Date().setDate(new Date().getDate() + Number(days)) / 1000),
            }),
        });
        const j = await response.json();
        //console.log(j);
        if (response.status === 200) {
            localStorage?.setItem('jwt', j.data);
        }
    }
</script>

<div class='jwt_helper'>
    <div>
        <span>Create JWT</span>
        <input bind:value={subject} type="text" placeholder="Subject"/>
        <input bind:value={username} type="text" placeholder="Username"/>
        <input bind:value={days} type="number" placeholder="Expiration"/>
        <button onclick={generateJWT}>Generate</button>
    </div>
    <div>
        <span>Test JWT auth</span>
        <button onclick={async () => {
            const jwt = localStorage?.getItem('jwt');
            console.log(jwt);
            const response = await fetch(`${backend}/api/test/auth`, {
                headers: {'Authorization': `Bearer ${jwt}`},
            });
            console.log(response);
            const j = await response.json();
            console.log(j);
        }}>Test Auth</button>
    </div>
</div>
