from flask import Flask
import os

app = Flask(__name__)

@app.route("/health")
def health():
    return "Product service is healthy!", 200

@app.route("/product")
def product():
    return {"message": "Hello from product-service"}, 200

if __name__ == "__main__":
    port = int(os.getenv("PORT", 8002))
    app.run(host="0.0.0.0", port=port)
