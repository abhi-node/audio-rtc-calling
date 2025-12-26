<script>
  let {
    label = '',
    type = 'text',
    value = $bindable(''),
    placeholder = '',
    required = false,
    autocomplete = ''
  } = $props()
</script>

<label class="field">
  {#if label}
    <span class="field-label">{label}</span>
  {/if}
  <div class="field-input-wrapper">
    <input
      class="field-input"
      {type}
      bind:value
      {placeholder}
      {required}
      {autocomplete}
    />
    <div class="field-border"></div>
  </div>
</label>

<style>
  .field {
    display: block;
    width: 100%;
  }

  .field-label {
    display: block;
    font-size: 12px;
    font-weight: 500;
    letter-spacing: 0.05em;
    text-transform: uppercase;
    color: var(--text-medium);
    margin-bottom: 8px;
  }

  .field-input-wrapper {
    position: relative;
    width: 100%;
  }

  .field-input {
    display: block;
    width: 100%;
    height: 48px;
    padding: 0 16px;
    font-family: inherit;
    font-size: 14px;
    font-weight: 400;
    color: var(--text-high);
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    outline: none;
    transition:
      border-color var(--transition-fast),
      background-color var(--transition-fast),
      box-shadow var(--transition-fast);
  }

  .field-input::placeholder {
    color: var(--text-low);
  }

  .field-input:hover {
    border-color: var(--border-hover);
    background-color: var(--surface-elevated);
  }

  .field-input:focus {
    border-color: var(--accent);
    background-color: var(--surface-elevated);
    box-shadow: inset 0 0 0 1px var(--accent);
  }

  .field-border {
    position: absolute;
    bottom: 0;
    left: 0;
    right: 0;
    height: 2px;
    background-color: var(--accent);
    transform: scaleX(0);
    transform-origin: center;
    transition: transform var(--transition-normal);
  }

  .field-input:focus + .field-border {
    transform: scaleX(1);
  }
</style>
