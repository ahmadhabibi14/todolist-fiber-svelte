<script>
	import { onMount } from "svelte";
   import { writable } from "svelte/store";

   let isLoggedIn = writable(false);
   onMount(() => {
      getCookie("session_id");
   });
   function getCookie(cookie_name) {
      const cookies = document.cookie.split(";");
      for (let i = 0; i < cookies.length; i++) {
         const cookie = cookies[i].trim();
         if (cookie.startsWith(cookie_name + "=")) {
            const cookieValue = cookie.substring(cookie_name + 1);
            const decodedCookieValue = decodeURIComponent(cookieValue);

            const cookieParts = decodedCookieValue.split(";");
            // const value = cookieParts[0].trim();
            const expiration = cookieParts.find((part) => {
               part.trim().startsWith("expires=");
            });

            if (isCookieNotExpired(expiration)) {
               isLoggedIn.set(true)
            } else {
               isLoggedIn.set(false)
            }
         }
      }
   }

   function isCookieNotExpired(expiration) {
      const expirationDate = new Date(expiration.substring(8).trim());
      return expirationDate > new Date();
   }

</script>

<!-- NavBar -->
<nav class="shadow-md z-50 fixed h-16 w-full top-0 inset-x-0 px-44 bg-zinc-100 text-lg border-b border-zinc-200 flex flex-row justify-between items-center">
   <div class="text-3xl font-bold">
      <span class="bg-gradient-to-br from-sky-500 to-fuchsia-500 text-transparent bg-clip-text">ToDoList</span> - App
   </div>

   <div class="flex flex-row space-x-3">
      <a href="/" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Home</a>
      {#if !isLoggedIn}
         <a href="/login" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Login</a>
      {/if}
   </div>
</nav>
<!-- Content Goes Here -->