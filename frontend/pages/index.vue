<script setup lang="ts">
const currentYear = ref('')
const hours = ref('00')
const minutes = ref('00')
const seconds = ref('00')
const milliseconds = ref('000')

let timer: ReturnType<typeof setInterval> | null = null

const startTime = Date.now()

const updateCounter = () => {
  const elapsed = Date.now() - startTime

  const h = Math.floor(elapsed / (1000 * 60 * 60))
  const m = Math.floor((elapsed / (1000 * 60)) % 60)
  const s = Math.floor((elapsed / 1000) % 60)
  const ms = elapsed % 1000

  hours.value = h.toString().padStart(2, '0')
  minutes.value = m.toString().padStart(2, '0')
  seconds.value = s.toString().padStart(2, '0')
  milliseconds.value = ms.toString().padStart(3, '0')
}

onMounted(() => {
  currentYear.value = new Date().getFullYear().toString()
  updateCounter()
  timer = setInterval(updateCounter, 50)
})

onBeforeUnmount(() => {
  if (timer) {
    clearInterval(timer)
  }
})
</script>

<template>
  <div class="page">
    <section class="hero">
      <div class="hero-bg" />
      <div class="hero-overlay" />

      <div class="hero-content">
        <img
          class="logo"
          src="/kubernetes-logo.svg"
          alt="Kubernetes Logo"
          width="180"
          height="180"
          fetchpriority="high"
        >
        <h1 class="title">
          КУБЫПОБЕДА
        </h1>
        <p class="subtitle">
          Очередной год победы Kubernetes
        </p>
      </div>

      <a class="scroll-down" href="#main-content" aria-label="Перейти к основному содержимому">
        <svg width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true">
          <path d="M12 5v14M19 12l-7 7-7-7" />
        </svg>
      </a>
    </section>

    <main id="main-content" class="main">
      <section class="status-section">
        <h2 class="status-title">
          Статус: SYNCING
        </h2>
        <p class="status-subtitle">
          Почему {{ currentYear }} это год кубов?
        </p>

        <div class="cards">
          <article class="card">
            <div class="card-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
                <path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4" />
                <path d="M14 2v6h6" />
                <path d="M2 15h10" />
                <path d="M9 12v6" />
              </svg>
            </div>
            <h3>YAML Native</h3>
            <p>Родной язык конфигурации. Отступы важнее смысла, а количество строк измеряется тысячами.</p>
          </article>

          <article class="card">
            <div class="card-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
                <circle cx="12" cy="12" r="10" />
                <path d="M12 8v4" />
                <path d="M12 16h.01" />
              </svg>
            </div>
            <h3>OOMKilled</h3>
            <p>Не ошибка, а запланированная процедура очистки памяти. Решение всегда одно - увеличить лимиты.</p>
          </article>

          <article class="card">
            <div class="card-icon">
              <svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" aria-hidden="true">
                <rect x="2" y="2" width="20" height="8" rx="2" />
                <rect x="2" y="14" width="20" height="8" rx="2" />
                <path d="M6 6h.01" />
                <path d="M6 18h.01" />
              </svg>
            </div>
            <h3>Масштаб</h3>
            <p>Зачем один сервер, если можно поднять целый кластер ради статического сайта? (но, увы, не в этом случае)</p>
          </article>
        </div>
      </section>

      <section class="counter-section">
        <h2 class="counter-title">
          Стабильность кластера
        </h2>

        <div class="counter-grid">
          <div class="counter-item">
            <span class="counter-value">{{ hours }}</span>
            <span class="counter-label">Часов</span>
          </div>
          <div class="counter-item">
            <span class="counter-value">{{ minutes }}</span>
            <span class="counter-label">Минут</span>
          </div>
          <div class="counter-item">
            <span class="counter-value">{{ seconds }}</span>
            <span class="counter-label">Секунд</span>
          </div>
          <div class="counter-item">
            <span class="counter-value">{{ milliseconds }}</span>
            <span class="counter-label">мс</span>
          </div>
        </div>

        <p class="counter-note">
          Времени прошло с момента последнего Pending пода
        </p>
      </section>
    </main>

    <footer class="footer">
      <p>Крутится на Github Pages (денег на кубовый кластер не хватило)</p>
      <p>
        Сурсы на
        <a class="footer-link" href="https://github.com/misshanya/k8s-pobeda">GitHub</a>
      </p>
    </footer>
  </div>
</template>
