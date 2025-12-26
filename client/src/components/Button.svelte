<script>
  let {
    type = 'button',
    variant = 'primary',
    loading = false,
    disabled = false,
    onclick = () => {},
    children
  } = $props()
</script>

<button
  class="btn btn--{variant}"
  class:btn--loading={loading}
  {type}
  disabled={disabled || loading}
  {onclick}
>
  <span class="btn__content">
    {@render children()}
  </span>
  {#if loading}
    <span class="btn__loader">
      <span class="btn__loader-bar"></span>
    </span>
  {/if}
</button>

<style>
  .btn {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 48px;
    padding: 0 24px;
    font-family: inherit;
    font-size: 14px;
    font-weight: 500;
    letter-spacing: 0.02em;
    text-transform: uppercase;
    border: none;
    border-radius: var(--radius);
    cursor: pointer;
    overflow: hidden;
    transition:
      background-color var(--transition-fast),
      box-shadow var(--transition-fast),
      transform var(--transition-fast);
  }

  .btn:disabled {
    opacity: 0.5;
    cursor: not-allowed;
    transform: none !important;
    box-shadow: none !important;
  }

  /* Primary variant */
  .btn--primary {
    background-color: var(--accent);
    color: var(--surface-ground);
    box-shadow: var(--shadow-1);
  }

  .btn--primary:hover:not(:disabled) {
    background-color: var(--accent-light);
    box-shadow: var(--shadow-2);
    transform: translateY(-1px);
  }

  .btn--primary:active:not(:disabled) {
    background-color: var(--accent-dark);
    box-shadow: var(--shadow-1);
    transform: translateY(0);
  }

  /* Secondary variant */
  .btn--secondary {
    background-color: transparent;
    color: var(--text-high);
    border: 1px solid var(--border-default);
  }

  .btn--secondary:hover:not(:disabled) {
    background-color: var(--surface-elevated);
    border-color: var(--border-hover);
  }

  .btn--secondary:active:not(:disabled) {
    background-color: var(--surface-overlay);
  }

  /* Ghost variant */
  .btn--ghost {
    background-color: transparent;
    color: var(--accent);
  }

  .btn--ghost:hover:not(:disabled) {
    background-color: rgba(0, 188, 212, 0.08);
  }

  .btn--ghost:active:not(:disabled) {
    background-color: rgba(0, 188, 212, 0.16);
  }

  /* Content */
  .btn__content {
    position: relative;
    z-index: 1;
    display: flex;
    align-items: center;
    gap: 8px;
    transition: opacity var(--transition-fast);
  }

  .btn--loading .btn__content {
    opacity: 0;
  }

  /* Loader */
  .btn__loader {
    position: absolute;
    inset: 0;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .btn__loader-bar {
    width: 24px;
    height: 2px;
    background-color: currentColor;
    animation: loader 1s ease-in-out infinite;
  }

  @keyframes loader {
    0%, 100% {
      transform: scaleX(0.3);
      opacity: 0.4;
    }
    50% {
      transform: scaleX(1);
      opacity: 1;
    }
  }
</style>
