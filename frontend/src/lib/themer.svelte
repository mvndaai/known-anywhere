<script>
  import { onMount } from "svelte";
  const localStorageKey = "theme";

  const getPersistedThemeObject = () => {
    const v = localStorage.getItem(localStorageKey) || "{}";
    return JSON.parse(v);
  };

  const getPersistedThemeItem = (key) => {
    return getPersistedThemeObject()[key];
  };
  const setPersistedThemeItem = (key, value) => {
    const obj = getPersistedThemeObject();
    obj[key] = value;
    localStorage.setItem(localStorageKey, JSON.stringify(obj));
  };

  const applyThemeItem = (key, value) => {
    if (persistChanges) {
      setPersistedThemeItem(key, value);
    }
    document.documentElement.style.setProperty(key, value);
  };
  const applyFullTheme = () => {
    Object.entries(getPersistedThemeObject()).forEach(([key, value]) => {
      applyThemeItem(key, value);
    });
  };

  const getCSSVariables = () => {
    return [].slice
      .call(document.styleSheets)
      .map((styleSheet) => [].slice.call(styleSheet.cssRules))
      .flat()
      .filter((cssRule) => cssRule.selectorText === ":root")
      .map((cssRule) =>
        cssRule.cssText.split("{")[1].split("}")[0].trim().split(";")
      )
      .flat()
      .filter((text) => text !== "")
      .map((text) => text.split(":"))
      .map((parts) => ({
        key: parts[0].trim(),
        value: getPersistedThemeItem(parts[0].trim()) || parts[1].trim(),
      }));
  };

  let cssVariables = $state();
  let persistChanges = $state(getPersistedThemeItem("persistTheme") || false);

  onMount(async function () {
    applyFullTheme();
    cssVariables = getCSSVariables();
    //console.log("cssVariables", cssVariables);
    //console.log("persistedChanges", getPersistedThemeObject());
    //console.log("persistChanges", persistChanges);
  });

  $effect(() => {
    setPersistedThemeItem("persistTheme", persistChanges);
  });
</script>

<div>
    <div>
        <label>
            Persist Changes
            <input type="checkbox" bind:checked={persistChanges} />
        </label>
        <button onclick={() => {
            delete localStorage[localStorageKey];
        }}>Clear storage</button>

    </div>


  {#each cssVariables as { key, value }, i}
    <div>
      <span>{key}</span>
      <input
        type="text"
        bind:value={cssVariables[i].value}
        onchange={() => {
          applyThemeItem(key, value);
        }}
      />
    </div>
  {/each}
  <!--TODO add an upload button to upload a theme file-->
</div>
