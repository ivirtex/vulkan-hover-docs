# VkAccelerationStructureVersionInfoKHR(3) Manual Page

## Name

VkAccelerationStructureVersionInfoKHR - Acceleration structure version information



## [](#_c_specification)C Specification

The `VkAccelerationStructureVersionInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkAccelerationStructureVersionInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    const uint8_t*     pVersionData;
} VkAccelerationStructureVersionInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `pVersionData` is a pointer to the version header of an acceleration structure as defined in [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html)

## [](#_description)Description

Note

`pVersionData` is a *pointer* to an array of 2×`VK_UUID_SIZE` `uint8_t` values instead of two `VK_UUID_SIZE` arrays as the expected use case for this member is to be pointed at the header of a previously serialized acceleration structure (via [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html) or [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html)) that is loaded in memory. Using arrays would necessitate extra memory copies of the UUIDs.

Valid Usage (Implicit)

- [](#VUID-VkAccelerationStructureVersionInfoKHR-sType-sType)VUID-VkAccelerationStructureVersionInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_ACCELERATION_STRUCTURE_VERSION_INFO_KHR`
- [](#VUID-VkAccelerationStructureVersionInfoKHR-pNext-pNext)VUID-VkAccelerationStructureVersionInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkAccelerationStructureVersionInfoKHR-pVersionData-parameter)VUID-VkAccelerationStructureVersionInfoKHR-pVersionData-parameter  
  `pVersionData` **must** be a valid pointer to an array of 2×VK\_UUID\_SIZE `uint8_t` values

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetDeviceAccelerationStructureCompatibilityKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDeviceAccelerationStructureCompatibilityKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAccelerationStructureVersionInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0