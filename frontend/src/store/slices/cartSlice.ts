import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface CartItem {
  id: string;
  productId: string;
  name: string;
  price: number;
  size: string;
  quantity: number;
  imageUrl: string;
}

interface CartState {
  items: CartItem[];
  loading: boolean;
  error: string | null;
}

const initialState: CartState = {
  items: [],
  loading: false,
  error: null,
};

const cartSlice = createSlice({
  name: 'cart',
  initialState,
  reducers: {
    addToCartStart: (state) => {
      state.loading = true;
      state.error = null;
    },
    addToCartSuccess: (state, action: PayloadAction<CartItem>) => {
      state.loading = false;
      const existingItem = state.items.find(
        item => item.productId === action.payload.productId && item.size === action.payload.size
      );
      if (existingItem) {
        existingItem.quantity += action.payload.quantity;
      } else {
        state.items.push(action.payload);
      }
    },
    addToCartFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },
    removeFromCart: (state, action: PayloadAction<string>) => {
      state.items = state.items.filter(item => item.id !== action.payload);
    },
    updateQuantity: (state, action: PayloadAction<{ id: string; quantity: number }>) => {
      const item = state.items.find(item => item.id === action.payload.id);
      if (item) {
        item.quantity = action.payload.quantity;
      }
    },
    clearCart: (state) => {
      state.items = [];
    },
    setCartItems: (state, action: PayloadAction<CartItem[]>) => {
      state.items = action.payload;
    },
  },
});

export const {
  addToCartStart,
  addToCartSuccess,
  addToCartFailure,
  removeFromCart,
  updateQuantity,
  clearCart,
  setCartItems,
} = cartSlice.actions;

export default cartSlice.reducer; 