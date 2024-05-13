# VkBindImageMemorySwapchainInfoKHR(3) Manual Page

## Name

VkBindImageMemorySwapchainInfoKHR - Structure specifying swapchain image
memory to bind to



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of
[VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryInfo.html) includes a
`VkBindImageMemorySwapchainInfoKHR` structure, then that structure
includes a swapchain handle and image index indicating that the image
will be bound to memory from that swapchain.

The `VkBindImageMemorySwapchainInfoKHR` structure is defined as:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkBindImageMemorySwapchainInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint32_t           imageIndex;
} VkBindImageMemorySwapchainInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchain` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a swapchain
  handle.

- `imageIndex` is an image index within `swapchain`.

## <a href="#_description" class="anchor"></a>Description

If `swapchain` is not `NULL`, the `swapchain` and `imageIndex` are used
to determine the memory that the image is bound to, instead of `memory`
and `memoryOffset`.

Memory **can** be bound to a swapchain and use the `pDeviceIndices` or
`pSplitInstanceBindRegions` members of
[VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindImageMemoryDeviceGroupInfo.html).

Valid Usage

- <a href="#VUID-VkBindImageMemorySwapchainInfoKHR-imageIndex-01644"
  id="VUID-VkBindImageMemorySwapchainInfoKHR-imageIndex-01644"></a>
  VUID-VkBindImageMemorySwapchainInfoKHR-imageIndex-01644  
  `imageIndex` **must** be less than the number of images in `swapchain`

- <a href="#VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-07756"
  id="VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-07756"></a>
  VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-07756  
  If the `swapchain` has been created with
  `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT`, `imageIndex`
  **must** be one that has previously been returned by
  [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html) or
  [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html)

Valid Usage (Implicit)

- <a href="#VUID-VkBindImageMemorySwapchainInfoKHR-sType-sType"
  id="VUID-VkBindImageMemorySwapchainInfoKHR-sType-sType"></a>
  VUID-VkBindImageMemorySwapchainInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`

- <a href="#VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-parameter"
  id="VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-parameter"></a>
  VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBindImageMemorySwapchainInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
