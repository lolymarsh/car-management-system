import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import AppLayout from './components/AppLayout';
import CarList from './pages/CarList';
import CarForm from './pages/CarForm';

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<AppLayout />}>
          <Route path="/cars" element={<CarList />} />
          <Route path="/cars/new" element={<CarForm />} />
          <Route path="/cars/:car_id/edit" element={<CarForm />} />
        </Route>
        <Route path="*" element={<Navigate to="/cars" replace />} />
      </Routes>
    </BrowserRouter>
  );
}
