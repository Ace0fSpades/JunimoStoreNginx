import React from 'react';

interface JunimoLogoProps {
  color?: string;
  strokeColor?: string;
  size?: number;
  className?: string;
  strokeWidth?: number;
}

const JunimoLogo: React.FC<JunimoLogoProps> = ({
  color = 'none',
  strokeColor = '#000',
  size = 48,
  className = '',
  strokeWidth = 2,
}) => {
  const svgStyle = {
    width: size,
    height: size,
  };

  return (
    <svg
      style={svgStyle}
      viewBox="0 0 106 100"
      fill={color}
      xmlns="http://www.w3.org/2000/svg"
      className={className}
    >
      <path
        d="M49.245 18.128c10.516 0 20.158 5.706 27.218 14.239 7.061 8.533 11.409 19.747 11.409 30.446 0 5.338-1.082 9.387-2.94 12.481-1.853 3.086-4.546 5.334-7.96 6.969-6.92 3.312-16.63 4.035-27.727 4.035s-20.809-.723-27.728-4.035c-3.415-1.635-6.107-3.883-7.96-6.969-1.86-3.094-2.94-7.143-2.94-12.481 0-10.699 4.348-21.913 11.409-30.446 7.06-8.534 16.702-14.239 27.219-14.239Z"
        stroke={strokeColor}
        strokeWidth={strokeWidth}
      />
      <path
        d="M50.328 17.9C46.543 11.086 47.3 9.57 42 2m46 54.778s9.085-4.543 15.9 8.328M9.63 59.784S-.968 48.427 2.818 37.828"
        stroke={strokeColor}
        strokeWidth={strokeWidth}
        strokeLinecap="round"
      />
      <circle cx="34.859" cy="60.541" fill={strokeColor} r="6.057" />
      <circle cx="69.687" cy="60.541" fill={strokeColor} r="6.057" />
      <path
        d="M62.873 86.526s0 5.3 3.028 9.086c3.029 3.785 9.086 2.271 9.086 2.271M34.859 86.77s0 5.3-3.029 9.085-9.085 2.271-9.085 2.271"
        stroke={strokeColor}
        strokeWidth={strokeWidth}
        strokeLinecap="round"
      />
    </svg>
  );
};

export default JunimoLogo;