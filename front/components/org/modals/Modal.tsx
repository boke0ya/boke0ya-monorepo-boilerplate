import styles from '../../../styles/org/modals/Modal.module.scss'

interface ModalProps {
  isOpen: boolean;
  title: string;
  className?: string;
  width?: number;
  onClose?(): void;
  children: React.ReactNode;
}

const Modal = ({
  isOpen,
  title,
  className,
  onClose,
  children,
  width = 480 
}: ModalProps) => {
  return (
    <div className={`${styles.container} ${isOpen ? styles.open : ''}`}>
      <div className={styles.overlay} onClick={onClose}/>
      <div className={`${className} ${styles.main}`} style={{
        width: `${width}px`
      }}>
        <h3 className={styles.title}>{title}</h3>
        {children}
      </div>
    </div>
  )
}

export default Modal

interface ModalFooterProps {
  children: React.ReactNode;
}

export const ModalFooter = ({ children }: ModalFooterProps) => {
  return (
    <div className={styles.footer}>
      {children}
    </div>
  )
}
