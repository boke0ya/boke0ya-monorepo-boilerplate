import styles from '@/styles/utils/Box.module.scss';

interface BoxProps {
  children: React.ReactNode;
}

const Box = ({
  children
}: BoxProps) => {
  return (
    <div className={styles.container}>
      {children}
    </div>
  )
}

export default Box;
