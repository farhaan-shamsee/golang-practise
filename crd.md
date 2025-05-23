# CRD, RBAC, COntroller and OPerators

## CRD: Custom Resource Definition

- CRD is a way to extend Kubernetes API.
- CR is a custom resource
- CRD is a way to define a custom resource. CRD validates the CR.
- A CRD lets you define a new Kubernetes resource type — like Keycloak, KeycloakRealm, or KeycloakUser — just like native resources (e.g., Pods, Services).
- It acts as the schema for custom resources.

## CR (Custom Resource)

- A CR is an instance of the resource defined by the CRD — like a Pod is an instance of the Pod kind.
- It describes what you want, not how it’s achieved.
- This is a CR that creates an actual Keycloak deployment:

    ```yaml
    apiVersion: keycloak.org/v1alpha1
    kind: Keycloak
    metadata:
        name: example-keycloak
    spec:
        instances: 1
    ```

- This tells Kubernetes: "I want a Keycloak deployment with 1 instance."

## Controller

- A controller is a loop running in the background that watches the CRs (e.g., Keycloak) and takes action to make the actual state match the desired state.

- It implements reconciliation logic.
