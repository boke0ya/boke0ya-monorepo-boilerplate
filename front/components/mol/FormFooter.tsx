import styles from '../../styles/mol/FormFooter.module.scss'

interface FormFooterProps {
  children: React.ReactNode;
}

const FormFooter = ({
  children
}: FormFooterProps) => {
  return (
    <div className={styles.container}>
      {children}
    </div>
  )
}

export default FormFooter
