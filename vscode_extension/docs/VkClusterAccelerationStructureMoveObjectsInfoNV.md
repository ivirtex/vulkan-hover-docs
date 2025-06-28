# VkClusterAccelerationStructureMoveObjectsInfoNV(3) Manual Page

## Name

VkClusterAccelerationStructureMoveObjectsInfoNV - Parameters describing move operation for a cluster acceleration structure



## [](#_c_specification)C Specification

The [VkClusterAccelerationStructureMoveObjectsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkClusterAccelerationStructureMoveObjectsInfoNV.html) structure is defined as:

```c++
// Provided by VK_NV_cluster_acceleration_structure
typedef struct VkClusterAccelerationStructureMoveObjectsInfoNV {
    VkDeviceAddress    srcAccelerationStructure;
} VkClusterAccelerationStructureMoveObjectsInfoNV;
```

## [](#_members)Members

- `srcAccelerationStructure` is the device address of the source cluster acceleration structure that will be moved.

## [](#_description)Description

Valid Usage

- [](#VUID-VkClusterAccelerationStructureMoveObjectsInfoNV-srcAccelerationStructure-10483)VUID-VkClusterAccelerationStructureMoveObjectsInfoNV-srcAccelerationStructure-10483  
  `srcAccelerationStructure` **must** be a type of [cluster acceleration structure](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#cluster-geometry)

## [](#_see_also)See Also

[VK\_NV\_cluster\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cluster_acceleration_structure.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkClusterAccelerationStructureMoveObjectsInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0