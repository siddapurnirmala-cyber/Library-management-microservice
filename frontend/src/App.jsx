import React, { useState, useEffect } from 'react';
import './App.css';
import { Book, Users, Repeat, LogOut, ChevronRight, Plus, Search, Loader2 } from 'lucide-react';
import { api } from './services/api';

function App() {
  const [activeTab, setActiveTab] = useState('dashboard');
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [books, setBooks] = useState([]);
  const [members, setMembers] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // For demo/prototype, simulate auth based on localStorage
    const authStatus = localStorage.getItem('lib_auth');
    if (authStatus === 'true') setIsAuthenticated(true);

    fetchData();
  }, []);

  const fetchData = async () => {
    setLoading(true);
    try {
      const booksData = await api.getBooks();
      const membersData = await api.getMembers();
      setBooks(booksData.books || []);
      setMembers(membersData.members || []);
    } catch (error) {
      console.error("Failed to fetch data:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleLogin = () => {
    window.location.href = 'http://localhost:8082/auth/google/login';
  };

  const handleLogout = () => {
    localStorage.removeItem('lib_auth');
    setIsAuthenticated(false);
  };

  if (!isAuthenticated) {
    return (
      <div className="login-container">
        <div className="glass card fade-in login-card">
          <h1>LibFlow</h1>
          <p>The modern library management system.</p>
          <button className="btn-primary" onClick={handleLogin}>
            Sign in with Google
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="app-layout">
      <nav className="glass sidebar">
        <div className="logo">
          <Book size={32} color="#6366f1" />
          <span>LibFlow</span>
        </div>
        <div className="nav-items">
          <button
            className={`nav-item ${activeTab === 'dashboard' ? 'active' : ''}`}
            onClick={() => setActiveTab('dashboard')}
          >
            <Repeat size={20} /> Dashboard
          </button>
          <button
            className={`nav-item ${activeTab === 'books' ? 'active' : ''}`}
            onClick={() => setActiveTab('books')}
          >
            <Book size={20} /> Books
          </button>
          <button
            className={`nav-item ${activeTab === 'members' ? 'active' : ''}`}
            onClick={() => setActiveTab('members')}
          >
            <Users size={20} /> Members
          </button>
        </div>
        <div className="nav-footer">
          <button className="btn-logout" onClick={handleLogout}>
            <LogOut size={20} /> Logout
          </button>
        </div>
      </nav>

      <main className="content">
        <header className="glass top-nav">
          <div className="title-section">
            <h2 className="title-text">{activeTab.charAt(0).toUpperCase() + activeTab.slice(1)}</h2>
          </div>
          <div className="user-profile">
            <div className="avatar">AD</div>
          </div>
        </header>

        <section className="view-content fade-in">
          {loading ? (
            <div className="loading-state">
              <Loader2 className="spin" size={48} color="#6366f1" />
              <p>Fetching library statistics...</p>
            </div>
          ) : (
            <>
              {activeTab === 'dashboard' && <DashboardView books={books} members={members} />}
              {activeTab === 'books' && <BooksView books={books} refresh={fetchData} />}
              {activeTab === 'members' && <MembersView members={members} refresh={fetchData} />}
            </>
          )}
        </section>
      </main>
    </div>
  );
}

const DashboardView = ({ books, members }) => {
  const availableBooks = books.reduce((acc, book) => acc + book.available_copies, 0);
  const totalCopies = books.reduce((acc, book) => acc + book.total_copies, 0);

  return (
    <div className="dashboard-grid">
      <div className="glass stat-card">
        <div className="stat-icon"><Book color="#06b6d4" /></div>
        <div className="stat-info">
          <h3>{books.length}</h3>
          <p>Unique Titles</p>
        </div>
      </div>
      <div className="glass stat-card">
        <div className="stat-icon"><Users color="#10b981" /></div>
        <div className="stat-info">
          <h3>{members.length}</h3>
          <p>Total Members</p>
        </div>
      </div>
      <div className="glass stat-card">
        <div className="stat-icon"><Repeat color="#6366f1" /></div>
        <div className="stat-info">
          <h3>{totalCopies - availableBooks}</h3>
          <p>Active Borrows</p>
        </div>
      </div>
    </div>
  );
};

const BooksView = ({ books, refresh }) => (
  <div className="data-view">
    <div className="actions-bar">
      <div className="search-box">
        <Search size={18} />
        <input type="text" placeholder="Search books..." />
      </div>
      <button className="btn-primary">
        <Plus size={18} /> Add Book
      </button>
    </div>
    <div className="glass table-container">
      <table>
        <thead>
          <tr>
            <th>Title</th>
            <th>Author</th>
            <th>Year</th>
            <th>Availability</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {books.map(book => (
            <tr key={book.id}>
              <td>{book.title}</td>
              <td>{book.author}</td>
              <td>{book.published_year}</td>
              <td>
                <div className="progress-bar-bg">
                  <div
                    className="progress-bar-fill"
                    style={{ width: `${(book.available_copies / book.total_copies) * 100}%` }}
                  ></div>
                </div>
                <span className="availability-text">{book.available_copies}/{book.total_copies}</span>
              </td>
              <td><button className="btn-action">Edit</button></td>
            </tr>
          ))}
          {books.length === 0 && (
            <tr>
              <td colSpan="5" className="empty-row">No books found in the library.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  </div>
);

const MembersView = ({ members, refresh }) => (
  <div className="data-view">
    <div className="actions-bar">
      <div className="search-box">
        <Search size={18} />
        <input type="text" placeholder="Search members..." />
      </div>
      <button className="btn-primary">
        <Plus size={18} /> Add Member
      </button>
    </div>
    <div className="glass table-container">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Joined At</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {members.map(member => (
            <tr key={member.id}>
              <td>{member.name}</td>
              <td>{member.email}</td>
              <td>{new Date(member.joined_at).toLocaleDateString()}</td>
              <td><button className="btn-action">Edit</button></td>
            </tr>
          ))}
          {members.length === 0 && (
            <tr>
              <td colSpan="4" className="empty-row">No members registered yet.</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  </div>
);

export default App;
