import os
import psycopg2
from flask import Flask, jsonify, request

app = Flask(__name__)


def get_db_connection():
    conn = psycopg2.connect(host=os.environ['POSTGRE_URL'],
                            database=os.environ['POSTGRES_DB'],
                            user=os.environ['POSTGRES_USER'],
                            password=os.environ['POSTGRES_PASSWORD'])
    return conn

@app.route('/books', methods=['GET'])
def read_all_books():
    try:
        cur = get_db_connection().cursor()
        cur.execute('SELECT * FROM books;')
        books = cur.fetchall()
        cur.close()

        response = []
        for fila in books:
            book={
                'code': fila[0],
                'name': fila[1],
                'author': fila[2],
                'pages': fila[3],
                'review': fila[4],
                'publish_date': fila[5]
            }
            response.append(book)

        return jsonify({'books': response, 'messages': 'books find'})
    except Exception as ex:
        return jsonify({'error': f'{ex}'})



@app.route('/books/<code>', methods=['GET'])
def read_books(code):
    try:
        cur = get_db_connection().cursor()
        cur.execute("SELECT * FROM books WHERE id = '{0}';".format(code))
        book = cur.fetchone()
        cur.close()

        print(book)
        if book == None:
            return jsonify({'error': 'INVALID_CODE_BOOK', 'message': f'The book with code {code} not found'})

        response = {
            'code': book[0],
            'name': book[1],
            'author': book[2],
            'pages': book[3],
            'review': book[4],
            'publish_date': book[5]
        }
        return jsonify(response)
    except Exception as ex:
        return jsonify({'error': f'{ex}'})        

@app.route('/books', methods=['POST'])
def save_book():
    try:
        cur = get_db_connection().cursor()
        sql=f"""INSERT INTO books (id, name, author, pages, review, publish_date) 
                VALUES('{request.json['code']}', '{request.json['name']}', '{request.json['author']}', 
                {request.json['pages']}, '{request.json['review']}', '{request.json['publish_date']}')"""
        print(sql)
        cur.execute(sql)
        cur.connection.commit()
        cur.close()

        return jsonify({'message': f'The book \'{request.json["name"]}\' is save.'})    
    except Exception as ex:
        print(ex)
        return jsonify({'error': f'{ex}'})

def error_page(error):
    return "<h1>Page Not found</h1>"

if __name__ == "__main__":
    app.register_error_handler(404, error_page)
    app.run(debug=False)