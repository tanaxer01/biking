
# Basic setup
I added `created_at` and `updated_at` for all entities, its adds important info that could prove
usefull in debugging wierd cases in a real product. Also could make some of the fields in `Rental`
not needed but I'll leave them for now. (Using `goose` for migrations)

I'm not sure if using `Chi` will bring something new to the table, so for now I wont use it (may change)

For the project layout, I'm not sure what would be best for this case. Ideally having the business 
logic & infra separated will make testing a looot easier.

After looking around diferent blogs and presentations I liked this post by Aviv Carmi It goes through
some diferent layouts and arrives at something that I fill could be easily maintained.

https://avivcarmi.com/finding-the-best-go-project-structure-part-2/

# Draft of business logic impl
1. This way I can implement `core/user` without even touching the `infra` side of things. 

2. Admin logic takes its own service maily because those endpoints wont need to look at the jwt token
and also that way we can see how a bigger service could look like. 

