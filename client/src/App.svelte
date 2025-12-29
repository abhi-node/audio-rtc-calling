<script>
  import { onMount } from 'svelte'
  import Login from './routes/Login.svelte'
  import Register from './routes/Register.svelte'
  import Dashboard from './routes/Dashboard.svelte'
  import { token, isAuthenticated } from './lib/stores.js'

  let route = $state(window.location.hash || '#/login')
  let authenticated = $state(false)

  isAuthenticated.subscribe(value => {
    authenticated = value
  })

  function handleHashChange() {
    route = window.location.hash || '#/login'
  }

  function navigate(path) {
    window.location.hash = path
  }

  onMount(() => {
    window.addEventListener('hashchange', handleHashChange)
    return () => window.removeEventListener('hashchange', handleHashChange)
  })
</script>

{#if authenticated}
  <Dashboard onLogout={() => navigate('#/login')} />
{:else if route === '#/register'}
  <Register onNavigate={navigate} />
{:else}
  <Login onNavigate={navigate} />
{/if}
