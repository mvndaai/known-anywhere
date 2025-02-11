<script>
  import Dialog from "./dialog.svelte";
  import JwtHelper from "./admin/jwt_helper.svelte";

  const backend = 'http://localhost:8080';
  let loggedIn = $state(false);
  let username = $state('');
  let dialogOpen = $state(false);
  const handleShowModal = () => dialogOpen = true;
  const jwt = localStorage?.getItem('jwt');
  let claims = null;
  if (jwt) {
    try {
      claims = JSON.parse(atob(jwt.split('.')[1]));
    } catch (e) {
      console.error('Invalid JWT token', e);
    }
  }
  if (claims?.exp && claims.exp * 1000 > Date.now()) {
    loggedIn = true;
    username = claims.username;
  }

  // Function to handle logout
  const logout = () => {
    fetch(`${backend}/api/protected/logout`, {
      method: 'POST',
      headers: { 'Authorization': `Bearer ${jwt}` },
    });
    localStorage.removeItem('jwt');
    loggedIn = false;
  };
</script>

<div>
  {#if loggedIn}
    <div>
      <button onclick={logout}>{username}</button>
      <!--<button onclick={logout}>Logout</button>-->
    </div>
  {:else}
    <button onclick={(handleShowModal)}>Login</button>
    <!--<Button size="small" onClick={onLogin} label="Log in" />-->
    <!--<Button primary size="small" onClick={onCreateAccount} label="Sign up" />-->

    <Dialog bind:open={dialogOpen}>
      <div class='stuff'>
        <button>Create Account</button>
        <!--<input type="text" bind:value={username} placeholder="Username" />-->
        <button>Log in</button>

        <br/><br/><br/>
        <JwtHelper />
      </div>
    </Dialog>

  {/if}
</div>