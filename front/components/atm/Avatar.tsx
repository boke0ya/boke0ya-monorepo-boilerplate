import Image from "next/image";
import styles from '../../styles/atm/Avatar.module.scss';
import { BiUser } from 'react-icons/bi';

interface AvatarProps {
  url: string;
  size?: number;
}

const Avatar = ({
  url,
  size = 64
}: AvatarProps) => {
  return (
    <div className={styles.container} style={{
      width: `${size}px`,
      height: `${size}px`,
      borderRadius: `${size / 2}px`,
    }}>
      {
        url ? (
          <Image src={url} alt={'Profile icon'} fill unoptimized />
        ) : (
          <BiUser />
        )
      }
    </div>
  )
}

export default Avatar
