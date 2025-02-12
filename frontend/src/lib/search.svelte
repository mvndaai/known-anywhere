<script>
    let text = $state('');
    let select = $state('username');

    const search = () => {
        if (text === '') {
            return;
        }
        window.location.href = `/search/?text=${text}&type=${select}`;
        //console.log(`Searching for ${select} with text ${text}`);
    };
</script>

<div class="search-bar">
    <select bind:value={select}>
        <option value="username">Username</option>
        <option value="coupon">Coupons</option>
        <option value="name">Name/Alias</option> <!-- TODO is there a better word for a name or alias Moniker-->
        <!--<option>Domain</option>-->
    </select>
    <input type="text" placeholder="Search" bind:value={text}/>
    <button onclick={search} disabled={text === ''}>Search</button>
</div>

<style>
    .search-bar {
        display: flex;
        align-items: center;
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
        padding: 0 1rem;
        background-color: #f0f0f0;
        cursor: pointer;
    }

    .search-bar button:hover {
        background-color: #e0e0e0;
    }
</style>