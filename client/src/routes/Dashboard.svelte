<script>
  import { onMount, onDestroy } from 'svelte'
  import TabBar from '../components/TabBar.svelte'
  import FriendCard from '../components/FriendCard.svelte'
  import EmptyState from '../components/EmptyState.svelte'
  import Modal from '../components/Modal.svelte'
  import Button from '../components/Button.svelte'
  import Input from '../components/Input.svelte'
  import { friendsStore } from '../lib/friendsStore.js'
  import { sendFriendRequest, acceptFriendRequest } from '../lib/api.js'
  import { token } from '../lib/stores.js'

  let { onLogout = () => {} } = $props()

  const { friends, incoming, outgoing, loading, error, incomingCount, outgoingCount, startPolling, stopPolling, refetch } = friendsStore

  let activeTab = $state('friends')
  let modalOpen = $state(false)
  let friendEmail = $state('')
  let sendingRequest = $state(false)
  let sendError = $state('')
  let acceptingId = $state(null)

  let friendsList = $state([])
  let incomingList = $state([])
  let outgoingList = $state([])
  let incomingBadge = $state(0)
  let outgoingBadge = $state(0)
  let isLoading = $state(true)
  let errorMessage = $state(null)

  friends.subscribe(v => friendsList = v)
  incoming.subscribe(v => incomingList = v)
  outgoing.subscribe(v => outgoingList = v)
  incomingCount.subscribe(v => incomingBadge = v)
  outgoingCount.subscribe(v => outgoingBadge = v)
  loading.subscribe(v => isLoading = v)
  error.subscribe(v => errorMessage = v)

  const tabs = $derived([
    { id: 'friends', label: 'Friends', count: 0 },
    { id: 'incoming', label: 'Incoming', count: incomingBadge },
    { id: 'outgoing', label: 'Outgoing', count: outgoingBadge }
  ])

  let cleanup

  onMount(() => {
    cleanup = startPolling(10000)
  })

  onDestroy(() => {
    if (cleanup) cleanup()
    stopPolling()
  })

  async function handleSendRequest() {
    if (!friendEmail.trim()) return

    sendingRequest = true
    sendError = ''

    try {
      await sendFriendRequest(friendEmail.trim())
      friendEmail = ''
      modalOpen = false
      await refetch()
    } catch (e) {
      sendError = e.message
    } finally {
      sendingRequest = false
    }
  }

  async function handleAccept(user) {
    acceptingId = user.user_id

    try {
      await acceptFriendRequest(user.user_id)
      await refetch()
    } catch (e) {
      console.error('Failed to accept:', e)
    } finally {
      acceptingId = null
    }
  }

  function handleLogout() {
    token.clear()
    onLogout()
  }

  function getCurrentList() {
    if (activeTab === 'friends') return friendsList
    if (activeTab === 'incoming') return incomingList
    if (activeTab === 'outgoing') return outgoingList
    return []
  }
</script>

<div class="dashboard">
  <header class="dashboard__header">
    <h1 class="dashboard__title">Friends</h1>
    <div class="dashboard__actions">
      <button class="dashboard__add" onclick={() => modalOpen = true}>
        <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <line x1="12" y1="5" x2="12" y2="19"/>
          <line x1="5" y1="12" x2="19" y2="12"/>
        </svg>
        <span>Add Friend</span>
      </button>
      <button class="dashboard__logout" onclick={handleLogout} aria-label="Sign out">
        <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"/>
          <polyline points="16 17 21 12 16 7"/>
          <line x1="21" y1="12" x2="9" y2="12"/>
        </svg>
      </button>
    </div>
  </header>

  <TabBar {tabs} bind:activeTab />

  <main class="dashboard__content">
    {#if isLoading}
      <div class="dashboard__loading">
        <div class="dashboard__skeleton"></div>
        <div class="dashboard__skeleton"></div>
        <div class="dashboard__skeleton"></div>
      </div>
    {:else if errorMessage}
      <div class="dashboard__error">
        <p>{errorMessage}</p>
        <button onclick={refetch}>Try again</button>
      </div>
    {:else if getCurrentList().length === 0}
      <EmptyState variant={activeTab} />
    {:else}
      <div class="dashboard__list">
        {#each getCurrentList() as user, i (user.user_id)}
          <FriendCard
            {user}
            variant={activeTab === 'friends' ? 'friend' : activeTab}
            animationDelay={i * 50}
            loading={acceptingId === user.user_id}
            onaction={handleAccept}
          />
        {/each}
      </div>
    {/if}
  </main>
</div>

<Modal bind:open={modalOpen} title="Add Friend">
  <form class="add-form" onsubmit={(e) => { e.preventDefault(); handleSendRequest() }}>
    <Input
      label="Email Address"
      type="email"
      bind:value={friendEmail}
      placeholder="friend@example.com"
      required
    />
    {#if sendError}
      <div class="add-form__error">
        <svg width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/>
          <line x1="12" y1="8" x2="12" y2="12"/>
          <line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        <span>{sendError}</span>
      </div>
    {/if}
    <div class="add-form__actions">
      <Button type="submit" loading={sendingRequest}>
        Send Request
      </Button>
    </div>
  </form>
</Modal>

<style>
  .dashboard {
    width: 100%;
    max-width: 600px;
    display: flex;
    flex-direction: column;
    gap: 16px;
  }

  .dashboard__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 4px;
  }

  .dashboard__title {
    font-size: 28px;
    font-weight: 600;
    color: var(--text-high);
    margin: 0;
    letter-spacing: -0.02em;
  }

  .dashboard__actions {
    display: flex;
    align-items: center;
    gap: 8px;
  }

  .dashboard__add {
    display: inline-flex;
    align-items: center;
    gap: 8px;
    height: 40px;
    padding: 0 16px;
    font-family: inherit;
    font-size: 13px;
    font-weight: 500;
    color: var(--surface-ground);
    background-color: var(--accent);
    border: none;
    border-radius: var(--radius);
    cursor: pointer;
    box-shadow: var(--shadow-1);
    transition:
      background-color var(--transition-fast),
      transform var(--transition-fast),
      box-shadow var(--transition-fast);
  }

  .dashboard__add:hover {
    background-color: var(--accent-light);
    transform: translateY(-1px);
    box-shadow: var(--shadow-2);
  }

  .dashboard__add:active {
    background-color: var(--accent-dark);
    transform: translateY(0);
  }

  .dashboard__logout {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 40px;
    height: 40px;
    padding: 0;
    color: var(--text-medium);
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    cursor: pointer;
    transition:
      background-color var(--transition-fast),
      border-color var(--transition-fast),
      color var(--transition-fast);
  }

  .dashboard__logout:hover {
    background-color: var(--surface-elevated);
    border-color: var(--border-hover);
    color: var(--error);
  }

  .dashboard__content {
    min-height: 300px;
  }

  .dashboard__list {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .dashboard__loading {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }

  .dashboard__skeleton {
    height: 80px;
    background: linear-gradient(
      90deg,
      var(--surface-card) 0%,
      var(--surface-elevated) 50%,
      var(--surface-card) 100%
    );
    background-size: 200% 100%;
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    animation: skeleton-shimmer 1.5s ease-in-out infinite;
  }

  @keyframes skeleton-shimmer {
    0% {
      background-position: 200% 0;
    }
    100% {
      background-position: -200% 0;
    }
  }

  .dashboard__error {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 16px;
    padding: 48px 24px;
    color: var(--error);
    text-align: center;
  }

  .dashboard__error button {
    padding: 8px 16px;
    font-family: inherit;
    font-size: 13px;
    color: var(--text-high);
    background-color: var(--surface-elevated);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    cursor: pointer;
  }

  .dashboard__error button:hover {
    background-color: var(--surface-overlay);
  }

  /* Add Friend Form */
  .add-form {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }

  .add-form__error {
    display: flex;
    align-items: center;
    gap: 8px;
    padding: 12px;
    font-size: 13px;
    color: var(--error);
    background-color: rgba(244, 67, 54, 0.1);
    border: 1px solid var(--error);
    border-radius: var(--radius);
  }

  .add-form__actions {
    padding-top: 4px;
  }
</style>
