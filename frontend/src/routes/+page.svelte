<script>
  import Header from "../lib/header.svelte";
  import I18n from "$lib/i18n/I18n.svelte";
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Known anywhere" />
</svelte:head>
<Header />

<!--
<h2>Welcome to SvelteKit</h2>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>
-->

<h1><I18n key="home.welcome"/></h1>
<p><I18n key="home.overview"/></p>

<h2><I18n key="home.howItWorks"/></h2>
<p><I18n key="home.createGroup"/></p>
<p><I18n key="home.logIn"/></p>
<p><I18n key="home.shareGroup"/></p>

<h2><I18n key="home.features"/></h2>
<ul>
  <li><I18n key="home.searchPeople"/></li>
  <li><I18n key="home.searchCoupons"/></li>
  <!--<li><I18n key="home.followGroups"/></li>-->
  <!--<li><I18n key="home.showCoupons"/></li>-->
  <!--<li><I18n key="home.notifications"/></li>-->
  <!--<li><I18n key="home.linkGroups"/></li>-->
</ul>


<!--
<I18n key="test.string"/><br/>
<I18n key="test.withCount" count=1 args={[1]}/><br/>
<I18n key="test.withCount" count=0 args={[0]}/><br/>
-->
