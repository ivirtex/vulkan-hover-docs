# VkClusterAccelerationStructureGetTemplateIndicesInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureGetTemplateIndicesInfoNV - Parameters describing addresses of cluster template acceleration structure whose index data is requested



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureGetTemplateIndicesInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureGetTemplateIndicesInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureGetTemplateIndicesInfoNV {
    VkDeviceAddress    clusterTemplateAddress;
} VkClusterAccelerationStructureGetTemplateIndicesInfoNV;
```

## [](#_members)Members

- `clusterTemplateAddress` is the device address of the cluster template acceleration structure whose index data is being fetched.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureGetTemplateIndicesInfoNV-clusterTemplateAddress-10833)VUID-VkClusterAccelerationStructureGetTemplateIndicesInfoNV-clusterTemplateAddress-10833  
  `clusterTemplateAddress` **must** be a [template cluster acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-clas-template)

Valid Usage (Implicit)

- [](#VUID-VkClusterAccelerationStructureGetTemplateIndicesInfoNV-clusterTemplateAddress-parameter)VUID-VkClusterAccelerationStructureGetTemplateIndicesInfoNV-clusterTemplateAddress-parameter  
  `clusterTemplateAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureGetTemplateIndicesInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0