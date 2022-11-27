import styles from '../../styles/atm/Input.module.scss'

interface InputProps {
  value?: string;
  placeholder?: string;
  className?: string;
  type?: 'text' | 'password' | 'email' | 'search';
  pattern?: string;
  readonly?: boolean;
  disabled?: boolean;
  onChange?(e: React.ChangeEvent<HTMLInputElement>): void;
  onInput?(e: React.FormEvent<HTMLInputElement>): void;
  onFocus?(e: React.FocusEvent<HTMLInputElement>): void;
}

const Input = ({
  type = 'text',
  placeholder = '',
  className,
  ...props
}: InputProps) => {
  return (
    <input
      type={type}
      placeholder={placeholder}
      className={`${className} ${styles.input}`}
      {...props}
    />
  )
}

export default Input;
