from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/api/v1/sensitive', methods=['GET'])
def sensitive_data():
    # Return some dummy sensitive text
    return jsonify({"message": "This is some sensitive data!"})

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
