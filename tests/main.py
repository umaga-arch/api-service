import os
from flask import Flask, jsonify
from flask_sqlalchemy import SQLAlchemy

def create_app():
    app = Flask(__name__)
    app.config['SQLALCHEMY_DATABASE_URI'] = os.environ.get('DATABASE_URL')
    app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    db = SQLAlchemy(app)

    from api_service.models import User, Post
    db.create_all()

    @app.route('/')
    def index():
        return jsonify({'message': 'Hello, World!'})

    @app.route('/users')
    def get_users():
        users = User.query.all()
        return jsonify([{'id': user.id, 'name': user.name} for user in users])

    @app.route('/users/<int:user_id>')
    def get_user(user_id):
        user = User.query.get(user_id)
        if user:
            return jsonify({'id': user.id, 'name': user.name})
        return jsonify({'error': 'User not found'}), 404

    return app

if __name__ == '__main__':
    app = create_app()
    app.run(debug=True)