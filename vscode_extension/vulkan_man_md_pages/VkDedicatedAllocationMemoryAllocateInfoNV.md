# VkDedicatedAllocationMemoryAllocateInfoNV(3) Manual Page

## Name

VkDedicatedAllocationMemoryAllocateInfoNV - Specify a dedicated memory
allocation resource



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a
`VkDedicatedAllocationMemoryAllocateInfoNV` structure, then that
structure includes a handle of the sole buffer or image resource that
the memory **can** be bound to.

The `VkDedicatedAllocationMemoryAllocateInfoNV` structure is defined as:

``` c
// Provided by VK_NV_dedicated_allocation
typedef struct VkDedicatedAllocationMemoryAllocateInfoNV {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    VkBuffer           buffer;
} VkDedicatedAllocationMemoryAllocateInfoNV;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `image` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a handle of an
  image which this memory will be bound to.

- `buffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a handle of a
  buffer which this memory will be bound to.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00649"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00649"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00649  
  At least one of `image` and `buffer` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00650"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00650"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00650  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the image
  **must** have been created with
  [VkDedicatedAllocationImageCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationImageCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00651"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00651"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00651  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the buffer
  **must** have been created with
  [VkDedicatedAllocationBufferCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDedicatedAllocationBufferCreateInfoNV.html)::`dedicatedAllocation`
  equal to `VK_TRUE`

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00652"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00652"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00652  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `VkMemoryAllocateInfo`::`allocationSize` **must** equal the
  `VkMemoryRequirements`::`size` of the image

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00653"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00653"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00653  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `VkMemoryAllocateInfo`::`allocationSize` **must** equal the
  `VkMemoryRequirements`::`size` of the buffer

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00654"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00654"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-00654  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation, the memory being imported **must** also be a
  dedicated image allocation and `image` **must** be identical to the
  image associated with the imported memory

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00655"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00655"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-00655  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation, the memory being imported **must** also be a
  dedicated buffer allocation and `buffer` **must** be identical to the
  buffer associated with the imported memory

Valid Usage (Implicit)

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-sType-sType"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-sType-sType"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_DEDICATED_ALLOCATION_MEMORY_ALLOCATE_INFO_NV`

- <a
  href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-parameter"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-parameter"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-image-parameter  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a
  href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-parameter"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-parameter"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-buffer-parameter  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkDedicatedAllocationMemoryAllocateInfoNV-commonparent"
  id="VUID-VkDedicatedAllocationMemoryAllocateInfoNV-commonparent"></a>
  VUID-VkDedicatedAllocationMemoryAllocateInfoNV-commonparent  
  Both of `buffer`, and `image` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_dedicated_allocation](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_dedicated_allocation.html),
[VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDedicatedAllocationMemoryAllocateInfoNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
