import styles from '../../styles/atm/FormLabel.module.scss'

interface FormLabelProps {
  children: React.ReactNode;
}

const FormLabel = ({
  children
}: FormLabelProps) => {
  return (
    <div className={styles.label}>
      {children}
    </div>
  )
}

export default FormLabel;
