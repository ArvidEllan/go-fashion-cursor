import React, { useState, useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import {
  Container,
  Grid,
  Paper,
  Typography,
  Button,
  Box,
  CircularProgress,
  Alert,
} from '@mui/material';
import { CloudUpload as UploadIcon } from '@mui/icons-material';
import { RootState } from '../store';
import {
  uploadStart,
  uploadSuccess,
  uploadFailure,
  processStart,
  processSuccess,
  processFailure,
} from '../store/slices/tryOnSlice';
import { api } from '../services/api';

const TryOn: React.FC = () => {
  const dispatch = useDispatch();
  const { currentTryOn, loading, error } = useSelector((state: RootState) => state.tryOn);
  const [selectedFile, setSelectedFile] = useState<File | null>(null);
  const [previewUrl, setPreviewUrl] = useState<string | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);

  const handleFileSelect = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setSelectedFile(file);
      const reader = new FileReader();
      reader.onloadend = () => {
        setPreviewUrl(reader.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleUpload = async () => {
    if (!selectedFile) return;

    dispatch(uploadStart());
    const formData = new FormData();
    formData.append('photo', selectedFile);

    try {
      const response = await api.post('/api/try-on/upload', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      dispatch(uploadSuccess(response.data));
    } catch (err) {
      dispatch(uploadFailure('Failed to upload photo'));
    }
  };

  const handleProcess = async () => {
    if (!currentTryOn) return;

    dispatch(processStart());
    try {
      const response = await api.post('/api/try-on/process', {
        try_on_id: currentTryOn.id,
        product_id: 'selected-product-id', // This should come from product selection
      });

      dispatch(processSuccess({
        id: currentTryOn.id,
        resultImage: response.data.result_image,
      }));
    } catch (err) {
      dispatch(processFailure('Failed to process try-on'));
    }
  };

  return (
    <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h4" gutterBottom>
              Virtual Try-On
            </Typography>
            <Typography variant="body1" color="text.secondary" paragraph>
              Upload a photo of yourself to try on different clothes virtually.
            </Typography>
          </Paper>
        </Grid>

        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 3, height: '100%' }}>
            <Typography variant="h6" gutterBottom>
              Upload Photo
            </Typography>
            <Box
              sx={{
                border: '2px dashed #ccc',
                borderRadius: 2,
                p: 3,
                textAlign: 'center',
                mb: 2,
                minHeight: 200,
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'center',
              }}
            >
              {previewUrl ? (
                <img
                  src={previewUrl}
                  alt="Preview"
                  style={{ maxWidth: '100%', maxHeight: 300 }}
                />
              ) : (
                <>
                  <UploadIcon sx={{ fontSize: 48, color: 'text.secondary', mb: 2 }} />
                  <Typography color="text.secondary">
                    Drag and drop a photo here, or click to select
                  </Typography>
                </>
              )}
              <input
                type="file"
                accept="image/*"
                onChange={handleFileSelect}
                style={{ display: 'none' }}
                ref={fileInputRef}
              />
              <Button
                variant="contained"
                onClick={() => fileInputRef.current?.click()}
                sx={{ mt: 2 }}
              >
                Select Photo
              </Button>
            </Box>
            <Button
              variant="contained"
              color="primary"
              onClick={handleUpload}
              disabled={!selectedFile || loading}
              fullWidth
            >
              {loading ? <CircularProgress size={24} /> : 'Upload Photo'}
            </Button>
          </Paper>
        </Grid>

        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 3, height: '100%' }}>
            <Typography variant="h6" gutterBottom>
              Try-On Result
            </Typography>
            {error && (
              <Alert severity="error" sx={{ mb: 2 }}>
                {error}
              </Alert>
            )}
            {currentTryOn ? (
              <Box sx={{ textAlign: 'center' }}>
                {currentTryOn.resultImage ? (
                  <img
                    src={currentTryOn.resultImage}
                    alt="Try-on result"
                    style={{ maxWidth: '100%', maxHeight: 300 }}
                  />
                ) : (
                  <Box sx={{ p: 3 }}>
                    <Typography color="text.secondary" gutterBottom>
                      {currentTryOn.status === 'processing'
                        ? 'Processing your try-on...'
                        : 'Select a product to try on'}
                    </Typography>
                    {currentTryOn.status === 'processing' && (
                      <CircularProgress sx={{ mt: 2 }} />
                    )}
                  </Box>
                )}
                {currentTryOn.status === 'pending' && (
                  <Button
                    variant="contained"
                    color="primary"
                    onClick={handleProcess}
                    disabled={loading}
                    sx={{ mt: 2 }}
                  >
                    Try On Selected Product
                  </Button>
                )}
              </Box>
            ) : (
              <Box sx={{ p: 3, textAlign: 'center' }}>
                <Typography color="text.secondary">
                  Upload a photo to start your virtual try-on
                </Typography>
              </Box>
            )}
          </Paper>
        </Grid>
      </Grid>
    </Container>
  );
};

export default TryOn; 