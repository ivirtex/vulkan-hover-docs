# VkMemoryDedicatedAllocateInfo(3) Manual Page

## Name

VkMemoryDedicatedAllocateInfo - Specify a dedicated memory allocation resource



## [](#_c_specification)C Specification

If the `pNext` chain includes a `VkMemoryDedicatedAllocateInfo` structure, then that structure includes a handle of the sole buffer or image resource that the memory **can** be bound to.

The `VkMemoryDedicatedAllocateInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_1
typedef struct VkMemoryDedicatedAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkImage            image;
    VkBuffer           buffer;
} VkMemoryDedicatedAllocateInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_dedicated_allocation
typedef VkMemoryDedicatedAllocateInfo VkMemoryDedicatedAllocateInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `image` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a handle of an image which this memory will be bound to.
- `buffer` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a handle of a buffer which this memory will be bound to.

## [](#_description)Description

Valid Usage

- [](#VUID-VkMemoryDedicatedAllocateInfo-image-01432)VUID-VkMemoryDedicatedAllocateInfo-image-01432  
  At least one of `image` and `buffer` **must** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-02964)VUID-VkMemoryDedicatedAllocateInfo-image-02964  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the memory is not an imported Android Hardware Buffer or an imported QNX Screen buffer , `VkMemoryAllocateInfo`::`allocationSize` **must** equal the `VkMemoryRequirements`::`size` of the image
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-01434)VUID-VkMemoryDedicatedAllocateInfo-image-01434  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** have been created without `VK_IMAGE_CREATE_SPARSE_BINDING_BIT` set in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-02965)VUID-VkMemoryDedicatedAllocateInfo-buffer-02965  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and the memory is not an imported Android Hardware Buffer or an imported QNX Screen buffer , `VkMemoryAllocateInfo`::`allocationSize` **must** equal the `VkMemoryRequirements`::`size` of the buffer
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-01436)VUID-VkMemoryDedicatedAllocateInfo-buffer-01436  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** have been created without `VK_BUFFER_CREATE_SPARSE_BINDING_BIT` set in [VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBufferCreateInfo.html)::`flags`
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-01876)VUID-VkMemoryDedicatedAllocateInfo-image-01876  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, and the external handle was created by the Vulkan API, then the memory being imported **must** also be a dedicated image allocation and `image` **must** be identical to the image associated with the imported memory
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-01877)VUID-VkMemoryDedicatedAllocateInfo-buffer-01877  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_WIN32_KMT_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D11_TEXTURE_KMT_BIT`, `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_HEAP_BIT`, or `VK_EXTERNAL_MEMORY_HANDLE_TYPE_D3D12_RESOURCE_BIT`, and the external handle was created by the Vulkan API, then the memory being imported **must** also be a dedicated buffer allocation and `buffer` **must** be identical to the buffer associated with the imported memory
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-01878)VUID-VkMemoryDedicatedAllocateInfo-image-01878  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, the memory being imported **must** also be a dedicated image allocation and `image` **must** be identical to the image associated with the imported memory
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-01879)VUID-VkMemoryDedicatedAllocateInfo-buffer-01879  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_OPAQUE_FD_BIT`, the memory being imported **must** also be a dedicated buffer allocation and `buffer` **must** be identical to the buffer associated with the imported memory
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-01797)VUID-VkMemoryDedicatedAllocateInfo-image-01797  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** not have been created with `VK_IMAGE_CREATE_DISJOINT_BIT` set in [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-04751)VUID-VkMemoryDedicatedAllocateInfo-image-04751  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the memory being imported **must** also be a dedicated image allocation and `image` **must** be identical to the image associated with the imported memory
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-04752)VUID-VkMemoryDedicatedAllocateInfo-buffer-04752  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) and [VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryAllocateInfo.html) defines a memory import operation with handle type `VK_EXTERNAL_MEMORY_HANDLE_TYPE_ZIRCON_VMO_BIT_FUCHSIA`, the memory being imported **must** also be a dedicated buffer allocation and `buffer` **must** be identical to the buffer associated with the imported memory

Valid Usage (Implicit)

- [](#VUID-VkMemoryDedicatedAllocateInfo-sType-sType)VUID-VkMemoryDedicatedAllocateInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_MEMORY_DEDICATED_ALLOCATE_INFO`
- [](#VUID-VkMemoryDedicatedAllocateInfo-image-parameter)VUID-VkMemoryDedicatedAllocateInfo-image-parameter  
  If `image` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `image` **must** be a valid [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) handle
- [](#VUID-VkMemoryDedicatedAllocateInfo-buffer-parameter)VUID-VkMemoryDedicatedAllocateInfo-buffer-parameter  
  If `buffer` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-VkMemoryDedicatedAllocateInfo-commonparent)VUID-VkMemoryDedicatedAllocateInfo-commonparent  
  Both of `buffer`, and `image` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkMemoryDedicatedAllocateInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0