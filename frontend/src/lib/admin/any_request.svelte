<script>
    let method = $state('GET');
    let url = $state('');
    let headers = $state('');
    let body = $state('');
    let response = $state('');
    let auth = $state(true);
</script>


<div class='any_request'>
    <div class='any_request__inputs'>
        <select bind:value={method}>
            <option value='GET'>GET</option>
            <option value='POST'>POST</option>
            <option value='PUT'>PUT</option>
            <option value='DELETE'>DELETE</option>
            <option value='OPTIONS'>OPTIONS</option>
        </select>
        <input bind:value={url} type='text' placeholder='URL'>
        <input bind:value={headers} type='text' placeholder='Headers'>
        <input bind:value={body} type='text' placeholder='Body'>
        <label>Auth <input type='checkbox' bind:checked={auth}/></label>
        <button onclick={async () => {
            const requestInfo = {method: method};
            if (auth) {
                const jwt = localStorage?.getItem('jwt');
                requestInfo.headers = {'Authorization': `Bearer ${jwt}`};
            }
            requestInfo
            if (headers && headers.length > 0) {
                requestInfo.headers = { ...requestInfo.headers || {}, ...JSON.parse(headers) };
            }
            if (body && body.length > 0) {
                requestInfo.body = JSON.stringify(body);
            }
            const r = await fetch(url, requestInfo);
            //console.log(r);
            const j = await r.json();
            //console.log(j);
            response = JSON.stringify(j, null, '\t');
        }}>Send</button>
    </div>
    <div class='any_request__response'>
        <pre>{response}</pre>
    </div>

</div>