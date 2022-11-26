import { ComponentProps } from "react";
import FormLabel from "../atm/FormLabel";
import Input from "../atm/Input";
import styles from '../../styles/mol/TextInput.module.scss'

interface TextInputProps extends ComponentProps<typeof Input> {
  label: string;
}

const TextInput = ({
  label,
  ...props
}: TextInputProps) => {
  return (
    <label className={styles.container}>
      <FormLabel>{label}</FormLabel>
      <Input {...props} />
    </label>
  )
}

export default TextInput;
