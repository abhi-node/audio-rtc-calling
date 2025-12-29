import { writable, derived } from 'svelte/store';
import { getFriends, getIncomingRequests, getOutgoingRequests } from './api.js';

function createFriendsStore() {
  const friends = writable([]);
  const incoming = writable([]);
  const outgoing = writable([]);
  const loading = writable(true);
  const error = writable(null);

  let pollInterval = null;
  let isPolling = false;

  async function fetchAll() {
    error.set(null);
    try {
      const [friendsRes, incomingRes, outgoingRes] = await Promise.all([
        getFriends(),
        getIncomingRequests(),
        getOutgoingRequests()
      ]);
      friends.set(friendsRes.requests || []);
      incoming.set(incomingRes.requests || []);
      outgoing.set(outgoingRes.requests || []);
    } catch (e) {
      error.set(e.message);
    } finally {
      loading.set(false);
    }
  }

  function startPolling(intervalMs = 10000) {
    if (isPolling) return;
    isPolling = true;

    fetchAll();

    pollInterval = setInterval(() => {
      if (document.visibilityState === 'visible') {
        fetchAll();
      }
    }, intervalMs);

    const handleVisibility = () => {
      if (document.visibilityState === 'visible') {
        fetchAll();
      }
    };
    document.addEventListener('visibilitychange', handleVisibility);

    return () => {
      clearInterval(pollInterval);
      document.removeEventListener('visibilitychange', handleVisibility);
      isPolling = false;
    };
  }

  function stopPolling() {
    if (pollInterval) {
      clearInterval(pollInterval);
      pollInterval = null;
    }
    isPolling = false;
  }

  async function refetch() {
    loading.set(true);
    await fetchAll();
  }

  const incomingCount = derived(incoming, $incoming => $incoming.length);
  const outgoingCount = derived(outgoing, $outgoing => $outgoing.length);

  return {
    friends,
    incoming,
    outgoing,
    loading,
    error,
    incomingCount,
    outgoingCount,
    startPolling,
    stopPolling,
    refetch
  };
}

export const friendsStore = createFriendsStore();
