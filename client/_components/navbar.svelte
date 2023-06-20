<script>
	import {readable} from "svelte/store";

	const session_id = readable(getCookie("session_id"));

	function getCookie(name) {
		const cookie = document.cookie.match('(^|;)\\s*' + name + '\\s*=\\s*([^;]+)');
		return cookie ? cookie.pop() : '';
	}
</script>

<!-- NavBar -->
<nav class="shadow-md z-50 fixed h-16 w-full top-0 inset-x-0 px-44 bg-zinc-100 text-lg border-b border-zinc-200 flex flex-row justify-between items-center">
   <div class="text-3xl font-bold">
      <span class="bg-gradient-to-br from-sky-500 to-fuchsia-500 text-transparent bg-clip-text">ToDoList</span> - App
   </div>

   <div class="flex flex-row space-x-3">
      <a href="/" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Home</a>
      {#if $session_id}
         <a href="/login" class="py-1.5 px-3 rounded-xl hover:bg-gray-200 font-medium">Login</a>
      {/if}
   </div>
</nav>
<!-- Content Goes Here -->