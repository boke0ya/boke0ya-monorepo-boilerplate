import '../styles/globals.scss'
import type { AppProps } from 'next/app'
import useGlobalViewModel, { GlobalContext } from '../hooks/viewmodels/useGlobalViewModel'
import { useEffect } from 'react'

export default function App({ Component, pageProps }: AppProps) {
  const globalViewModel = useGlobalViewModel()
  useEffect(() => {
    globalViewModel.loadSessionUser()
  }, [])
  return (
    <GlobalContext.Provider value={globalViewModel}>
      <Component {...pageProps} />
    </GlobalContext.Provider>
  )
}
