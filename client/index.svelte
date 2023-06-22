<script>
   import { onMount } from "svelte";
   import Navbar from "./_components/navbar.svelte";
   import Footer from "./_components/footer.svelte";

   let Todos = [];
   async function getTodos() {
      const resp = await fetch("/api/todo/list", { method: "GET" });
      Todos = await resp.json();
   }

   onMount(() => {
      getTodos()
   })
</script>

<Navbar></Navbar>
<main class="min-h-[80vh] mx-44 mt-20">

   {#if Todos.length > 0}
      <ul>
         {#each Todos as todo}
            <li>{todo.text}</li>
         {/each}
      </ul>
   {:else}
      <p>NO Content</p>
   {/if}

</main>
<Footer></Footer>