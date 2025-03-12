import { SVGProps } from "react";

export type IconSvgProps = SVGProps<SVGSVGElement> & {
  size?: number;
};

export interface SearchIconProps {
  size?: number;
  strokeWidth?: number;
  width?: number | string; 
  height?: number | string; 
}