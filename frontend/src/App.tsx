import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import './App.scss';
import Home from './pages/home/Home';
import Login from './pages/auth/Login';
import Register from './pages/auth/Register';
import ProfileView from './pages/profile/ProfileView';
import ProfileEdit from './pages/profile/ProfileEdit';
import PublicationView from './pages/publication/PublicationView';

function App() {
  return (
    <Router>
      <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
      <Route path="/profile" element={<ProfileView />} />
      <Route path="/profile/edit" element={<ProfileEdit />} />
      <Route path="/publication" element={<PublicationView />} />
      </Routes>
    </Router>
  );
}

export default App;
