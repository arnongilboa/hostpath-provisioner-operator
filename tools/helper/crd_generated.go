package helper

//hppCRD is a string yaml of the hpp CRD
var hppCRD string = 
`apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.9.2
  creationTimestamp: null
  name: hostpathprovisioners.hostpathprovisioner.kubevirt.io
spec:
  group: hostpathprovisioner.kubevirt.io
  names:
    kind: HostPathProvisioner
    listKind: HostPathProvisionerList
    plural: hostpathprovisioners
    singular: hostpathprovisioner
  scope: Cluster
  versions:
  - name: v1beta1
    schema:
      openAPIV3Schema:
        description: HostPathProvisioner is the Schema for the hostpathprovisioners
          API
        properties:
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: HostPathProvisionerSpec defines the desired state of HostPathProvisioner
            properties:
              featureGates:
                description: FeatureGates are a list of specific enabled feature gates
                items:
                  type: string
                type: array
                x-kubernetes-list-type: set
              imagePullPolicy:
                description: ImagePullPolicy is the container pull policy for the
                  host path provisioner containers
                type: string
              pathConfig:
                description: PathConfig describes the location and layout of PV storage
                  on nodes. Deprecated
                properties:
                  path:
                    description: Path The path the directories for the PVs are created
                      under
                    type: string
                  useNamingPrefix:
                    description: UseNamingPrefix Use the name of the PVC requesting
                      the PV as part of the directory created
                    type: boolean
                type: object
              storagePools:
                description: StoragePools are a list of storage pools
                items:
                  description: StoragePool defines how and where hostpath provisioner
                    can use storage to create volumes.
                  properties:
                    name:
                      description: Name specifies an identifier that is used in the
                        storage class arguments to identify the source to use.
                      type: string
                    path:
                      description: path the path to use on the host, this is a required
                        field
                      type: string
                    pvcTemplate:
                      description: PVCTemplate is the template of the PVC to create
                        as the source volume
                      properties:
                        accessModes:
                          description: 'accessModes contains the desired access modes
                            the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                          items:
                            type: string
                          type: array
                        dataSource:
                          description: 'dataSource field can be used to specify either:
                            * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)
                            * An existing PVC (PersistentVolumeClaim) If the provisioner
                            or an external controller can support the specified data
                            source, it will create a new volume based on the contents
                            of the specified data source. When the AnyVolumeDataSource
                            feature gate is enabled, dataSource contents will be copied
                            to dataSourceRef, and dataSourceRef contents will be copied
                            to dataSource when dataSourceRef.namespace is not specified.
                            If the namespace is specified, then dataSourceRef will
                            not be copied to dataSource.'
                          properties:
                            apiGroup:
                              description: APIGroup is the group for the resource
                                being referenced. If APIGroup is not specified, the
                                specified Kind must be in the core API group. For
                                any other third-party types, APIGroup is required.
                              type: string
                            kind:
                              description: Kind is the type of resource being referenced
                              type: string
                            name:
                              description: Name is the name of resource being referenced
                              type: string
                          required:
                          - kind
                          - name
                          type: object
                          x-kubernetes-map-type: atomic
                        dataSourceRef:
                          description: 'dataSourceRef specifies the object from which
                            to populate the volume with data, if a non-empty volume
                            is desired. This may be any object from a non-empty API
                            group (non core object) or a PersistentVolumeClaim object.
                            When this field is specified, volume binding will only
                            succeed if the type of the specified object matches some
                            installed volume populator or dynamic provisioner. This
                            field will replace the functionality of the dataSource
                            field and as such if both fields are non-empty, they must
                            have the same value. For backwards compatibility, when
                            namespace isn''t specified in dataSourceRef, both fields
                            (dataSource and dataSourceRef) will be set to the same
                            value automatically if one of them is empty and the other
                            is non-empty. When namespace is specified in dataSourceRef,
                            dataSource isn''t set to the same value and must be empty.
                            There are three important differences between dataSource
                            and dataSourceRef: * While dataSource only allows two
                            specific types of objects, dataSourceRef allows any non-core
                            object, as well as PersistentVolumeClaim objects. * While
                            dataSource ignores disallowed values (dropping them),
                            dataSourceRef preserves all values, and generates an error
                            if a disallowed value is specified. * While dataSource
                            only allows local objects, dataSourceRef allows objects
                            in any namespaces. (Beta) Using this field requires the
                            AnyVolumeDataSource feature gate to be enabled. (Alpha)
                            Using the namespace field of dataSourceRef requires the
                            CrossNamespaceVolumeDataSource feature gate to be enabled.'
                          properties:
                            apiGroup:
                              description: APIGroup is the group for the resource
                                being referenced. If APIGroup is not specified, the
                                specified Kind must be in the core API group. For
                                any other third-party types, APIGroup is required.
                              type: string
                            kind:
                              description: Kind is the type of resource being referenced
                              type: string
                            name:
                              description: Name is the name of resource being referenced
                              type: string
                            namespace:
                              description: Namespace is the namespace of resource
                                being referenced Note that when a namespace is specified,
                                a gateway.networking.k8s.io/ReferenceGrant object
                                is required in the referent namespace to allow that
                                namespace's owner to accept the reference. See the
                                ReferenceGrant documentation for details. (Alpha)
                                This field requires the CrossNamespaceVolumeDataSource
                                feature gate to be enabled.
                              type: string
                          required:
                          - kind
                          - name
                          type: object
                        resources:
                          description: 'resources represents the minimum resources
                            the volume should have. If RecoverVolumeExpansionFailure
                            feature is enabled users are allowed to specify resource
                            requirements that are lower than previous value but must
                            still be higher than capacity recorded in the status field
                            of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                          properties:
                            claims:
                              description: "Claims lists the names of resources, defined
                                in spec.resourceClaims, that are used by this container.
                                \n This is an alpha field and requires enabling the
                                DynamicResourceAllocation feature gate. \n This field
                                is immutable."
                              items:
                                description: ResourceClaim references one entry in
                                  PodSpec.ResourceClaims.
                                properties:
                                  name:
                                    description: Name must match the name of one entry
                                      in pod.spec.resourceClaims of the Pod where
                                      this field is used. It makes that resource available
                                      inside a container.
                                    type: string
                                required:
                                - name
                                type: object
                              type: array
                              x-kubernetes-list-map-keys:
                              - name
                              x-kubernetes-list-type: map
                            limits:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: 'Limits describes the maximum amount of
                                compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                              type: object
                            requests:
                              additionalProperties:
                                anyOf:
                                - type: integer
                                - type: string
                                pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                x-kubernetes-int-or-string: true
                              description: 'Requests describes the minimum amount
                                of compute resources required. If Requests is omitted
                                for a container, it defaults to Limits if that is
                                explicitly specified, otherwise to an implementation-defined
                                value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                              type: object
                          type: object
                        selector:
                          description: selector is a label query over volumes to consider
                            for binding.
                          properties:
                            matchExpressions:
                              description: matchExpressions is a list of label selector
                                requirements. The requirements are ANDed.
                              items:
                                description: A label selector requirement is a selector
                                  that contains values, a key, and an operator that
                                  relates the key and values.
                                properties:
                                  key:
                                    description: key is the label key that the selector
                                      applies to.
                                    type: string
                                  operator:
                                    description: operator represents a key's relationship
                                      to a set of values. Valid operators are In,
                                      NotIn, Exists and DoesNotExist.
                                    type: string
                                  values:
                                    description: values is an array of string values.
                                      If the operator is In or NotIn, the values array
                                      must be non-empty. If the operator is Exists
                                      or DoesNotExist, the values array must be empty.
                                      This array is replaced during a strategic merge
                                      patch.
                                    items:
                                      type: string
                                    type: array
                                required:
                                - key
                                - operator
                                type: object
                              type: array
                            matchLabels:
                              additionalProperties:
                                type: string
                              description: matchLabels is a map of {key,value} pairs.
                                A single {key,value} in the matchLabels map is equivalent
                                to an element of matchExpressions, whose key field
                                is "key", the operator is "In", and the values array
                                contains only "value". The requirements are ANDed.
                              type: object
                          type: object
                          x-kubernetes-map-type: atomic
                        storageClassName:
                          description: 'storageClassName is the name of the StorageClass
                            required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1'
                          type: string
                        volumeMode:
                          description: volumeMode defines what type of volume is required
                            by the claim. Value of Filesystem is implied when not
                            included in claim spec.
                          type: string
                        volumeName:
                          description: volumeName is the binding reference to the
                            PersistentVolume backing this claim.
                          type: string
                      type: object
                  required:
                  - name
                  - path
                  type: object
                type: array
                x-kubernetes-list-type: atomic
              workload:
                description: Restrict on which nodes HPP workload pods will be scheduled
                properties:
                  affinity:
                    description: affinity enables pod affinity/anti-affinity placement
                      expanding the types of constraints that can be expressed with
                      nodeSelector. affinity is going to be applied to the relevant
                      kind of pods in parallel with nodeSelector See https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity
                    properties:
                      nodeAffinity:
                        description: Describes node affinity scheduling rules for
                          the pod.
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node matches the corresponding matchExpressions;
                              the node(s) with the highest sum are the most preferred.
                            items:
                              description: An empty preferred scheduling term matches
                                all objects with implicit weight 0 (i.e. it's a no-op).
                                A null preferred scheduling term matches no objects
                                (i.e. is also a no-op).
                              properties:
                                preference:
                                  description: A node selector term, associated with
                                    the corresponding weight.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                weight:
                                  description: Weight associated with matching the
                                    corresponding nodeSelectorTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - preference
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to an update), the system
                              may or may not try to eventually evict the pod from
                              its node.
                            properties:
                              nodeSelectorTerms:
                                description: Required. A list of node selector terms.
                                  The terms are ORed.
                                items:
                                  description: A null or empty node selector term
                                    matches no objects. The requirements of them are
                                    ANDed. The TopologySelectorTerm type implements
                                    a subset of the NodeSelectorTerm.
                                  properties:
                                    matchExpressions:
                                      description: A list of node selector requirements
                                        by node's labels.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchFields:
                                      description: A list of node selector requirements
                                        by node's fields.
                                      items:
                                        description: A node selector requirement is
                                          a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: The label key that the selector
                                              applies to.
                                            type: string
                                          operator:
                                            description: Represents a key's relationship
                                              to a set of values. Valid operators
                                              are In, NotIn, Exists, DoesNotExist.
                                              Gt, and Lt.
                                            type: string
                                          values:
                                            description: An array of string values.
                                              If the operator is In or NotIn, the
                                              values array must be non-empty. If the
                                              operator is Exists or DoesNotExist,
                                              the values array must be empty. If the
                                              operator is Gt or Lt, the values array
                                              must have a single element, which will
                                              be interpreted as an integer. This array
                                              is replaced during a strategic merge
                                              patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                  type: object
                                  x-kubernetes-map-type: atomic
                                type: array
                            required:
                            - nodeSelectorTerms
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      podAffinity:
                        description: Describes pod affinity scheduling rules (e.g.
                          co-locate this pod in the same node, zone, etc. as some
                          other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the affinity expressions specified
                              by this field, but it may choose a node that violates
                              one or more of the expressions. The node that is most
                              preferred is the one with the greatest sum of weights,
                              i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the affinity requirements specified by
                              this field are not met at scheduling time, the pod will
                              not be scheduled onto the node. If the affinity requirements
                              specified by this field cease to be met at some point
                              during pod execution (e.g. due to a pod label update),
                              the system may or may not try to eventually evict the
                              pod from its node. When there are multiple elements,
                              the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                      podAntiAffinity:
                        description: Describes pod anti-affinity scheduling rules
                          (e.g. avoid putting this pod in the same node, zone, etc.
                          as some other pod(s)).
                        properties:
                          preferredDuringSchedulingIgnoredDuringExecution:
                            description: The scheduler will prefer to schedule pods
                              to nodes that satisfy the anti-affinity expressions
                              specified by this field, but it may choose a node that
                              violates one or more of the expressions. The node that
                              is most preferred is the one with the greatest sum of
                              weights, i.e. for each node that meets all of the scheduling
                              requirements (resource request, requiredDuringScheduling
                              anti-affinity expressions, etc.), compute a sum by iterating
                              through the elements of this field and adding "weight"
                              to the sum if the node has pods which matches the corresponding
                              podAffinityTerm; the node(s) with the highest sum are
                              the most preferred.
                            items:
                              description: The weights of all of the matched WeightedPodAffinityTerm
                                fields are added per-node to find the most preferred
                                node(s)
                              properties:
                                podAffinityTerm:
                                  description: Required. A pod affinity term, associated
                                    with the corresponding weight.
                                  properties:
                                    labelSelector:
                                      description: A label query over a set of resources,
                                        in this case pods.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaceSelector:
                                      description: A label query over the set of namespaces
                                        that the term applies to. The term is applied
                                        to the union of the namespaces selected by
                                        this field and the ones listed in the namespaces
                                        field. null selector and null or empty namespaces
                                        list means "this pod's namespace". An empty
                                        selector ({}) matches all namespaces.
                                      properties:
                                        matchExpressions:
                                          description: matchExpressions is a list
                                            of label selector requirements. The requirements
                                            are ANDed.
                                          items:
                                            description: A label selector requirement
                                              is a selector that contains values,
                                              a key, and an operator that relates
                                              the key and values.
                                            properties:
                                              key:
                                                description: key is the label key
                                                  that the selector applies to.
                                                type: string
                                              operator:
                                                description: operator represents a
                                                  key's relationship to a set of values.
                                                  Valid operators are In, NotIn, Exists
                                                  and DoesNotExist.
                                                type: string
                                              values:
                                                description: values is an array of
                                                  string values. If the operator is
                                                  In or NotIn, the values array must
                                                  be non-empty. If the operator is
                                                  Exists or DoesNotExist, the values
                                                  array must be empty. This array
                                                  is replaced during a strategic merge
                                                  patch.
                                                items:
                                                  type: string
                                                type: array
                                            required:
                                            - key
                                            - operator
                                            type: object
                                          type: array
                                        matchLabels:
                                          additionalProperties:
                                            type: string
                                          description: matchLabels is a map of {key,value}
                                            pairs. A single {key,value} in the matchLabels
                                            map is equivalent to an element of matchExpressions,
                                            whose key field is "key", the operator
                                            is "In", and the values array contains
                                            only "value". The requirements are ANDed.
                                          type: object
                                      type: object
                                      x-kubernetes-map-type: atomic
                                    namespaces:
                                      description: namespaces specifies a static list
                                        of namespace names that the term applies to.
                                        The term is applied to the union of the namespaces
                                        listed in this field and the ones selected
                                        by namespaceSelector. null or empty namespaces
                                        list and null namespaceSelector means "this
                                        pod's namespace".
                                      items:
                                        type: string
                                      type: array
                                    topologyKey:
                                      description: This pod should be co-located (affinity)
                                        or not co-located (anti-affinity) with the
                                        pods matching the labelSelector in the specified
                                        namespaces, where co-located is defined as
                                        running on a node whose value of the label
                                        with key topologyKey matches that of any node
                                        on which any of the selected pods is running.
                                        Empty topologyKey is not allowed.
                                      type: string
                                  required:
                                  - topologyKey
                                  type: object
                                weight:
                                  description: weight associated with matching the
                                    corresponding podAffinityTerm, in the range 1-100.
                                  format: int32
                                  type: integer
                              required:
                              - podAffinityTerm
                              - weight
                              type: object
                            type: array
                          requiredDuringSchedulingIgnoredDuringExecution:
                            description: If the anti-affinity requirements specified
                              by this field are not met at scheduling time, the pod
                              will not be scheduled onto the node. If the anti-affinity
                              requirements specified by this field cease to be met
                              at some point during pod execution (e.g. due to a pod
                              label update), the system may or may not try to eventually
                              evict the pod from its node. When there are multiple
                              elements, the lists of nodes corresponding to each podAffinityTerm
                              are intersected, i.e. all terms must be satisfied.
                            items:
                              description: Defines a set of pods (namely those matching
                                the labelSelector relative to the given namespace(s))
                                that this pod should be co-located (affinity) or not
                                co-located (anti-affinity) with, where co-located
                                is defined as running on a node whose value of the
                                label with key <topologyKey> matches that of any node
                                on which a pod of the set of pods is running
                              properties:
                                labelSelector:
                                  description: A label query over a set of resources,
                                    in this case pods.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaceSelector:
                                  description: A label query over the set of namespaces
                                    that the term applies to. The term is applied
                                    to the union of the namespaces selected by this
                                    field and the ones listed in the namespaces field.
                                    null selector and null or empty namespaces list
                                    means "this pod's namespace". An empty selector
                                    ({}) matches all namespaces.
                                  properties:
                                    matchExpressions:
                                      description: matchExpressions is a list of label
                                        selector requirements. The requirements are
                                        ANDed.
                                      items:
                                        description: A label selector requirement
                                          is a selector that contains values, a key,
                                          and an operator that relates the key and
                                          values.
                                        properties:
                                          key:
                                            description: key is the label key that
                                              the selector applies to.
                                            type: string
                                          operator:
                                            description: operator represents a key's
                                              relationship to a set of values. Valid
                                              operators are In, NotIn, Exists and
                                              DoesNotExist.
                                            type: string
                                          values:
                                            description: values is an array of string
                                              values. If the operator is In or NotIn,
                                              the values array must be non-empty.
                                              If the operator is Exists or DoesNotExist,
                                              the values array must be empty. This
                                              array is replaced during a strategic
                                              merge patch.
                                            items:
                                              type: string
                                            type: array
                                        required:
                                        - key
                                        - operator
                                        type: object
                                      type: array
                                    matchLabels:
                                      additionalProperties:
                                        type: string
                                      description: matchLabels is a map of {key,value}
                                        pairs. A single {key,value} in the matchLabels
                                        map is equivalent to an element of matchExpressions,
                                        whose key field is "key", the operator is
                                        "In", and the values array contains only "value".
                                        The requirements are ANDed.
                                      type: object
                                  type: object
                                  x-kubernetes-map-type: atomic
                                namespaces:
                                  description: namespaces specifies a static list
                                    of namespace names that the term applies to. The
                                    term is applied to the union of the namespaces
                                    listed in this field and the ones selected by
                                    namespaceSelector. null or empty namespaces list
                                    and null namespaceSelector means "this pod's namespace".
                                  items:
                                    type: string
                                  type: array
                                topologyKey:
                                  description: This pod should be co-located (affinity)
                                    or not co-located (anti-affinity) with the pods
                                    matching the labelSelector in the specified namespaces,
                                    where co-located is defined as running on a node
                                    whose value of the label with key topologyKey
                                    matches that of any node on which any of the selected
                                    pods is running. Empty topologyKey is not allowed.
                                  type: string
                              required:
                              - topologyKey
                              type: object
                            type: array
                        type: object
                    type: object
                  nodeSelector:
                    additionalProperties:
                      type: string
                    description: 'nodeSelector is the node selector applied to the
                      relevant kind of pods It specifies a map of key-value pairs:
                      for the pod to be eligible to run on a node, the node must have
                      each of the indicated key-value pairs as labels (it can have
                      additional labels as well). See https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector'
                    type: object
                  tolerations:
                    description: tolerations is a list of tolerations applied to the
                      relevant kind of pods See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
                      for more info. These are additional tolerations other than default
                      ones.
                    items:
                      description: The pod this Toleration is attached to tolerates
                        any taint that matches the triple <key,value,effect> using
                        the matching operator <operator>.
                      properties:
                        effect:
                          description: Effect indicates the taint effect to match.
                            Empty means match all taint effects. When specified, allowed
                            values are NoSchedule, PreferNoSchedule and NoExecute.
                          type: string
                        key:
                          description: Key is the taint key that the toleration applies
                            to. Empty means match all taint keys. If the key is empty,
                            operator must be Exists; this combination means to match
                            all values and all keys.
                          type: string
                        operator:
                          description: Operator represents a key's relationship to
                            the value. Valid operators are Exists and Equal. Defaults
                            to Equal. Exists is equivalent to wildcard for value,
                            so that a pod can tolerate all taints of a particular
                            category.
                          type: string
                        tolerationSeconds:
                          description: TolerationSeconds represents the period of
                            time the toleration (which must be of effect NoExecute,
                            otherwise this field is ignored) tolerates the taint.
                            By default, it is not set, which means tolerate the taint
                            forever (do not evict). Zero and negative values will
                            be treated as 0 (evict immediately) by the system.
                          format: int64
                          type: integer
                        value:
                          description: Value is the taint value the toleration matches
                            to. If the operator is Exists, the value should be empty,
                            otherwise just a regular string.
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          status:
            description: HostPathProvisionerStatus defines the observed state of HostPathProvisioner
            properties:
              conditions:
                description: Conditions contains the current conditions observed by
                  the operator
                items:
                  description: Condition represents the state of the operator's reconciliation
                    functionality.
                  properties:
                    lastHeartbeatTime:
                      format: date-time
                      type: string
                    lastTransitionTime:
                      format: date-time
                      type: string
                    message:
                      type: string
                    reason:
                      type: string
                    status:
                      type: string
                    type:
                      description: ConditionType is the state of the operator's reconciliation
                        functionality.
                      type: string
                  required:
                  - status
                  - type
                  type: object
                type: array
                x-kubernetes-list-type: atomic
              observedVersion:
                description: ObservedVersion The observed version of the HostPathProvisioner
                  deployment
                type: string
              operatorVersion:
                description: OperatorVersion The version of the HostPathProvisioner
                  Operator
                type: string
              storagePoolStatuses:
                items:
                  description: StoragePoolStatus is the status of the named storage
                    pool
                  properties:
                    claimStatuses:
                      description: The status of all the claims.
                      items:
                        description: ClaimStatus defines the storage claim status
                          for each PVC in a storage pool
                        properties:
                          name:
                            description: Name of the PersistentVolumeClaim
                            type: string
                          status:
                            description: Status of the PersistentVolumeClaim
                            properties:
                              accessModes:
                                description: 'accessModes contains the actual access
                                  modes the volume backing the PVC has. More info:
                                  https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1'
                                items:
                                  type: string
                                type: array
                              allocatedResources:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: allocatedResources is the storage resource
                                  within AllocatedResources tracks the capacity allocated
                                  to a PVC. It may be larger than the actual capacity
                                  when a volume expansion operation is requested.
                                  For storage quota, the larger value from allocatedResources
                                  and PVC.spec.resources is used. If allocatedResources
                                  is not set, PVC.spec.resources alone is used for
                                  quota calculation. If a volume expansion capacity
                                  request is lowered, allocatedResources is only lowered
                                  if there are no expansion operations in progress
                                  and if the actual volume capacity is equal or lower
                                  than the requested capacity. This is an alpha field
                                  and requires enabling RecoverVolumeExpansionFailure
                                  feature.
                                type: object
                              capacity:
                                additionalProperties:
                                  anyOf:
                                  - type: integer
                                  - type: string
                                  pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                                  x-kubernetes-int-or-string: true
                                description: capacity represents the actual resources
                                  of the underlying volume.
                                type: object
                              conditions:
                                description: conditions is the current Condition of
                                  persistent volume claim. If underlying persistent
                                  volume is being resized then the Condition will
                                  be set to 'ResizeStarted'.
                                items:
                                  description: PersistentVolumeClaimCondition contails
                                    details about state of pvc
                                  properties:
                                    lastProbeTime:
                                      description: lastProbeTime is the time we probed
                                        the condition.
                                      format: date-time
                                      type: string
                                    lastTransitionTime:
                                      description: lastTransitionTime is the time
                                        the condition transitioned from one status
                                        to another.
                                      format: date-time
                                      type: string
                                    message:
                                      description: message is the human-readable message
                                        indicating details about last transition.
                                      type: string
                                    reason:
                                      description: reason is a unique, this should
                                        be a short, machine understandable string
                                        that gives the reason for condition's last
                                        transition. If it reports "ResizeStarted"
                                        that means the underlying persistent volume
                                        is being resized.
                                      type: string
                                    status:
                                      type: string
                                    type:
                                      description: PersistentVolumeClaimConditionType
                                        is a valid value of PersistentVolumeClaimCondition.Type
                                      type: string
                                  required:
                                  - status
                                  - type
                                  type: object
                                type: array
                              phase:
                                description: phase represents the current phase of
                                  PersistentVolumeClaim.
                                type: string
                              resizeStatus:
                                description: resizeStatus stores status of resize
                                  operation. ResizeStatus is not set by default but
                                  when expansion is complete resizeStatus is set to
                                  empty string by resize controller or kubelet. This
                                  is an alpha field and requires enabling RecoverVolumeExpansionFailure
                                  feature.
                                type: string
                            type: object
                        required:
                        - name
                        - status
                        type: object
                      type: array
                      x-kubernetes-list-type: atomic
                    currentReady:
                      description: CurrentReady is the number of currently ready replicasets.
                      type: integer
                    desiredReady:
                      description: DesiredReady is the number of desired ready replicasets.
                      type: integer
                    name:
                      description: Name is the name of the storage pool
                      type: string
                    phase:
                      description: StoragePoolPhase indicates which phase the storage
                        pool is in.
                      type: string
                  required:
                  - name
                  - phase
                  type: object
                type: array
                x-kubernetes-list-type: atomic
              targetVersion:
                description: TargetVersion The targeted version of the HostPathProvisioner
                  deployment
                type: string
            type: object
        type: object
    served: true
    storage: true
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: null
  storedVersions: null
`
