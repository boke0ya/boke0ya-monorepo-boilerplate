import styles from '../../styles/atm/IconButton.module.scss'

interface IconButtonProps {
  size?: number;
  className?: string;
  disabled?: boolean;
  children: React.ReactNode;
  onClick?(e: React.MouseEvent): void;
}

const IconButton = ({
  size = 32,
  className,
  children,
  ...props
}: IconButtonProps) => {
  return (
    <button className={`${className} ${styles.button}`} style={{
      width: `${size}px`,
      height: `${size}px`,
    }} {...props}>
      {children}
    </button>
  )
}

export default IconButton
