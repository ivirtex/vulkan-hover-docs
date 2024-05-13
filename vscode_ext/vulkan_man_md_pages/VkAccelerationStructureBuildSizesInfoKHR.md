# VkAccelerationStructureBuildSizesInfoKHR(3) Manual Page

## Name

VkAccelerationStructureBuildSizesInfoKHR - Structure specifying build
sizes for an acceleration structure



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkAccelerationStructureBuildSizesInfoKHR` structure describes the
required build sizes for an acceleration structure and scratch buffers
and is defined as:

``` c
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureBuildSizesInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceSize       accelerationStructureSize;
    VkDeviceSize       updateScratchSize;
    VkDeviceSize       buildScratchSize;
} VkAccelerationStructureBuildSizesInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `accelerationStructureSize` is the size in bytes required in a
  [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAccelerationStructureKHR.html) for a
  build or update operation.

- `updateScratchSize` is the size in bytes required in a scratch buffer
  for an update operation.

- `buildScratchSize` is the size in bytes required in a scratch buffer
  for a build operation.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkAccelerationStructureBuildSizesInfoKHR-sType-sType"
  id="VUID-VkAccelerationStructureBuildSizesInfoKHR-sType-sType"></a>
  VUID-VkAccelerationStructureBuildSizesInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_BUILD_SIZES_INFO_KHR`

- <a href="#VUID-VkAccelerationStructureBuildSizesInfoKHR-pNext-pNext"
  id="VUID-VkAccelerationStructureBuildSizesInfoKHR-pNext-pNext"></a>
  VUID-VkAccelerationStructureBuildSizesInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_acceleration_structure](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_acceleration_structure.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[vkGetAccelerationStructureBuildSizesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetAccelerationStructureBuildSizesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkAccelerationStructureBuildSizesInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
