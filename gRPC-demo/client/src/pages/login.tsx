import type { NextPage } from 'next'
import { useContext } from 'react'
import { UserContext } from '../context/AuthProvider'
import { Container } from '../components/Container'

const LoginPage: NextPage = () => {
  const { login } = useContext(UserContext)

  return (
    <Container>
      <button onClick={login}>slack login</button>
    </Container>
  )
}

export default LoginPage
