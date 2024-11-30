# Server binary name
SERVER_BIN=coffee-shop-server

# Client binary name
CLIENT_BIN=coffee-cli

# Build the server binary
build-server:
	@echo "Building the coffee shop server..."
	go build -o $(SERVER_BIN) ./main.go

# Build the client binary
build-client:
	@echo "Building the coffee CLI..."
	go build -o $(CLIENT_BIN) ./client/main.go

# Run the server
run-server: build-server
	@echo "Running the coffee shop server..."
	./$(SERVER_BIN)

# Run the client with the correct subcommand
run-client: build-client
	@echo "Running the coffee CLI..."
	@echo "Please use 'make view-orders' to view all orders, 'make place-order' to place an order, 'make get-order' to view a particular order"
	@echo "Example usage:"
	@echo "    make view-orders"
	@echo "    make place-order"
	@echo "    make get-order"
	@echo ""
	@echo "To get detailed help on each command, run the following:"
	@echo "    ./$(CLIENT_BIN) --help    # View help for the root command"
	@echo "    ./$(CLIENT_BIN) view --help    # View help for viewing orders"
	@echo "    ./$(CLIENT_BIN) place --help    # View help for placing orders"
	@echo "    ./$(CLIENT_BIN) get --help    # View help for getting an order by ID"

# Clean all build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(SERVER_BIN) $(CLIENT_BIN)

# Install dependencies
install-deps:
	@echo "Installing dependencies..."
	go mod tidy

# View all orders (client functionality)
view-orders:
	@echo "Viewing all orders..."
	./$(CLIENT_BIN) view

# Place an order (client functionality)
place-order:
	@echo "Please enter the following details to place a new order:"
	@read -p "Client Name: " client; \
	read -p "Coffee Type: " coffee; \
	read -p "Size: " size; \
	read -p "Quantity: " quantity; \
	read -p "Comment (optional): " comment; \
	./$(CLIENT_BIN) place "$$client" "$$coffee" "$$size" "$$quantity" "$$comment"

# Get details of a specific order by ID (client functionality)
get-order:
	@echo "Please enter the order ID to view details:"
	@read -p "Order ID: " orderID; \
	./$(CLIENT_BIN) get $$orderID

# Build everything (server + client)
all: build-server build-client
	@echo "Build complete."

# Help section
help:
	@echo "Makefile Commands:"
	@echo "  make build-server    - Builds the server binary (coffee-shop-server)"
	@echo "  make build-client    - Builds the client binary (coffee-cli)"
	@echo "  make run-server      - Builds and Runs the coffee shop server directly"
	@echo "  make run-client      - Builds and Runs the coffee CLI directly"
	@echo "  make clean           - Cleans up build artifacts"
	@echo "  make install-deps    - Installs the Go dependencies"
	@echo "  make view-orders     - (Generate coffee-cli binary first) View all orders"
	@echo "  make place-order     - (Generate coffee-cli binary first) Places a coffee order (Enter details at runtime)"
	@echo "  make get-order       - (Generate coffee-cli binary first) Get details of a specific order by ID (Enter details at runtime)"
	@echo "  make help            - Displays this help message"
	@echo ""
	@echo "(OPTIONAL)To view detailed help for a command, use '--help' with the command:"
	@echo "  Example: ./coffee-cli view --help    # View help for viewing all orders"
	@echo "  Example: ./coffee-cli place --help   # View help for placing an order"
	@echo "  Example: ./coffee-cli get --help     # View help for getting an order by ID"
