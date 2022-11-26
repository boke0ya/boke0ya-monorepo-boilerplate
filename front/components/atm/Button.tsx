import styles from '../../styles/atm/Button.module.scss'

interface ButtonProps {
  children: React.ReactNode;
  className?: string;
  disabled?: boolean;
  isLoading?: boolean;
  onClick?(e: React.MouseEvent): void;
}

const Button = ({
  children,
  className,
  disabled,
  isLoading,
  ...props
}: ButtonProps) => {
  return (
    <button
      className={`${className} ${styles.button}`}
      disabled={disabled || isLoading}
      {...props}
    >
      {
        isLoading ? (
          <div className={styles.loading} />
        ) : (<></>)
      }
      {children}
    </button>
  )
}

export default Button;
