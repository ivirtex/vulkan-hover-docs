# VkClusterAccelerationStructureInstantiateClusterInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureInstantiateClusterInfoNV - Parameters describing instantiate operation for a template cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureInstantiateClusterInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureInstantiateClusterInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureInstantiateClusterInfoNV {
    uint32_t                    clusterIdOffset;
    uint32_t                    geometryIndexOffset:24;
    uint32_t                    reserved:8;
    VkDeviceAddress             clusterTemplateAddress;
    VkStridedDeviceAddressNV    vertexBuffer;
} VkClusterAccelerationStructureInstantiateClusterInfoNV;
```

## [](#_members)Members

- `clusterIdOffset` is an unsigned offset applied to the `clusterID` value stored in the cluster template.
- `geometryIndexOffset` is a signed offset applied to the geometry index of each triangle.
- `reserved` is reserved for future use.
- `clusterTemplateAddress` is the address of a previously built cluster template.
- `vertexBuffer` is either `0` or a [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html) structure containing the vertex data for the indexed triangles stored in the cluster template.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-vertexBuffer-10507)VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-vertexBuffer-10507  
  `vertexBuffer` **must** not be `0` if the template was built without vertex data
- [](#VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-vertexBuffer-10508)VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-vertexBuffer-10508  
  The format in `vertexBuffer` **must** match the original format specified in [VkClusterAccelerationStructureTriangleClusterInputNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureTriangleClusterInputNV.html)
- [](#VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-reserved-10509)VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-reserved-10509  
  `reserved` **must** be `0`
- [](#VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-geometryIndexOffset-10510)VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-geometryIndexOffset-10510  
  The maximum geometry index after using the value in `geometryIndexOffset` **must** be less than [VkPhysicalDeviceClusterAccelerationStructurePropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceClusterAccelerationStructurePropertiesNV.html)::`maxClusterGeometryIndex`

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-clusterTemplateAddress-parameter)VUID-VkClusterAccelerationStructureInstantiateClusterInfoNV-clusterTemplateAddress-parameter  
  `clusterTemplateAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkStridedDeviceAddressNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStridedDeviceAddressNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureInstantiateClusterInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0