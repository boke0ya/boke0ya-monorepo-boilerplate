import '../styles/globals.scss'
import type { AppProps } from 'next/app'
import { GlobalContext, GlobalReducer } from '../hooks/viewmodels/useGlobalViewModel'
import { useReducer } from 'react'

export default function App({ Component, pageProps }: AppProps) {
  const [state, dispatch] = useReducer(GlobalReducer, {
    sessionUser: null,
  })
  return (
    <GlobalContext.Provider value={{
      state,
      dispatch
    }}>
      <Component {...pageProps} />
    </GlobalContext.Provider>
  )
}
