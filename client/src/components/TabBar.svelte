<script>
  let {
    tabs = [],
    activeTab = $bindable(''),
    onchange = () => {}
  } = $props()

  function handleClick(tabId) {
    activeTab = tabId
    onchange(tabId)
  }

  function getIndicatorStyle() {
    const index = tabs.findIndex(t => t.id === activeTab)
    if (index === -1) return ''
    const width = 100 / tabs.length
    return `transform: translateX(${index * 100}%); width: ${width}%`
  }
</script>

<div class="tabs">
  <div class="tabs__list">
    {#each tabs as tab}
      <button
        class="tabs__tab"
        class:tabs__tab--active={activeTab === tab.id}
        onclick={() => handleClick(tab.id)}
      >
        <span class="tabs__label">{tab.label}</span>
        {#if tab.count > 0}
          <span class="tabs__badge">{tab.count}</span>
        {/if}
      </button>
    {/each}
    <div class="tabs__indicator" style={getIndicatorStyle()}></div>
  </div>
</div>

<style>
  .tabs {
    width: 100%;
    background-color: var(--surface-card);
    border: 1px solid var(--border-default);
    border-radius: var(--radius);
    overflow: hidden;
  }

  .tabs__list {
    position: relative;
    display: flex;
  }

  .tabs__tab {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 8px;
    height: 52px;
    padding: 0 16px;
    font-family: inherit;
    font-size: 13px;
    font-weight: 500;
    letter-spacing: 0.03em;
    text-transform: uppercase;
    color: var(--text-medium);
    background-color: transparent;
    border: none;
    cursor: pointer;
    transition: color var(--transition-fast);
  }

  .tabs__tab:hover {
    color: var(--text-high);
  }

  .tabs__tab--active {
    color: var(--accent);
  }

  .tabs__label {
    position: relative;
  }

  .tabs__badge {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    min-width: 20px;
    height: 20px;
    padding: 0 6px;
    font-size: 11px;
    font-weight: 600;
    color: var(--surface-ground);
    background-color: var(--accent);
    border-radius: 10px;
    animation: badge-pop 200ms cubic-bezier(0.34, 1.56, 0.64, 1);
  }

  .tabs__indicator {
    position: absolute;
    bottom: 0;
    left: 0;
    height: 2px;
    background: linear-gradient(90deg, var(--accent-dark), var(--accent), var(--accent-light));
    transition: transform var(--transition-normal);
  }

  @keyframes badge-pop {
    0% {
      transform: scale(0);
    }
    100% {
      transform: scale(1);
    }
  }
</style>
