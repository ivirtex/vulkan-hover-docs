# VkMemoryDedicatedAllocateInfo(3) Manual Page

## Name

VkMemoryDedicatedAllocateInfo - Specify a dedicated memory allocation
resource



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain includes a `VkMemoryDedicatedAllocateInfo`
structure, then that structure includes a handle of the sole buffer or
image resource that the memory **can** be bound to.

The `VkMemoryDedicatedAllocateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_1
typedef struct VkMemoryDedicatedAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    VkBuffer           buffer;
} VkMemoryDedicatedAllocateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_dedicated_allocation
typedef VkMemoryDedicatedAllocateInfo VkMemoryDedicatedAllocateInfoKHR;
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

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-01432"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-01432"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-01432  
  At least one of `image` and `buffer` **must** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-02964"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-02964"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-02964  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the memory
  is not an imported Android Hardware Buffer or an imported QNX Screen
  buffer , `VkMemoryAllocateInfo`::`allocationSize` **must** equal the
  `VkMemoryRequirements`::`size` of the image

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-01434"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-01434"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-01434  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** have been created without
  `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` set in
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-02965"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-02965"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-02965  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the
  memory is not an imported Android Hardware Buffer or an imported QNX
  Screen buffer , `VkMemoryAllocateInfo`::`allocationSize` **must**
  equal the `VkMemoryRequirements`::`size` of the buffer

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-01436"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-01436"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-01436  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** have been created without
  `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` set in
  [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html)::`flags`

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-01876"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-01876"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-01876  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, and the external
  handle was created by the Vulkan API, then the memory being imported
  **must** also be a dedicated image allocation and `image` **must** be
  identical to the image associated with the imported memory

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-01877"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-01877"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-01877  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`,
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, and the external
  handle was created by the Vulkan API, then the memory being imported
  **must** also be a dedicated buffer allocation and `buffer` **must**
  be identical to the buffer associated with the imported memory

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-01878"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-01878"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-01878  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, the memory being
  imported **must** also be a dedicated image allocation and `image`
  **must** be identical to the image associated with the imported memory

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-01879"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-01879"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-01879  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, the memory being
  imported **must** also be a dedicated buffer allocation and `buffer`
  **must** be identical to the buffer associated with the imported
  memory

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-01797"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-01797"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-01797  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** not have been created with `VK_IMAGE_CREATE_DISJOINT_BIT` set
  in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-04751"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-04751"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-04751  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the memory
  being imported **must** also be a dedicated image allocation and
  `image` **must** be identical to the image associated with the
  imported memory

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-04752"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-04752"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-04752  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and
  [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) defines a memory
  import operation with handle type
  `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the memory
  being imported **must** also be a dedicated buffer allocation and
  `buffer` **must** be identical to the buffer associated with the
  imported memory

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-sType-sType"
  id="VUID-VkMemoryDedicatedAllocateInfo-sType-sType"></a>
  VUID-VkMemoryDedicatedAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO`

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-image-parameter"
  id="VUID-VkMemoryDedicatedAllocateInfo-image-parameter"></a>
  VUID-VkMemoryDedicatedAllocateInfo-image-parameter  
  If `image` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `image`
  **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html) handle

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-buffer-parameter"
  id="VUID-VkMemoryDedicatedAllocateInfo-buffer-parameter"></a>
  VUID-VkMemoryDedicatedAllocateInfo-buffer-parameter  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-VkMemoryDedicatedAllocateInfo-commonparent"
  id="VUID-VkMemoryDedicatedAllocateInfo-commonparent"></a>
  VUID-VkMemoryDedicatedAllocateInfo-commonparent  
  Both of `buffer`, and `image` that are valid handles of non-ignored
  parameters **must** have been created, allocated, or retrieved from
  the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryDedicatedAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
