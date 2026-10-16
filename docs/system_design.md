# URL Shortener like Bit.ly 

www.mylongurl.com/long/is/annoying -> www.bit.ly/GHJ23 -> redirected to the first url

## Process 

> Functional Requirements 
> core features

> Non-Functional Requirements
> service quality related things

1. Requirements 
> Functional Requirements 
> Non-Functional Requirements
    
2. Core Entities 
> High level list of that gets persisted 
> Names of tables and collections in data model

3. API or Interface 
> User facing api 
> Contract with the users and backend

4. Data Flow 
> Skip 
> Rate Limitter, Message Queues

5. High-Level Design 
> Drawing stuff -> core stuff

6. Deep Dives
> Non-Functional evolve the high-level


### Requirements 

- Ask Questions

- Functional Requirements 
1. Users can create a short url from long url 
2. Users can be redirected to original from shorturl
3. Users can support custom alias
4. Expiration time

- Non-Functional Requirements 
- Low Latency on redirections (200ms) (Humans perceive real-time)
- Scalability (100 million daily) (1B URLs)
- Ensure Uniqueness of shortcodes 
- CAP Theorem (Consistency, Availability, Partition Tolerance) 
- Read-After Write consistency (Banking, Ticketing)
- High Availability, Eventual Consistency for url shortening - i disagree
- back of the envelope - scale, latency, storage


### Core Entities 

- Originla URL 
- Short URL 
- User

### API 
almost a one to one mapping of functional requirements 

// shorten a url 
POST /urls -> shortURL 
{
    Original URL
    alias
    expiration time
}

// redirection
GET {shortURL} -> Redirect to OriginalURL 
Send code 300X code Permanent or Temporary
Send a 301 redirect: Permanent: No analytics; but compute costs
Send a 302 redirect: Temporary

### High Level Design 

URL Table
- shorturl/customAlias
- longurl
- expirationTime
- creationTime
- userId

8 + 100 + 8 + 100 + 8 = 500bytes * 1B = 500GB

User
- userId
- ... or Name, email


- ShortCode Creation 
- fast
- unique 
- short
- 5 to 7 characters
- random number generator
    - base62 encoding 0-9, A-Z, a-z
    - Birthday Paradox
    - we need to check for collisions for check 
- Hash the long url 
    - md5(longURL) -> hash -> base62(slice[:6])
    - same as above 
- Avoid checking the database
    - incrementing a counter -> base62 
    - predictability bad for security -> they can see how many urls are there
    - warning don't shorten private url 
    - rate limiter
    - bijective function - squids.org

- low latency on redirects 
- indexing
- redis/memcached
    - readthrough - cache miss - read db
    - LRU least recently used
    - cache in a cdn 

- Back of the envelope math 
- 10^8 / 10^5 = 10^3 = 1000 rps * (10k - 100k rps) when scaling 
- scale horizontally 
- more reads vs writes (system is ready heavy)

- read service vs write service (microservices)
- more read services and less write services (overkill)
- we need a global counter (Redis INCR)
- or take a slice of counts and use them from memory 

- to increase readthrouput in database do sharding 
- redis high availability mode 
- database replicas

