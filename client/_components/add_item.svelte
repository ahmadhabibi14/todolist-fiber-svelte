<script>
   import { createEventDispatcher } from "svelte";

   export let isOpen = false;
   export let session_id = "";
   const dispatch = createEventDispatcher();

   let data = {
      text: "",
   }

   const submitText = async (e) => {
      const reqBody = JSON.stringify({
         text: data.text
      })
      const actionURL = e.target.action;
      try {
         const resp = await fetch(actionURL, {
            method: "POST",
            headers: {
               "X-Session-ID": session_id,
               "Content-Type": "application/json"
            },
            body: reqBody
         });
         if (resp.ok) {
            const msg = await resp.json();
            console.error(msg.message);
            alert(msg.message);
            dispatch("close");
            location.reload();
         }
      } catch (error) {
         console.error(error);
         alert(msg.message);
      }
      data.text = "";
   }
</script>

{#if isOpen}
<div class="w-full h-[96%] overflow-y-hidden top-0 left-0 absolute bg-zinc-800 bg-opacity-70 backdrop-blur-md flex justify-center">
   <form action="/api/todo/add" on:submit|preventDefault={submitText}
      class="mt-24 w-6/12 h-fit p-4 bg-white rounded-xl border-zinc-300 shadow-lg flex flex-col space-y-3"
   >
      <textarea
         rows="6"
         class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 p-4 outline-0 w-full"
         type="text"
         id="text"
         name="text"
         placeholder="Text"
         bind:value={data.text}
         required
      ></textarea>
      <button type="submit" class="py-1.5 px-3 bg-sky-600 rounded-lg font-medium text-zinc-50 hover:bg-sky-500">
         Submit
      </button>
   </form>
</div>
{/if}