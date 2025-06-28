# VkCopyMemoryToAccelerationStructureInfoKHR(3) Manual Page

## Name

VkCopyMemoryToAccelerationStructureInfoKHR - Parameters for deserializing an acceleration structure



## [](#_c_specification)C Specification

The `VkCopyMemoryToAccelerationStructureInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyMemoryToAccelerationStructureInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkDeviceOrHostAddressConstKHR         src;
    VkAccelerationStructureKHR            dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyMemoryToAccelerationStructureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the device or host address to memory containing the source data for the copy.
- `dst` is the target acceleration structure for the copy.
- `mode` is a [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-src-04960)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-src-04960  
  The source memory pointed to by `src` **must** contain data previously serialized using [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html), potentially modified to relocate acceleration structure references as described in that command
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-03413)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-03413  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR`
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pInfo-03414)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pInfo-03414  
  The data in `src` **must** have a format compatible with the destination physical device as returned by [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-03746)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-03746  
  `dst` **must** have been created with a `size` greater than or equal to that used to serialize the data in `src`

Valid Usage (Implicit)

- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-sType-sType)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_MEMORY_TO_ACCELERATION_STRUCTURE_INFO_KHR`
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pNext-pNext)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-parameter)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-dst-parameter  
  `dst` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-parameter)VUID-VkCopyMemoryToAccelerationStructureInfoKHR-mode-parameter  
  `mode` **must** be a valid [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html), [VkDeviceOrHostAddressConstKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressConstKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyMemoryToAccelerationStructureKHR.html), [vkCopyMemoryToAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyMemoryToAccelerationStructureKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMemoryToAccelerationStructureInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0