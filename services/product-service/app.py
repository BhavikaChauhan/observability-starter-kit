from flask import Flask
from telemetry import init_telemetry, tracer

app = Flask(__name__)
init_telemetry()
tr = tracer()

@app.route('/health')
def health():
    with tr.start_as_current_span('health-span'):
        return 'Product Service is healthy!'

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000)