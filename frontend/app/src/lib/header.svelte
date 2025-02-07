<script>
</script>

<header>
    <div>
        <h1>Known Socially</h1>
        <br>
        <div>
            <a href="/">Home</a>
            <a href="/admin/request">Requests</a>
            <a href="/admin">Admin</a>
        </div>
    </div>
</header>



<style>
    @import './variables.css';
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
    }
    div {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: row;
    }
    h1 {
        margin: 0;
    }
</style>