<script>
  import { onMount } from 'svelte';
  import { i18n } from './i18n';

  export let key;
  export let count = 1;
  export let args = {};

  let translation = '';
  onMount(async() => {
    translation = await i18n(key, count, args);
  });
</script>

{translation}
