<script>
    import I18n from "./i18n/I18n.svelte";
    let text = $state('');
    let select = $state('username');

    const search = (event) => {
        event.preventDefault();
        if (text === '') {
            return;
        }
        window.location.href = `/search/?type=${select}&text=${text}`;
    };
</script>

<div class="search-bar">
    <form onsubmit={search}>
        <select bind:value={select}>
            <option value="username"><I18n key="search.select.username"/></option>
            <option value="coupon"><I18n key="search.select.coupon"/></option>
            <option value="name"><I18n key="search.select.name"/></option> <!-- TODO is there a better word for a name or alias Moniker-->
        </select>
        <input type="text" placeholder="Search" bind:value={text}/> <!-- TODO figure out how to i18n a placeholder-->
        <button type="submit" disabled={text === ''}><I18n key="search.search"/></button>
    </form>
</div>

<style>
    .search-bar {
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: #f9f9f9;
        border-radius: 8px;
        box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    }

    .search-bar form {
        display: flex;
        width: 100%;
        max-width: 600px;
    }

    .search-bar select, .search-bar input, .search-bar button {
        height: 2.0rem; /* Ensure all elements have the same height */
        border: 1px solid #ccc;
        padding: 0.5rem;
        margin: 0;
        box-sizing: border-box; /* Ensure padding is included in the height */
    }

    .search-bar select {
        border-radius: 4px 0 0 4px;
    }

    .search-bar input {
        border-radius: 0;
        flex: 1;
    }

    .search-bar button {
        border-radius: 0 4px 4px 0;
        padding: 0 1.5rem;
        background-color: #f0f0f0;
        cursor: pointer;
        border: none;
    }

    .search-bar button:hover {
        background-color: #e0e0e0;
    }

    .search-bar button:disabled {
        background-color: #cccccc;
        cursor: not-allowed;
    }
</style>