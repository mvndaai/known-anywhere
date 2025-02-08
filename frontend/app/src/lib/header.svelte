<script>
  import Profile from "./profile.svelte";
  import Search from "./search.svelte";
</script>

<header>
    <a href="/"><h1>Known Socially</h1></a>
    <Search />
    <Profile />
</header>

<style>
    @import './theme.css';
    header {
        background-color: var(--color-dark);
        color: var(--color-white);
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

        a {
            color: var(--color-white);
            text-decoration: none;
        }
    }
</style>