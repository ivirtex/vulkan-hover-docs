# VkCopyAccelerationStructureToMemoryInfoKHR(3) Manual Page

## Name

VkCopyAccelerationStructureToMemoryInfoKHR - Parameters for serializing an acceleration structure



## [](#_c_specification)C Specification

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyAccelerationStructureToMemoryInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkAccelerationStructureKHR            src;
    VkDeviceOrHostAddressKHR              dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyAccelerationStructureToMemoryInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the source acceleration structure for the copy
- `dst` is the device or host address to memory which is the target for the copy
- `mode` is a [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-04959)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-04959  
  The source acceleration structure `src` **must** have been constructed prior to the execution of this command
- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-dst-03561)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-dst-03561  
  The memory pointed to by `dst` **must** be at least as large as the serialization size of `src`, as reported by [vkWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkWriteAccelerationStructuresPropertiesKHR.html) or [vkCmdWriteAccelerationStructuresPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteAccelerationStructuresPropertiesKHR.html) with a query type of `VK_QUERY_TYPE_ACCELERATION_STRUCTURE_SERIALIZATION_SIZE_KHR`
- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-03412)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-03412  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR`

Valid Usage (Implicit)

- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-sType-sType)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_TO_MEMORY_INFO_KHR`
- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-pNext-pNext)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-parameter)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-src-parameter  
  `src` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-parameter)VUID-VkCopyAccelerationStructureToMemoryInfoKHR-mode-parameter  
  `mode` **must** be a valid [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html), [VkDeviceOrHostAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceOrHostAddressKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureToMemoryKHR.html), [vkCopyAccelerationStructureToMemoryKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureToMemoryKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyAccelerationStructureToMemoryInfoKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0