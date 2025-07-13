# VkBindImageMemorySwapchainInfoKHR(3) Manual Page

## Name

VkBindImageMemorySwapchainInfoKHR - Structure specifying swapchain image memory to bind to



## [](#_c_specification)C Specification

If the `pNext` chain of [VkBindImageMemoryInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryInfo.html) includes a `VkBindImageMemorySwapchainInfoKHR` structure, then that structure includes a swapchain handle and image index indicating that the image will be bound to memory from that swapchain.

The `VkBindImageMemorySwapchainInfoKHR` structure is defined as:

```c++
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkBindImageMemorySwapchainInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
    uint32_t           imageIndex;
} VkBindImageMemorySwapchainInfoKHR;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `swapchain` is [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) or a swapchain handle.
- `imageIndex` is an image index within `swapchain`.

## [](#_description)Description

If `swapchain` is not `NULL`, the `swapchain` and `imageIndex` are used to determine the memory that the image is bound to, instead of `memory` and `memoryOffset`.

Memory **can** be bound to a swapchain and use the `pDeviceIndices` or `pSplitInstanceBindRegions` members of [VkBindImageMemoryDeviceGroupInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBindImageMemoryDeviceGroupInfo.html).

Valid Usage

- [](#VUID-VkBindImageMemorySwapchainInfoKHR-imageIndex-01644)VUID-VkBindImageMemorySwapchainInfoKHR-imageIndex-01644  
  `imageIndex` **must** be less than the number of images in `swapchain`
- [](#VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-07756)VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-07756  
  If the `swapchain` has been created with `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_KHR`, `imageIndex` **must** be one that has previously been returned by [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImageKHR.html) or [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkAcquireNextImage2KHR.html)

Valid Usage (Implicit)

- [](#VUID-VkBindImageMemorySwapchainInfoKHR-sType-sType)VUID-VkBindImageMemorySwapchainInfoKHR-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BIND_IMAGE_MEMORY_SWAPCHAIN_INFO_KHR`
- [](#VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-parameter)VUID-VkBindImageMemorySwapchainInfoKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

## [](#_see_also)See Also

[VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBindImageMemorySwapchainInfoKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0