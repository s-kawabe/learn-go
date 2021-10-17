import type { FC } from 'react'
import { ReactNode } from 'react'
import styles from './Container.module.css'

type Props = {
  children: ReactNode
}

export const Container: FC<Props> = (props) => {
  return (
    <div className={styles.container}>
      {props.children}
    </div>
  )
}
