<script>
  import Input from '../components/Input.svelte'
  import Button from '../components/Button.svelte'
  import { register } from '../lib/api.js'
  import { token } from '../lib/stores.js'

  let { onNavigate = () => {} } = $props()

  let email = $state('')
  let password = $state('')
  let firstName = $state('')
  let lastName = $state('')
  let loading = $state(false)
  let error = $state('')

  async function handleSubmit(e) {
    e.preventDefault()
    error = ''
    loading = true

    try {
      const response = await register(email, password, firstName, lastName)
      token.set(response.token)
    } catch (err) {
      error = err.message || 'Registration failed'
    } finally {
      loading = false
    }
  }
</script>

<div class="auth">
  <div class="auth__card">
    <header class="auth__header">
      <div class="auth__logo">
        <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
          <rect x="3" y="3" width="18" height="18" rx="2"/>
          <path d="M9 12h6M12 9v6"/>
        </svg>
      </div>
      <h1 class="auth__title">Create account</h1>
      <p class="auth__subtitle">Enter your details to get started</p>
    </header>

    <form class="auth__form" onsubmit={handleSubmit}>
      {#if error}
        <div class="auth__error">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <circle cx="12" cy="12" r="10"/>
            <line x1="12" y1="8" x2="12" y2="12"/>
            <line x1="12" y1="16" x2="12.01" y2="16"/>
          </svg>
          <span>{error}</span>
        </div>
      {/if}

      <div class="auth__fields">
        <div class="auth__row">
          <Input
            label="First name"
            type="text"
            bind:value={firstName}
            placeholder="John"
            required
            autocomplete="given-name"
          />

          <Input
            label="Last name"
            type="text"
            bind:value={lastName}
            placeholder="Doe"
            required
            autocomplete="family-name"
          />
        </div>

        <Input
          label="Email"
          type="email"
          bind:value={email}
          placeholder="name@company.com"
          required
          autocomplete="email"
        />

        <Input
          label="Password"
          type="password"
          bind:value={password}
          placeholder="Create a password"
          required
          autocomplete="new-password"
        />
      </div>

      <Button type="submit" {loading}>
        Create account
      </Button>
    </form>

    <footer class="auth__footer">
      <span>Have an account?</span>
      <a href="#/login">Sign in</a>
    </footer>
  </div>
</div>

<style>
  .auth {
    width: 100%;
    max-width: 400px;
  }

  .auth__card {
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    padding: 40px;
  }

  .auth__header {
    text-align: center;
    margin-bottom: 32px;
  }

  .auth__logo {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 48px;
    height: 48px;
    margin-bottom: 20px;
    color: var(--accent);
  }

  .auth__title {
    font-size: 24px;
    font-weight: 500;
    letter-spacing: -0.01em;
    color: var(--text-high);
    margin: 0 0 8px;
  }

  .auth__subtitle {
    font-size: 14px;
    color: var(--text-medium);
    margin: 0;
  }

  .auth__form {
    display: flex;
    flex-direction: column;
    gap: 24px;
  }

  .auth__fields {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .auth__row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 16px;
  }

  .auth__error {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 12px 16px;
    background-color: rgba(244, 67, 54, 0.08);
    border: 1px solid var(--error);
    border-radius: var(--radius);
    color: var(--error);
    font-size: 13px;
  }

  .auth__error svg {
    flex-shrink: 0;
  }

  .auth__footer {
    margin-top: 24px;
    padding-top: 24px;
    border-top: 1px solid var(--border-default);
    text-align: center;
    font-size: 14px;
    color: var(--text-medium);
  }

  .auth__footer a {
    margin-left: 4px;
  }
</style>
