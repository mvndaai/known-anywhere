<script>
  import JwtHelper from "./admin/jwt_helper.svelte";
    const backend = 'http://localhost:8080'
    const ls = (typeof window !== 'undefined') ? window.localStorage : null;
    let loggedIn = $state(false);
    let username = $state('');
    let dialogOpen = $state(false);


    const jwt = ls?.getItem('jwt');
    const claims = jwt ? JSON.parse(atob(jwt.split('.')[1])) : null;
    //console.log(claims);
    //console.log(claims?.exp);

    if (claims?.exp && claims.exp * 1000 > Date.now()) {
        loggedIn = true;
        username = claims.username;
    }

    const logout = () => {
        fetch(`${backend}/api/protected/logout`, {
            method: 'POST',
            headers: {'Authorization': `Bearer ${jwt}`},
        });
        ls.removeItem('jwt');
        loggedIn = false;
    }

    const closeDialog = () => dialogOpen = false;
</script>

<div>
    {#if loggedIn}
        <div>
            <button onclick={logout}>{username}</button>
            <!--<button onclick={logout}>Logout</button>-->
        </div>
    {:else}
        <button onclick={() => {dialogOpen = true}}>Login</button>
        <!--<Button size="small" onClick={onLogin} label="Log in" />-->
        <!--<Button primary size="small" onClick={onCreateAccount} label="Sign up" />-->

        <dialog open={dialogOpen}>
            <button onclick={closeDialog}>X</button>
            <div>
                <input type="text" bind:value={username} placeholder="Username" />
                <button>Log in</button>

                <JwtHelper />
            </div>
        </dialog>
    {/if}
</div>

<style>
    dialog::backdrop {
       background-color: salmon;
    }
</style>