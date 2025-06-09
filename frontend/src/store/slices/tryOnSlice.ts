import { createSlice, PayloadAction } from '@reduxjs/toolkit';

interface TryOnImage {
  id: string;
  originalImage: string;
  resultImage: string | null;
  status: 'pending' | 'processing' | 'completed' | 'failed';
  productId: string | null;
  createdAt: string;
}

interface TryOnState {
  currentTryOn: TryOnImage | null;
  history: TryOnImage[];
  loading: boolean;
  error: string | null;
}

const initialState: TryOnState = {
  currentTryOn: null,
  history: [],
  loading: false,
  error: null,
};

const tryOnSlice = createSlice({
  name: 'tryOn',
  initialState,
  reducers: {
    uploadStart: (state) => {
      state.loading = true;
      state.error = null;
    },
    uploadSuccess: (state, action: PayloadAction<TryOnImage>) => {
      state.loading = false;
      state.currentTryOn = action.payload;
      state.history.unshift(action.payload);
    },
    uploadFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
    },
    processStart: (state) => {
      if (state.currentTryOn) {
        state.currentTryOn.status = 'processing';
      }
      state.loading = true;
      state.error = null;
    },
    processSuccess: (state, action: PayloadAction<{ id: string; resultImage: string }>) => {
      state.loading = false;
      if (state.currentTryOn && state.currentTryOn.id === action.payload.id) {
        state.currentTryOn.status = 'completed';
        state.currentTryOn.resultImage = action.payload.resultImage;
      }
      const historyItem = state.history.find(item => item.id === action.payload.id);
      if (historyItem) {
        historyItem.status = 'completed';
        historyItem.resultImage = action.payload.resultImage;
      }
    },
    processFailure: (state, action: PayloadAction<string>) => {
      state.loading = false;
      state.error = action.payload;
      if (state.currentTryOn) {
        state.currentTryOn.status = 'failed';
      }
    },
    setHistory: (state, action: PayloadAction<TryOnImage[]>) => {
      state.history = action.payload;
    },
    clearCurrentTryOn: (state) => {
      state.currentTryOn = null;
    },
  },
});

export const {
  uploadStart,
  uploadSuccess,
  uploadFailure,
  processStart,
  processSuccess,
  processFailure,
  setHistory,
  clearCurrentTryOn,
} = tryOnSlice.actions;

export default tryOnSlice.reducer; 