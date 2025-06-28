# VkCopyAccelerationStructureInfoKHR(3) Manual Page

## Name

VkCopyAccelerationStructureInfoKHR - Parameters for copying an acceleration structure



## [](#_c_specification)C Specification

The `VkCopyAccelerationStructureInfoKHR` structure is defined as:

```c++
// Provided by VK_KHR_acceleration_structure
typedef struct VkCopyAccelerationStructureInfoKHR {
    VkStructureType                       sType;
    const void*                           pNext;
    VkAccelerationStructureKHR            src;
    VkAccelerationStructureKHR            dst;
    VkCopyAccelerationStructureModeKHR    mode;
} VkCopyAccelerationStructureInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `src` is the source acceleration structure for the copy.
- `dst` is the target acceleration structure for the copy.
- `mode` is a [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Valid Usage

- [](#VUID-VkCopyAccelerationStructureInfoKHR-mode-03410)VUID-VkCopyAccelerationStructureInfoKHR-mode-03410  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR` or `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR`
- [](#VUID-VkCopyAccelerationStructureInfoKHR-src-04963)VUID-VkCopyAccelerationStructureInfoKHR-src-04963  
  The source acceleration structure `src` **must** have been constructed prior to the execution of this command
- [](#VUID-VkCopyAccelerationStructureInfoKHR-src-03411)VUID-VkCopyAccelerationStructureInfoKHR-src-03411  
  If `mode` is `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`, `src` **must** have been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` in the build
- [](#VUID-VkCopyAccelerationStructureInfoKHR-buffer-03718)VUID-VkCopyAccelerationStructureInfoKHR-buffer-03718  
  The `buffer` used to create `src` **must** be bound to device memory
- [](#VUID-VkCopyAccelerationStructureInfoKHR-buffer-03719)VUID-VkCopyAccelerationStructureInfoKHR-buffer-03719  
  The `buffer` used to create `dst` **must** be bound to device memory
- [](#VUID-VkCopyAccelerationStructureInfoKHR-dst-07791)VUID-VkCopyAccelerationStructureInfoKHR-dst-07791  
  The range of memory backing `dst` that is accessed by this command **must** not overlap the memory backing `src` that is accessed by this command

Valid Usage (Implicit)

- [](#VUID-VkCopyAccelerationStructureInfoKHR-sType-sType)VUID-VkCopyAccelerationStructureInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COPY_ACCELERATION_STRUCTURE_INFO_KHR`
- [](#VUID-VkCopyAccelerationStructureInfoKHR-pNext-pNext)VUID-VkCopyAccelerationStructureInfoKHR-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkCopyAccelerationStructureInfoKHR-src-parameter)VUID-VkCopyAccelerationStructureInfoKHR-src-parameter  
  `src` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-VkCopyAccelerationStructureInfoKHR-dst-parameter)VUID-VkCopyAccelerationStructureInfoKHR-dst-parameter  
  `dst` **must** be a valid [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html) handle
- [](#VUID-VkCopyAccelerationStructureInfoKHR-mode-parameter)VUID-VkCopyAccelerationStructureInfoKHR-mode-parameter  
  `mode` **must** be a valid [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value
- [](#VUID-VkCopyAccelerationStructureInfoKHR-commonparent)VUID-VkCopyAccelerationStructureInfoKHR-commonparent  
  Both of `dst`, and `src` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_KHR\_acceleration\_structure](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_acceleration_structure.html), [VkAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureKHR.html), [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkCmdCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdCopyAccelerationStructureKHR.html), [vkCopyAccelerationStructureKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCopyAccelerationStructureKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyAccelerationStructureInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0