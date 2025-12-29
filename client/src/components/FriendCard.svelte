<script>
  let {
    user,
    variant = 'friend',
    onaction = () => {},
    loading = false,
    animationDelay = 0
  } = $props()

  function getInitials(firstName, lastName) {
    return `${firstName?.[0] || ''}${lastName?.[0] || ''}`.toUpperCase()
  }

  function getGradient(name) {
    const gradients = [
      'linear-gradient(135deg, #667eea 0%, #764ba2 100%)',
      'linear-gradient(135deg, #f093fb 0%, #f5576c 100%)',
      'linear-gradient(135deg, #4facfe 0%, #00f2fe 100%)',
      'linear-gradient(135deg, #43e97b 0%, #38f9d7 100%)',
      'linear-gradient(135deg, #fa709a 0%, #fee140 100%)',
      'linear-gradient(135deg, #a8edea 0%, #fed6e3 100%)',
      'linear-gradient(135deg, #ff9a9e 0%, #fecfef 100%)',
      'linear-gradient(135deg, #00c6fb 0%, #005bea 100%)',
    ]
    const hash = (name || '').split('').reduce((acc, c) => acc + c.charCodeAt(0), 0)
    return gradients[hash % gradients.length]
  }

  function getActionText() {
    if (variant === 'incoming') return 'Accept'
    if (variant === 'outgoing') return 'Pending'
    return ''
  }

  function getActionVariant() {
    if (variant === 'incoming') return 'accept'
    if (variant === 'outgoing') return 'pending'
    return ''
  }
</script>

<div class="card" style="animation-delay: {animationDelay}ms">
  <div class="card__avatar" style="background: {getGradient(user.first_name + user.last_name)}">
    <span class="card__initials">{getInitials(user.first_name, user.last_name)}</span>
  </div>

  <div class="card__info">
    <span class="card__name">{user.first_name} {user.last_name}</span>
    <span class="card__email">{user.email}</span>
  </div>

  {#if variant !== 'friend'}
    <button
      class="card__action card__action--{getActionVariant()}"
      disabled={variant === 'outgoing' || loading}
      onclick={() => onaction(user)}
    >
      {#if loading}
        <span class="card__loader"></span>
      {:else}
        {getActionText()}
      {/if}
    </button>
  {/if}
</div>

<style>
  .card {
    display: flex;
    align-items: center;
    gap: 16px;
    padding: 16px;
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    animation: card-enter 300ms cubic-bezier(0.34, 1.56, 0.64, 1) backwards;
    transition:
      background-color var(--transition-fast),
      border-color var(--transition-fast),
      transform var(--transition-fast),
      box-shadow var(--transition-fast);
  }

  .card:hover {
    background-color: var(--surface-elevated);
    border-color: var(--border-hover);
    transform: translateY(-2px);
    box-shadow: var(--shadow-2);
  }

  @keyframes card-enter {
    0% {
      opacity: 0;
      transform: translateY(16px);
    }
    100% {
      opacity: 1;
      transform: translateY(0);
    }
  }

  .card__avatar {
    flex-shrink: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 48px;
    height: 48px;
    border-radius: var(--radius);
    box-shadow: var(--shadow-1);
  }

  .card__initials {
    font-size: 16px;
    font-weight: 600;
    color: white;
    text-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
  }

  .card__info {
    flex: 1;
    min-width: 0;
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .card__name {
    font-size: 15px;
    font-weight: 500;
    color: var(--text-high);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .card__email {
    font-size: 13px;
    color: var(--text-medium);
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
  }

  .card__action {
    flex-shrink: 0;
    display: inline-flex;
    align-items: center;
    justify-content: center;
    min-width: 80px;
    height: 36px;
    padding: 0 16px;
    font-family: inherit;
    font-size: 12px;
    font-weight: 600;
    letter-spacing: 0.03em;
    text-transform: uppercase;
    border: none;
    border-radius: var(--radius);
    cursor: pointer;
    transition:
      background-color var(--transition-fast),
      transform var(--transition-fast),
      box-shadow var(--transition-fast);
  }

  .card__action--accept {
    color: var(--surface-ground);
    background-color: var(--success);
    box-shadow: var(--shadow-1);
  }

  .card__action--accept:hover:not(:disabled) {
    background-color: #5cb860;
    transform: scale(1.02);
    box-shadow: var(--shadow-2);
  }

  .card__action--accept:active:not(:disabled) {
    transform: scale(0.98);
  }

  .card__action--pending {
    color: var(--text-medium);
    background-color: var(--surface-elevated);
    cursor: not-allowed;
  }

  .card__action:disabled {
    opacity: 0.7;
  }

  .card__loader {
    width: 16px;
    height: 16px;
    border: 2px solid transparent;
    border-top-color: currentColor;
    border-radius: 50%;
    animation: spin 600ms linear infinite;
  }

  @keyframes spin {
    to {
      transform: rotate(360deg);
    }
  }
</style>
