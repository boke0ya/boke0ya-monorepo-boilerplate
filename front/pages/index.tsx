import Head from 'next/head'
import { useState } from 'react'
import LoginModal from '../components/org/modals/LoginModal'
import Modal from '../components/org/modals/Modal'
import SignupModal from '../components/org/modals/SignupModal'
import styles from '../styles/Home.module.scss'

export default function Home() {
  const [isLoginOpen, setIsLoginOpen] = useState(false)
  const [isSignupOpen, setIsSignupOpen] = useState(false)
  return (
    <>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <button onClick={() => setIsLoginOpen(true)}>ログイン</button>
      <button onClick={() => setIsSignupOpen(true)}>新規登録</button>
      <LoginModal isOpen={isLoginOpen} onClose={() => setIsLoginOpen(false)} />
      <SignupModal isOpen={isSignupOpen} onClose={() => setIsSignupOpen(false)} />
    </>
  )
}
