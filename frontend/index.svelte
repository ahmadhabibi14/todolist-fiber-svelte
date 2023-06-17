<script>
	let promise = fetch(
      'http://localhost:3000/api/data',
      {
         method: "GET"
      }
   ).then((x) => x.json());
</script>

{#await promise}
	<!-- optionally show something while promise is pending -->
   <div>Loading....</div>
{:then data}
	<!-- promise was fulfilled -->
	<div>
      <h1>Welcome {data.username}</h1>
   </div>
{:catch error}
	<!-- optionally show something while promise was rejected -->
   <div>Error</div>
{/await}