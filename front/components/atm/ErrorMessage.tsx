import styles from '../../styles/atm/ErrorMessage.module.scss';

interface ErrorMessageProps {
  children?: React.ReactNode;
}

const ErrorMessage = ({
  children
}: ErrorMessageProps) => {
  return (
    <div className={styles.message}>
      {children}
    </div>
  )
}

export default ErrorMessage;
