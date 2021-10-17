import { createContext, ReactNode, useState } from 'react'
import type { FC } from 'react'

type User = {
  id: string
  name: string
  icon: string
}

export type UserContextType = {
  user: User
  isLogin: boolean
  login: () => void
  // logout: () => void
  // getToken: () => void
}

const initialUser: User = {
  id: "",
  name: "",
  icon: "",
}

export const UserContext = createContext<UserContextType>({
  user: initialUser,
  isLogin: false,
  login: () => {
    return new Promise<string>((resolve) => {
      resolve("");
    });
  },
  // logout: () => {
  //   return new Promise<void>((resolve) => {
  //     resolve();
  //   });
  // },
  // getToken: () => {
  //   return new Promise<void>((resolve) => {
  //     resolve();
  //   });
  // },
})

const login = () => {
  alert('your successed login!')
}

type Props = {
  children: ReactNode
}

export const AuthProvider: FC<Props> = (props) => {
  const [user, setUser] = useState<User>(initialUser)
  const [isLogin, setIsLogin] = useState(false)

  return (
    <UserContext.Provider value={{ user, isLogin, login }}>
      {props.children}
    </UserContext.Provider>
  )
}
