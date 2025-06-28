# VkDedicatedAllocationBufferCreateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationBufferCreateInfoNV - Specify that a buffer is bound to a dedicated memory resource



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkDedicatedAllocationBufferCreateInfoNV` structure, then that structure includes an enable controlling whether the buffer will have a dedicated memory allocation bound to it.

The `VkDedicatedAllocationBufferCreateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationBufferCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           dedicatedAllocation;
} VkDedicatedAllocationBufferCreateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `dedicatedAllocation` specifies whether the buffer will have a dedicated allocation bound to it.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkDedicatedAllocationBufferCreateInfoNV-sType-sType)VUID-VkDedicatedAllocationBufferCreateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV`

## [](#_see_also)See Also

[VK\_NV\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_dedicated_allocation.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDedicatedAllocationBufferCreateInfoNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0