<script>
  import { onMount } from 'svelte';
  import Navbar from '../_components/navbar.svelte';
  import Footer from '../_components/footer.svelte';

  let user = { loggedIn: false };

  onMount(() => {
    const cookies = document.cookie.split(';');
    cookies.forEach(cookie => {
      const [name, value] = cookie.trim().split('=');
      if (name === 'session_id') {
        console.log('session_id:', value);
        user.loggedIn = true;
      }
    });
  });

  let data = {
    username: '',
    password: '',
  };

  const handleSubmit = async e => {
    const reqBody = JSON.stringify({
      username: data.username,
      password: data.password,
    });
    const actionURL = e.target.action;
    try {
      const resp = await fetch(actionURL, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: reqBody,
      });

      if (resp.ok) {
        const creds = await resp.json();
        const cookieHeader = resp.headers.get('Set-Cookie');
        console.log(creds);
        localStorage.setItem('user_id', creds.user_id);
        document.cookie = cookieHeader;
        window.location.href = '/';
      } else {
        const errorData = await resp.json();
        alert('Login failed');
        location.reload();
      }
    } catch (error) {
      // Handle any network or server errors
      console.error(error);
      alert('Unexpected error occurred during login');
      location.reload();
    }
    data.username = '';
    data.password = '';
  };
</script>

<Navbar />
<main class="min-h-[78vh] mx-44 mt-20">
  <div class="w-full h-fit flex flex-col items-center justify-center pt-16">
    {#if !user.loggedIn}
      <form
        action="/api/login"
        on:submit|preventDefault={handleSubmit}
        class="flex flex-col space-y-4 justify-start w-5/12 p-5 border border-zinc-200 rounded-2xl shadow-lg"
      >
        <h2 class="text-3xl font-bold text-sky-600 text-center">Login</h2>
        <input
          class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 py-2 px-4 outline-0"
          type="text"
          id="username"
          label="username"
          placeholder="Username"
          bind:value={data.username}
          required
        />
        <input
          class="border-zinc-400 border rounded-xl bg-transparent focus:border-blue-500 caret-blue-500 py-2 px-4 outline-0"
          type="password"
          id="password"
          label="password"
          placeholder="Password"
          bind:value={data.password}
          required
        />
        <button
          class="shadow-md hover:shadow-xl rounded-xl py-2 px-6 text-zinc-50 font-medium bg-gradient-to-br from-sky-500 hover:from-sky-600 to-fuchsia-500 hover:to-fuchsia-600"
          type="submit"
        >
          Login
        </button>
      </form>
    {:else}
      <div>Already Logged In</div>
    {/if}
  </div>
</main>
<Footer />
