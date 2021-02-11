from locust import HttpUser, task, between, events
from essential_generators import DocumentGenerator
from faker import Faker
import random
import json
import urllib.parse
import os
import logging

logging.basicConfig(format='%(asctime)s %(levelname)s:%(message)s', level=logging.INFO)

product_list = []
product_id = []

api_key = os.environ["API_KEY"]
api_secret = os.environ["API_SECRET"]

wait_time_lrange = os.getenv('WAIT_TIME_LRANGE')
if wait_time_lrange is None:
  wait_time_lrange = int(120)
else:
  wait_time_lrange = int(wait_time_lrange)

wait_time_rrange = os.getenv('WAIT_TIME_RRANGE')
if wait_time_rrange is None:
  wait_time_rrange = int(300)
else:
  wait_time_rrange = int(wait_time_rrange)

doc_gen = DocumentGenerator()

template = {
  'comment_post_ID': product_id, 
  'rating': [1, 2, 3, 4, 5],
  'comment': 'sentence',
  'author': 'name',
  'email': 'email',
  'submit': 'Submit',
  'comment_parent': 0
}

doc_gen.set_template(template)

fake = Faker()

def gen_addr(first_name, last_name, email):
  ret = {
          'first_name': first_name,
          'last_name': last_name,
          'company': '',
          'address_1': fake.street_address(),
          'address_2': '',
          'city': fake.city(),
          'state': fake.state(),
          'postcode': fake.zipcode(),
          'country': fake.country_code(),
          'email': email,
          'phone': fake.phone_number()
        }
  return ret

def gen_user_details():

  first_name = fake.first_name()
  last_name = fake.last_name()
  email = fake.email()
  addr = gen_addr(first_name, last_name, email)

  ret = {
    'email': email,
    'first_name': first_name,
    'last_name': last_name,
    'username': first_name + '.' + last_name,
    'billing': addr,
    'shipping': addr
  }
  logging.debug("Generating user details: %s", ret)
  return ret

#
# Sample: curl 'http://woodkraft.ntnxsherlock.com/wp-json/wc/store/products?catalog_visibility=catalog&per_page=100' \
#                -H 'Accept: application/json, */*;q=0.1'
#
def setup(l):
    if len(product_list) != 0:
      logging.info("Using Product List: %s", product_list)
      return

    l.client.headers['Content-Type'] = "application/json"
    response = l.client.get("/wp-json/wc/store/products?catalog_visibility=catalog&per_page=100").json()
    logging.debug("Product response: %s", response)
    for items in response:
      logging.debug("permalink: %s", items['permalink'])
      name = items['permalink'].split('/')[-2]
      product_list.append(name)
      product_id.append(items['id'])
    logging.info("Creating Product List: %s", product_list)
    logging.info("Product Id: %s", product_id)


#
# Sample: curl -X POST https://example.com/wp-json/wc/v3/customers \
#         -u consumer_key:consumer_secret -H "Content-Type: application/json" -d '{}'
#
def create_customer(l):
    logging.debug("Adding new customer: %s", l.user_details)
    logging.info("Adding new customer: %s, %s", l.user_details['email'], l.user_details['username'])
    customer_details = json.dumps(l.user_details)
    l.client.headers['Content-Type'] = "application/json"
    response = l.client.post("/wp-json/wc/v3/customers", data=customer_details, auth=(api_key, api_secret))
    logging.debug("create_customer response: %s", response.json())
    logging.info("create_customer response: %s %s", response.status_code, response.reason)
  
  
def place_order(l):
    l.client.headers['Content-Type'] = "application/json"
    order_details = {
      'payment_method': 'bac',
      'payment_method_title': 'Direct Bank Transfer',
      'set_paid': 'true',
      'billing': l.user_details['billing'],
      'shipping': l.user_details['shipping'],
      'line_items': [{
        'product_id': random.choice(product_id),
        'quantity': 1
      }],
      'shipping_lines': [{
        'method_id': 'flat_rate',
        'method_title': 'Flat Rate',
        'total': '3.99'
      }]
    }
    order_details_json = json.dumps(order_details)
    logging.debug("place order: %s", order_details_json)
    response = l.client.post("/wp-json/wc/v3/orders", data=order_details_json, auth=(api_key, api_secret))
    response_json = response.json()
    id = response_json['id']
    logging.debug("place_order response: ", response_json)
    logging.info("place_order response: %s %s", response.status_code, response.reason)
    update_order(l, str(id))


def update_order(l, id):
    update_details = {
      'status': 'completed'
    }
    update_details_json = json.dumps(update_details)
    response = l.client.put("/wp-json/wc/v3/orders/" + id, data=update_details_json, auth=(api_key, api_secret))
    logging.debug("place_order response: ", response.json())
    logging.info("place_order response: %s %s", response.status_code, response.reason)
    

#
#Sample: curl 'http://woodkraft.ntnxsherlock.com/wp-comments-post.php' -H 'Content-Type: application/x-www-form-urlencoded' \
#       --data 'rating=5&comment=hard+to+solve&author=pd&email=pd%40pd.com&submit=Submit&comment_post_ID=72&comment_parent=0' \
#
def post_review_guest(l):
  review = json.dumps(doc_gen.documents(1)[0])
  logging.debug("Adding review: %s", review)
  response = l.client.post("/wp-comments-post.php", data=review)
  logging.debug("post_review_guest response: %s", response)
  logging.info("post_review_guest response: %s %s", response.status_code, response.reason)


def post_review(l):
  review = {
    'product_id': random.choice(product_id),
    'review': doc_gen.sentence(),
    'reviewer': l.user_details['first_name'] + " " + l.user_details['last_name'],
    'reviewer_email': l.user_details['email'],
    'rating': random.choice([1,2,3,4,5])
  }
  logging.debug("Adding review: %s", review)
  review_json = json.dumps(review)
  response = l.client.post("/wp-json/wc/v3/products/reviews", data=review_json, auth=(api_key, api_secret))
  logging.debug("post_review response: %s", response.json())
  logging.info("post_review response: %s %s", response.status_code, response.reason)


def browse_products(l):
  response = l.client.get("/wp-json/wc/v3/products/" + str(random.choice(product_id)), auth=(api_key, api_secret))
  logging.debug("browse_products response: %s", response.json())
  logging.info("browse_products response: %s %s", response.status_code, response.reason)


#curl 'http://woodkraft.ntnxsherlock.com/?product=interlocking-puzzle'
#//woodkraft.ntnxsherlock.com/product/vneck-tee/
def browse_products_guest(l):
  response = l.client.get("/products/" + random.choice(product_list))
  logging.debug("browse_products_guest response: %s", response)
  logging.info("browse_products_guest response: %s %s", response.status_code, response.reason)


class User(HttpUser):
    wait_time = between(wait_time_lrange, wait_time_rrange)
    tasks = {browse_products: 16, post_review: 8, place_order: 4}

    def on_start(self):
      self.user_details = gen_user_details()
      self.client.verify = False
      setup(self)
      create_customer(self)
