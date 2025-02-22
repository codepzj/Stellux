import { IconSvgProps } from "@/types/icons";

export const SunLinearIcon = ({
  size = 24,
  width,
  height,
  ...props
}: IconSvgProps) => (
  <svg
    aria-hidden="true"
    fill="none"
    focusable="false"
    height={size || height}
    role="presentation"
    viewBox="0 0 24 24"
    width={size || width}
    xmlns="http://www.w3.org/2000/svg"
    {...props}
  >
    <path
      d="M12 17C14.7614 17 17 14.7614 17 12C17 9.23858 14.7614 7 12 7C9.23858 7 7 9.23858 7 12C7 14.7614 9.23858 17 12 17Z"
      stroke="currentColor"
      strokeWidth="1.5"
    />
    <path
      d="M12 2V4M12 20V22M4 12H2M22 12H20M19.778 4.223L17.556 6.254M4.222 4.223L6.444 6.254M6.444 17.556L4.222 19.778M19.778 19.777L17.556 17.555"
      stroke="currentColor"
      strokeLinecap="round"
      strokeWidth="1.5"
    />
  </svg>
);
