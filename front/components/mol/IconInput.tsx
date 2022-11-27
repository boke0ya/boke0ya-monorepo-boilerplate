import { useEffect, useRef, useState } from 'react'
import styles from '../../styles/mol/IconInput.module.scss'
import Avatar from '../atm/Avatar'
import Button from '../atm/Button';
import IconButton from '../atm/IconButton';
import { GrPowerReset } from 'react-icons/gr'

interface IconInputProps {
  defaultUrl?: string;
  value?: File | Blob;
  onChange?(file: File | Blob | null): void;
}

const IconInput = ({
  defaultUrl,
  value,
  onChange,
}: IconInputProps) => {
  const inputRef = useRef<HTMLInputElement>();
  const [url, setUrl] = useState<string>(null)
  const change = (e: React.ChangeEvent<HTMLInputElement>) => {
    if(e.target.files.length > 0){
      onChange(e.target.files[0])
    }
    e.target.value = null
  }
  useEffect(() => {
    setUrl(url => {
      if(url) URL.revokeObjectURL(url)
      if(value){
          return URL.createObjectURL(value)
      }
      return null
    })
  }, [value])
  return (
    <div className={styles.container}>
      <input type='file' onChange={change} ref={inputRef} style={{ display: 'none' }} accept='image/*' />
      <Avatar url={url ?? defaultUrl} />
      <Button onClick={() => inputRef.current.click()}>プロフィール画像を変更</Button>
      <div style={{ opacity: value ? 1 : 0 }}>
        <IconButton size={24} onClick={() => onChange(null)}>
          <GrPowerReset />
        </IconButton>
      </div>
    </div>
  )
}

export default IconInput
