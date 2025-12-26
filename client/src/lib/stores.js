import { writable, derived } from 'svelte/store'

const STORAGE_KEY = 'auth_token'

function createTokenStore() {
  const stored = typeof localStorage !== 'undefined'
    ? localStorage.getItem(STORAGE_KEY)
    : null

  const { subscribe, set, update } = writable(stored)

  return {
    subscribe,
    set: (value) => {
      if (typeof localStorage !== 'undefined') {
        if (value) {
          localStorage.setItem(STORAGE_KEY, value)
        } else {
          localStorage.removeItem(STORAGE_KEY)
        }
      }
      set(value)
    },
    clear: () => {
      if (typeof localStorage !== 'undefined') {
        localStorage.removeItem(STORAGE_KEY)
      }
      set(null)
    }
  }
}

export const token = createTokenStore()
export const isAuthenticated = derived(token, $token => !!$token)
