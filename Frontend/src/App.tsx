import './App.scss';

// Components
import Header from './components/Header';
import Footer from './components/Footer';
import AppRoutes from './components/AppRoutes';

// Context
import { AuthProvider } from './context/AuthContext';

function App() {
  return (
    <AuthProvider>
      <div className="App">
        <Header />
        <main className="main-content">
          <AppRoutes />
        </main>
        <Footer />
      </div>
    </AuthProvider>
  );
}

export default App; 