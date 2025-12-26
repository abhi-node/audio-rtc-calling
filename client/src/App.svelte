<script>
  import { onMount } from 'svelte'
  import Login from './routes/Login.svelte'
  import Register from './routes/Register.svelte'
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

  function logout() {
    token.clear()
  }

  onMount(() => {
    window.addEventListener('hashchange', handleHashChange)
    return () => window.removeEventListener('hashchange', handleHashChange)
  })
</script>

{#if authenticated}
  <div class="dashboard">
    <div class="dashboard__card">
      <div class="dashboard__icon">
        <svg width="28" height="28" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <polyline points="20 6 9 17 4 12"/>
        </svg>
      </div>
      <h1 class="dashboard__title">Authenticated</h1>
      <p class="dashboard__text">You have successfully signed in</p>
      <button class="dashboard__button" onclick={logout}>
        Sign out
      </button>
    </div>
  </div>
{:else if route === '#/register'}
  <Register onNavigate={navigate} />
{:else}
  <Login onNavigate={navigate} />
{/if}

<style>
  .dashboard {
    width: 100%;
    max-width: 400px;
  }

  .dashboard__card {
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    padding: 48px 40px;
    text-align: center;
  }

  .dashboard__icon {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 56px;
    height: 56px;
    margin-bottom: 24px;
    background-color: rgba(76, 175, 80, 0.1);
    border: 1px solid var(--success);
    border-radius: var(--radius);
    color: var(--success);
  }

  .dashboard__title {
    font-size: 24px;
    font-weight: 500;
    letter-spacing: -0.01em;
    color: var(--text-high);
    margin: 0 0 8px;
  }

  .dashboard__text {
    font-size: 14px;
    color: var(--text-medium);
    margin: 0 0 32px;
  }

  .dashboard__button {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    height: 40px;
    padding: 0 20px;
    font-family: inherit;
    font-size: 13px;
    font-weight: 500;
    letter-spacing: 0.02em;
    text-transform: uppercase;
    color: var(--text-high);
    background-color: transparent;
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    cursor: pointer;
    transition:
      background-color var(--transition-fast),
      border-color var(--transition-fast);
  }

  .dashboard__button:hover {
    background-color: var(--surface-elevated);
    border-color: var(--border-hover);
  }

  .dashboard__button:active {
    background-color: var(--surface-overlay);
  }
</style>
