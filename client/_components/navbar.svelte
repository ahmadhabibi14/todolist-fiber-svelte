<script>
	import { onMount } from "svelte";

   let user = { loggedIn: false }
   onMount(() => {
      const cookies = document.cookie.split(";");
      cookies.forEach(cookie => {
         const [name, value] = cookie.trim().split("=");
         if (name === "session_id") {
            console.log("session_id:", value)
            user.loggedIn = true;
         }
      })
   });
</script>

<!-- NavBar -->
<nav class="shadow-md z-50 fixed h-16 w-full top-0 inset-x-0 px-44 bg-zinc-100 text-lg border-b border-zinc-200 flex flex-row justify-between items-center">
   <div class="text-3xl font-bold">
      <span class="bg-gradient-to-br from-sky-500 to-fuchsia-500 text-transparent bg-clip-text">ToDoList</span> - App
   </div>

   <div class="flex flex-row space-x-3">
      <a href="/" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Home</a>
      {#if !user.loggedIn}
         <a href="/login" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Login</a>
      {/if}
   </div>
</nav>
<!-- Content Goes Here -->