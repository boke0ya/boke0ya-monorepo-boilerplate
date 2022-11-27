import Head from "next/head"
import SignupTemplate from "../components/tmp/SignupTemplate"
import useSignupViewModel from "../hooks/viewmodels/useSignupViewModel"

const Signup = () => {
  const viewmodel = useSignupViewModel()
  return (
    <>
      <Head>
        <title>新規登録</title>
      </Head>
      <SignupTemplate viewmodel={viewmodel} />
    </>
  )
}

export default Signup
