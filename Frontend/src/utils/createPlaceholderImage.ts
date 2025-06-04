/**
 * Утилита для создания placeholder изображений в случае отсутствия настоящих
 */

/**
 * Безопасное кодирование в base64 с поддержкой Unicode символов
 * @param str Строка для кодирования
 * @returns Закодированная строка
 */
const encodeUnicode = (str: string): string => {
  // Преобразуем UTF-8 строку в безопасную для base64
  return window.btoa(
    encodeURIComponent(str).replace(/%([0-9A-F]{2})/g, (_, p1) => {
      return String.fromCharCode(parseInt(p1, 16));
    })
  );
};

// SVG иконка игрового контроллера для плейсхолдера игры
const gameControllerIcon = `
  <path d="M36,18c-6.7,0-12.6,0-18,0C11.9,18,7,22.9,7,29s4.9,11,11,11h36c6.1,0,11-4.9,11-11s-4.9-11-11-11H36z M16,33h-4v4h-3
    v-4H5v-3h4v-4h3v4h4V33z M37,35c-1.7,0-3-1.3-3-3s1.3-3,3-3s3,1.3,3,3S38.7,35,37,35z M47,35c-1.7,0-3-1.3-3-3s1.3-3,3-3
    s3,1.3,3,3S48.7,35,47,35z" fill="currentColor"/>
`;

// SVG иконка газеты для плейсхолдера новостей
const newspaperIcon = `
  <path d="M48,12H16c-2.2,0-4,1.8-4,4v32c0,2.2,1.8,4,4,4h32c2.2,0,4-1.8,4-4V16C52,13.8,50.2,12,48,12z M46,44c0,1.1-0.9,2-2,2H20
    c-1.1,0-2-0.9-2-2V20c0-1.1,0.9-2,2-2h24c1.1,0,2,0.9,2,2V44z" fill="currentColor"/>
  <rect x="22" y="22" width="8" height="8" fill="currentColor"/>
  <rect x="34" y="22" width="8" height="2" fill="currentColor"/>
  <rect x="34" y="26" width="8" height="2" fill="currentColor"/>
  <rect x="22" y="34" width="20" height="2" fill="currentColor"/>
  <rect x="22" y="38" width="20" height="2" fill="currentColor"/>
`;

/**
 * Создает URL для placeholder изображения с текстом
 * @param width Ширина изображения
 * @param height Высота изображения
 * @param text Текст для отображения на изображении
 * @param backgroundColor Цвет фона
 * @param textColor Цвет текста
 * @param iconSvgPath SVG путь для иконки (опционально)
 * @returns URL для использования в src тега img
 */
export const createPlaceholderImage = (
  width: number,
  height: number,
  text: string,
  backgroundColor: string = '#2d3e50',
  textColor: string = '#ffffff',
  iconSvgPath?: string
): string => {
  try {
    // Обрезаем слишком длинный текст
    const safeText = text.length > 20 ? `${text.substring(0, 20)}...` : text;
    
    // Определяем, нужно ли добавлять иконку
    const hasIcon = !!iconSvgPath;
    
    // Создаем SVG изображение
    const svg = `
      <svg width="${width}" height="${height}" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ${width} ${height}">
        <rect width="100%" height="100%" fill="${backgroundColor}"/>
        ${hasIcon ? `
          <g transform="translate(${width/2 - 32}, ${height/2 - 40})">
            <svg width="64" height="64" viewBox="0 0 64 64" fill="${textColor}">
              ${iconSvgPath}
            </svg>
          </g>
        ` : ''}
        <text 
          x="50%" 
          y="${hasIcon ? '70%' : '50%'}" 
          font-family="Arial, sans-serif" 
          font-size="20" 
          text-anchor="middle" 
          dominant-baseline="middle" 
          fill="${textColor}"
        >
          ${safeText}
        </text>
      </svg>
    `;

    // Кодируем SVG в base64 для использования в URL с поддержкой Unicode
    const encoded = encodeUnicode(svg);
    return `data:image/svg+xml;base64,${encoded}`;
  } catch (error) {
    console.error('Ошибка при создании placeholder изображения:', error);
    // Возвращаем базовое изображение при ошибке
    return `data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iJHt3aWR0aH0iIGhlaWdodD0iJHtoZWlnaHR9IiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPjxyZWN0IHdpZHRoPSIxMDAlIiBoZWlnaHQ9IjEwMCUiIGZpbGw9IiMxYTI2MzQiLz48L3N2Zz4=`;
  }
};

/**
 * Создает URL для placeholder изображения новости
 * @param width Ширина изображения
 * @param height Высота изображения
 * @param text Текст для отображения на изображении
 * @returns URL для использования в src тега img
 */
export const createNewsPlaceholderImage = (
  width: number,
  height: number,
  text: string = 'Новости'
): string => {
  return createPlaceholderImage(width, height, text, '#1a2634', '#ffffff', newspaperIcon);
};

/**
 * Создает URL для placeholder изображения игры
 * @param width Ширина изображения
 * @param height Высота изображения
 * @param text Текст для отображения на изображении
 * @returns URL для использования в src тега img
 */
export const createGamePlaceholderImage = (
  width: number,
  height: number,
  text: string = 'Игра'
): string => {
  return createPlaceholderImage(width, height, text, '#1a2634', '#ffce20', gameControllerIcon);
};

export default createPlaceholderImage; 