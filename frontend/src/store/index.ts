import { configureStore } from '@reduxjs/toolkit';
import authReducer from './slices/authSlice';
import cartReducer from './slices/cartSlice';
import tryOnReducer from './slices/tryOnSlice';

export const store = configureStore({
  reducer: {
    auth: authReducer,
    cart: cartReducer,
    tryOn: tryOnReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch; 