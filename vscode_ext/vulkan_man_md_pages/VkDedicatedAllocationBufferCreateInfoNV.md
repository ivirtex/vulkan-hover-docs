# VkDedicatedAllocationBufferCreateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationBufferCreateInfoNV - Specify that a buffer is bound
to a dedicated memory resource



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a
`VkDedicatedAllocationBufferCreateInfoNV` structure, then that structure
includes an enable controlling whether the buffer will have a dedicated
memory allocation bound to it.

The `VkDedicatedAllocationBufferCreateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationBufferCreateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkBool32           dedicatedAllocation;
} VkDedicatedAllocationBufferCreateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `dedicatedAllocation` specifies whether the buffer will have a
  dedicated allocation bound to it.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkDedicatedAllocationBufferCreateInfoNV-sType-sType"
  id="VUID-VkDedicatedAllocationBufferCreateInfoNV-sType-sType"></a>
  VUID-VkDedicatedAllocationBufferCreateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_BUFFER_CREATE_INFO_NV`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_dedicated_allocation.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDedicatedAllocationBufferCreateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
