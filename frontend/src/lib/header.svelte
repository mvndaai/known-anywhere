<script>
  import Dialog from "./dialog.svelte";
  import Profile from "./header_profile.svelte";
  import Search from "./search.svelte";
  import { onMount } from 'svelte';
  import Themer from "./themer.svelte";

  let dialogOpen = $state(false);

  const handleKeydown = (event) => {
    if (event.key === '/' && !document.activeElement.matches('input, textarea')) {
        dialogOpen = true;;
    }
  };

  onMount(() => {
    window.addEventListener('keydown', handleKeydown);
    return () => {
      window.removeEventListener('keydown', handleKeydown);
    };
  });
</script>

<header>
    <a class='icon' href="/"><h1>Known Anywhere</h1></a>
    <Search />
    <Profile />
</header>
<Dialog bind:open={dialogOpen}>
    <h1>Theme</h1>
    <Themer />
    <h1>Admin links</h1>
    <a href="/admin/request">Requests</a>
    <a href="/admin">Admin</a>
</Dialog>

<style>
    @import './theme.css';
    header {
        background-color: var(--header-background-color);
        color: var(--header-text-color);
        text-align: center;
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        width: 100%;
        height: var(--header-height);
        box-sizing: border-box;
        z-index: 1000;

        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 0 1rem;

        .icon {
            color: var(--header-text-color);
            text-decoration: none;
        }
    }
</style>