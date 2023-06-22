<script>
   import { onMount } from "svelte";
   import Navbar from "./_components/navbar.svelte";
   import Footer from "./_components/footer.svelte";

   let Todos = [];
   let username = "";
   async function getTodos() {
      const resp = await fetch("/api/todo/list", { method: "GET" });
      Todos = await resp.json();
      console.log(Todos);
   }

   onMount(() => {
      getTodos();
      username = localStorage.getItem("username");
      console.log(username);
   })

   function formatTime(timestamp) {
      const optionsTime = { hour: 'numeric', minute: 'numeric', hour12: true };
      const date = new Date(timestamp);
      return date.toLocaleString('en-US', optionsTime);
   }

   function formatDate(timestamp) {
      const optionsDate = { day: 'numeric', month: 'long', year: 'numeric' };
      const date = new Date(timestamp);
      return date.toLocaleString('en-US', optionsDate);
   }
</script>

<Navbar></Navbar>
<main class="min-h-[80vh] mx-44 mt-20">

   {#if Todos.length > 0}
      <div class="flex flex-col w-9/12 space-y-3 justify-center mx-auto">
         {#each Todos as todo}
            <div class="flex flex-col space-y-2 rounded-xl shadow-md py-3 px-5 border border-zinc-300">
               <div class="flex flex-row justify-between">
                  <h3 class="text-2xl">{todo.text}</h3>
                  {#if todo.username === username}
                     <button>Edit</button>
                  {/if}
               </div>
               <p class="text-xs">Created: {formatDate(todo.created_at)} • {formatTime(todo.created_at)}</p>
            </div>
         {/each}
      </div>
   {:else}
      <p>NO Content</p>
   {/if}

</main>
<Footer></Footer>