# VkDedicatedAllocationMemoryAllocateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationMemoryAllocateInfoNV - Specify a dedicated memory allocation resource



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkDedicatedAllocationMemoryAllocateInfoNV` structure, then that structure includes a handle of the sole buffer or image resource that the memory **can** be bound to.

The `VkDedicatedAllocationMemoryAllocateInfoNV` structure is defined as:

```c++
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationMemoryAllocateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    VkBuffer           buffer;
} VkDedicatedAllocationMemoryAllocateInfoNV;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a handle of an image which this memory will be bound to.
- `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a handle of a buffer which this memory will be bound to.

## [](#_description)Description

Valid Usage

- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00649)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00649  
  At least one of `image` and `buffer` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00650)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00650  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the image **must** have been created with [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation` equal to `VK_TRUE`
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00651)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00651  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), the buffer **must** have been created with [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation` equal to `VK_TRUE`
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00652)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00652  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `VkMemoryAllocateInfo`::`allocationSize` **must** equal the `VkMemoryRequirements`::`size` of the image
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00653)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00653  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `VkMemoryAllocateInfo`::`allocationSize` **must** equal the `VkMemoryRequirements`::`size` of the buffer
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00654)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00654  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation, the memory being imported **must** also be a dedicated image allocation and `image` **must** be identical to the image associated with the imported memory
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00655)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00655  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation, the memory being imported **must** also be a dedicated buffer allocation and `buffer` **must** be identical to the buffer associated with the imported memory

Valid Usage (Implicit)

- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-sType-sType)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV`
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-parameter)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-parameter  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-parameter)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-parameter  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-commonparent)VUID-VkDedicatedAllocationMemoryAllocateInfoNV-commonparent  
  Both of `buffer`, and `image` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_NV\_dedicated\_allocation](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_dedicated_allocation.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDedicatedAllocationMemoryAllocateInfoNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0