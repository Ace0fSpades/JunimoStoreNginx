import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import styles from '../styles/pages/CartPage.module.scss';
import { FaTrash, FaArrowLeft, FaCreditCard } from 'react-icons/fa';

// Services
import { CartService } from '../services/CartService';
import { OrderService } from '../services/OrderService';

// Context
import { useAuth } from '../context/AuthContext';

// Types
import { ShoppingCart, CartItem } from '../models/Cart';

const CartPage: React.FC = () => {
  const { user } = useAuth();
  const [cart, setCart] = useState<ShoppingCart | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [orderLoading, setOrderLoading] = useState(false);
  const [orderSuccess, setOrderSuccess] = useState(false);
  const [orderError, setOrderError] = useState<string | null>(null);

  // Calculate cart totals
  const calculateTotals = () => {
    if (!cart?.cartItems?.length) {
      return { totalItems: 0, subtotal: 0 };
    }
    
    const totalItems = cart.cartItems.reduce((sum, item) => sum + item.quantity, 0);
    const subtotal = cart.cartItems.reduce((sum, item) => sum + (item.game.price * item.quantity), 0);
    
    return { totalItems, subtotal };
  };
  
  const { totalItems, subtotal } = calculateTotals();

  // Проверка на пустую корзину
  const isCartEmpty = !cart?.cartItems?.length;

  useEffect(() => {
    if (user) {
      fetchCart();
    } else {
      // Если пользователь не авторизован, устанавливаем загрузку в false
      setLoading(false);
    }
  }, [user]);

  const fetchCart = async () => {
    if (!user) {
      setLoading(false);
      return;
    }
    
    try {
      setLoading(true);
      const cartData = await CartService.getCart(user.id);
      setCart(cartData);
      setError(null);
    } catch (err) {
      console.error('Error fetching cart:', err);
      setError('Не удалось загрузить корзину. Пожалуйста, попробуйте позже.');
      // Устанавливаем cart в пустой объект, чтобы показать пустую корзину вместо загрузки
      setCart({ id: 0, userID: user.id, cartItems: [] });
    } finally {
      setLoading(false);
    }
  };

  const handleQuantityChange = async (item: CartItem, newQuantity: number) => {
    if (newQuantity < 1) {
      return;
    }
    
    try {
      await CartService.updateQuantity(user!.id, item.game.id, newQuantity);
      // Update local state for immediate UI update
      setCart(prevCart => {
        if (!prevCart) return null;
        
        return {
          ...prevCart,
          cartItems: prevCart.cartItems.map(cartItem => 
            cartItem.id === item.id 
              ? { ...cartItem, quantity: newQuantity } 
              : cartItem
          )
        };
      });
    } catch (err) {
      console.error('Error updating quantity:', err);
      setError('Не удалось обновить количество. Пожалуйста, попробуйте позже.');
    }
  };

  const handleRemoveItem = async (item: CartItem) => {
    try {
      await CartService.removeFromCart(user!.id, item.game.id);
      // Update local state for immediate UI update
      setCart(prevCart => {
        if (!prevCart) return null;
        
        return {
          ...prevCart,
          cartItems: prevCart.cartItems.filter(cartItem => cartItem.id !== item.id)
        };
      });
    } catch (err) {
      console.error('Error removing item:', err);
      setError('Не удалось удалить товар. Пожалуйста, попробуйте позже.');
    }
  };

  const handleClearCart = async () => {
    try {
      await CartService.clearCart(user!.id);
      setCart(prevCart => prevCart ? { ...prevCart, cartItems: [] } : null);
    } catch (err) {
      console.error('Error clearing cart:', err);
      setError('Не удалось очистить корзину. Пожалуйста, попробуйте позже.');
    }
  };

  const handleCreateOrder = async () => {
    if (isCartEmpty) {
      return;
    }
    
    try {
      setOrderLoading(true);
      setOrderError(null);
      
      await OrderService.createOrder(user!.id);
      
      // Clear cart after successful order
      setCart(prevCart => prevCart ? { ...prevCart, cartItems: [] } : null);
      setOrderSuccess(true);
    } catch (err) {
      console.error('Error creating order:', err);
      setOrderError('Не удалось создать заказ. Пожалуйста, попробуйте позже.');
    } finally {
      setOrderLoading(false);
    }
  };

  if (loading) {
    return <div className="loading">Загрузка...</div>;
  }

  if (error) {
    return <div className="error-message">{error}</div>;
  }

  if (orderSuccess) {
    return (
      <div className={styles.cartPage}>
        <div className="container">
          <div className={styles.orderSuccess}>
            <h2>Заказ успешно оформлен!</h2>
            <p>Спасибо за покупку! Вы можете просмотреть детали заказа в своем профиле.</p>
            <div className={styles.orderSuccessActions}>
              <Link to="/profile" className="btn btn-primary">Перейти в профиль</Link>
              <Link to="/games" className="btn btn-outline">Продолжить покупки</Link>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className={styles.cartPage}>
      <div className="container">
        <div className={styles.cartHeader}>
          <h1>Корзина</h1>
          <Link to="/games" className={styles.backToCatalog}>
            <FaArrowLeft /> Вернуться к каталогу
          </Link>
        </div>

        {!cart?.cartItems?.length ? (
          <div className={styles.emptyCart}>
            <h2>Ваша корзина пуста</h2>
            <p>Добавьте товары из каталога, чтобы оформить заказ.</p>
            <Link to="/games" className="btn btn-primary">Перейти в каталог</Link>
          </div>
        ) : (
          <div className={styles.cartContent}>
            <div className={styles.cartItems}>
              <div className={styles.cartItemsHeader}>
                <div>Товар</div>
                <div>Цена</div>
                <div>Количество</div>
                <div>Сумма</div>
                <div>Действия</div>
              </div>
              
              {cart.cartItems.map(item => (
                <div key={item.id} className={styles.cartItem}>
                  <div className={styles.cartItemProduct}>
                    <div className={styles.cartItemImage}>
                      <img src={`data:image/png;base64,${item.game.image_data}`} alt={item.game.title} />
                    </div>
                    <div className={styles.cartItemInfo}>
                      <Link to={`/games/${item.game.id}`} className={styles.cartItemTitle}>
                        {item.game.title}
                      </Link>
                      <div className={styles.cartItemMeta}>
                        <span>{item.game.category.name}</span>
                        <span>{item.game.developer.name}</span>
                      </div>
                    </div>
                  </div>
                  
                  <div className={styles.cartItemPrice}>
                    {item.game.price.toFixed(2)}₽
                  </div>
                  
                  <div>
                    <div className={styles.quantityControl}>
                      <button 
                        className={styles.quantityBtn}
                        onClick={() => handleQuantityChange(item, item.quantity - 1)}
                        disabled={item.quantity <= 1}
                      >
                        -
                      </button>
                      <span className={styles.quantityValue}>{item.quantity}</span>
                      <button 
                        className={styles.quantityBtn}
                        onClick={() => handleQuantityChange(item, item.quantity + 1)}
                      >
                        +
                      </button>
                    </div>
                  </div>
                  
                  <div className={styles.cartItemTotal}>
                    {(item.game.price * item.quantity).toFixed(2)}₽
                  </div>
                  
                  <div className={styles.cartItemActions}>
                    <button 
                      className={styles.removeItemBtn}
                      onClick={() => handleRemoveItem(item)}
                      title="Удалить"
                    >
                      <FaTrash />
                    </button>
                  </div>
                </div>
              ))}
            </div>
            
            <div className={styles.cartSummary}>
              <h2>Информация о заказе</h2>
              
              <div className={styles.summaryRow}>
                <span>Количество товаров:</span>
                <span>{totalItems}</span>
              </div>
              
              <div className={styles.summaryRow}>
                <span>Промежуточный итог:</span>
                <span>{subtotal} ₽</span>
              </div>
              
              <div className={`${styles.summaryRow} ${styles.total}`}>
                <span>Итого:</span>
                <span>{subtotal} ₽</span>
              </div>
              
              <div className={styles.summaryActions}>
                <button 
                  type="button" 
                  className={styles.checkoutBtn}
                  onClick={handleCreateOrder}
                  disabled={isCartEmpty || orderLoading}
                >
                  {orderLoading ? 'Оформление...' : 'Оформить заказ'} <FaCreditCard />
                </button>
                
                <button 
                  type="button" 
                  className={styles.clearCartBtn}
                  onClick={handleClearCart}
                  disabled={isCartEmpty}
                >
                  Очистить корзину <FaTrash />
                </button>
              </div>
              
              {orderError && (
                <div className="error-message">{orderError}</div>
              )}
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default CartPage; 