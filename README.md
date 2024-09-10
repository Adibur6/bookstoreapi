
# bookapi Kubernetes Deployment

This README provides basic commands to deploy the `bookapi` application on Kubernetes and access it using port forwarding.

## Commands

1.  **Deploy the Application**
    
    Apply the Deployment configuration:
    
 
    `kubectl apply -f deployment.yaml` 
    
2.  **Expose the Application**
    
    Apply the Service configuration:
    
    `kubectl apply -f service.yaml` 
    
3.  **Access the Application**
    
    Forward the Service port to your local machine:
    
    `kubectl port-forward svc/bookapi-service 8080:80` 
    
    You can now access the application at `http://localhost:8080` in your browser or via Postman.
