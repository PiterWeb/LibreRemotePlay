import { browser } from '$app/environment'
import log from '$lib/logger/logger'
import { init, register, getLocaleFromNavigator, locale} from 'svelte-i18n'

const defaultLocale = 'en'

register('en', () => import('./en.json'))
register('es', () => import('./es.json'))
register('gl', () => import('./gl.json'))
register('ru', () => import('./ru.json'))
register('fr', () => import('./fr.json'))

init({
	fallbackLocale: defaultLocale,
	initialLocale: browser ? new Intl.Locale(getLocaleFromLocalStorage() ?? getLocaleFromNavigator() ?? defaultLocale).language : defaultLocale,
})

export function getLocaleFromLocalStorage() {
	return localStorage.getItem('locale')
}

function saveLocaleToLocalStorage(locale: string) {
	browser && localStorage.setItem('locale', locale)
}

locale.subscribe((value) => {
  log(`locale changed to ${value}`)
	saveLocaleToLocalStorage(value ?? defaultLocale)
})


