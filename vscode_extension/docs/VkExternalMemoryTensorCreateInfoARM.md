# VkExternalMemoryTensorCreateInfoARM(3) Manual Page

## Name

VkExternalMemoryTensorCreateInfoARM - Specify that a tensor may be backed by external memory



## [](#_c_specification)C Specification

To define a set of external memory handle types that **may** be used as backing store for a tensor, add a [VkExternalMemoryTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryTensorCreateInfoARM.html) structure to the `pNext` chain of the [VkTensorCreateInfoARM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTensorCreateInfoARM.html) structure.

The `VkExternalMemoryTensorCreateInfoARM` structure is defined as:

```c++
// Provided by VK_ARM_tensors
typedef struct VkExternalMemoryTensorCreateInfoARM {
    VkStructureType                    sType;
    const void*                        pNext;
    VkExternalMemoryHandleTypeFlags    handleTypes;
} VkExternalMemoryTensorCreateInfoARM;
```

## [](#_members)Members

Note

A `VkExternalMemoryTensorCreateInfoARM` structure with a non-zero `handleTypes` field must be included in the creation parameters for a tensor that will be bound to memory that is either exported or imported.

## [](#_description)Description

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `handleTypes` is zero or a bitmask of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) specifying one or more external memory handle types.

Valid Usage (Implicit)

- [](#VUID-VkExternalMemoryTensorCreateInfoARM-sType-sType)VUID-VkExternalMemoryTensorCreateInfoARM-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_EXTERNAL_MEMORY_TENSOR_CREATE_INFO_ARM`
- [](#VUID-VkExternalMemoryTensorCreateInfoARM-handleTypes-parameter)VUID-VkExternalMemoryTensorCreateInfoARM-handleTypes-parameter  
  `handleTypes` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBits.html) values

## [](#_see_also)See Also

[VK\_ARM\_tensors](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_ARM_tensors.html), [VkExternalMemoryHandleTypeFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExternalMemoryTensorCreateInfoARM).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0