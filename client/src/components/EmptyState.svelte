<script>
  let {
    variant = 'friends'
  } = $props()

  const states = {
    friends: {
      icon: 'M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2 M23 21v-2a4 4 0 0 0-3-3.87 M16 3.13a4 4 0 0 1 0 7.75 M9 7a4 4 0 1 0 0 8 4 4 0 0 0 0-8z',
      title: 'No friends yet',
      description: 'Send a friend request to start connecting'
    },
    incoming: {
      icon: 'M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2 M8.5 7a4 4 0 1 0 0 8 4 4 0 0 0 0-8z M20 8v6 M23 11l-3 3-3-3',
      title: 'No incoming requests',
      description: 'When someone sends you a request, it will appear here'
    },
    outgoing: {
      icon: 'M16 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2 M8.5 7a4 4 0 1 0 0 8 4 4 0 0 0 0-8z M20 8v6 M17 11h6',
      title: 'No outgoing requests',
      description: 'Requests you send will appear here until accepted'
    }
  }

  const state = $derived(states[variant] || states.friends)
</script>

<div class="empty">
  <div class="empty__icon">
    <svg width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
      <path d={state.icon} />
    </svg>
  </div>
  <h3 class="empty__title">{state.title}</h3>
  <p class="empty__description">{state.description}</p>
</div>

<style>
  .empty {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 48px 24px;
    text-align: center;
    animation: empty-fade 400ms ease-out;
  }

  @keyframes empty-fade {
    0% {
      opacity: 0;
      transform: scale(0.95);
    }
    100% {
      opacity: 1;
      transform: scale(1);
    }
  }

  .empty__icon {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 72px;
    height: 72px;
    margin-bottom: 20px;
    background-color: var(--surface-elevated);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    color: var(--text-low);
  }

  .empty__title {
    font-size: 18px;
    font-weight: 500;
    color: var(--text-high);
    margin: 0 0 8px;
  }

  .empty__description {
    font-size: 14px;
    color: var(--text-medium);
    margin: 0;
    max-width: 280px;
  }
</style>
