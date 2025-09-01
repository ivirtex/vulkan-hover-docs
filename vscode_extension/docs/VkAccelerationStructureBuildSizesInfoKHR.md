# VkAccelerationStructureBuildSizesInfoKHR(3) Manual Page

## Name

VkAccelerationStructureBuildSizesInfoKHR - Structure specifying build sizes for an acceleration structure



## [](#_c_specification)C Specification

The `VkAccelerationStructureBuildSizesInfoKHR` structure describes the required build sizes for an acceleration structure and scratch buffers and is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureBuildSizesInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       accelerationStructureSize;
    VkDeviceSize       updateScratchSize;
    VkDeviceSize       buildScratchSize;
} VkAccelerationStructureBuildSizesInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `accelerationStructureSize` is the size in bytes required in a [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) for a build or update operation.
- `updateScratchSize` is the size in bytes required in a scratch buffer for an update operation.
- `buildScratchSize` is the size in bytes required in a scratch buffer for a build operation.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureBuildSizesInfoKHR-sType-sType)VUID-VkAccelerationStructureBuildSizesInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_SIZES_INFO_KHR`
- [](#VUID-VkAccelerationStructureBuildSizesInfoKHR-pNext-pNext)VUID-VkAccelerationStructureBuildSizesInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureBuildSizesKHR.html), [vkGetClusterAccelerationStructureBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetClusterAccelerationStructureBuildSizesNV.html), [vkGetPartitionedAccelerationStructuresBuildSizesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPartitionedAccelerationStructuresBuildSizesNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureBuildSizesInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0