<script>
  let {
    open = $bindable(false),
    title = '',
    children
  } = $props()

  let dialogRef = $state(null)

  function handleBackdrop(e) {
    if (e.target === dialogRef) {
      open = false
    }
  }

  function handleKeydown(e) {
    if (e.key === 'Escape') {
      open = false
    }
  }
</script>

{#if open}
  <!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
  <div
    class="modal"
    bind:this={dialogRef}
    onclick={handleBackdrop}
    onkeydown={handleKeydown}
    role="dialog"
    aria-modal="true"
    tabindex="-1"
  >
    <div class="modal__content">
      <div class="modal__header">
        <h2 class="modal__title">{title}</h2>
        <button class="modal__close" onclick={() => open = false} aria-label="Close">
          <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round">
            <line x1="18" y1="6" x2="6" y2="18"/>
            <line x1="6" y1="6" x2="18" y2="18"/>
          </svg>
        </button>
      </div>
      <div class="modal__body">
        {@render children()}
      </div>
    </div>
  </div>
{/if}

<style>
  .modal {
    position: fixed;
    inset: 0;
    z-index: 1000;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 24px;
    background-color: rgba(0, 0, 0, 0.7);
    backdrop-filter: blur(4px);
    animation: modal-backdrop 200ms ease-out;
  }

  @keyframes modal-backdrop {
    0% {
      opacity: 0;
    }
    100% {
      opacity: 1;
    }
  }

  .modal__content {
    width: 100%;
    max-width: 400px;
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    box-shadow: var(--shadow-3);
    animation: modal-slide 300ms cubic-bezier(0.34, 1.56, 0.64, 1);
  }

  @keyframes modal-slide {
    0% {
      opacity: 0;
      transform: translateY(24px) scale(0.95);
    }
    100% {
      opacity: 1;
      transform: translateY(0) scale(1);
    }
  }

  .modal__header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 20px 24px;
    border-bottom: 1px solid var(--border-default);
  }

  .modal__title {
    font-size: 16px;
    font-weight: 500;
    color: var(--text-high);
    margin: 0;
  }

  .modal__close {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;
    padding: 0;
    background-color: transparent;
    border: none;
    border-radius: var(--radius);
    color: var(--text-medium);
    cursor: pointer;
    transition:
      background-color var(--transition-fast),
      color var(--transition-fast);
  }

  .modal__close:hover {
    background-color: var(--surface-elevated);
    color: var(--text-high);
  }

  .modal__body {
    padding: 24px;
  }
</style>
