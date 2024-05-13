# VkImageSwapchainCreateInfoKHR(3) Manual Page

## Name

VkImageSwapchainCreateInfoKHR - Specify that an image will be bound to
swapchain memory



## <a href="#_c_specification" class="anchor"></a>C Specification

If the `pNext` chain of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)
includes a `VkImageSwapchainCreateInfoKHR` structure, then that
structure includes a swapchain handle indicating that the image will be
bound to memory from that swapchain.

The `VkImageSwapchainCreateInfoKHR` structure is defined as:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
typedef struct VkImageSwapchainCreateInfoKHR {
    VkStructureType    sType;
    const void*        pNext;
    VkSwapchainKHR     swapchain;
} VkImageSwapchainCreateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `swapchain` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) or a handle of a
  swapchain that the image will be bound to.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-VkImageSwapchainCreateInfoKHR-swapchain-00995"
  id="VUID-VkImageSwapchainCreateInfoKHR-swapchain-00995"></a>
  VUID-VkImageSwapchainCreateInfoKHR-swapchain-00995  
  If `swapchain` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the
  fields of [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html) **must** match
  the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#swapchain-wsi-image-create-info"
  target="_blank" rel="noopener">implied image creation parameters</a>
  of the swapchain

Valid Usage (Implicit)

- <a href="#VUID-VkImageSwapchainCreateInfoKHR-sType-sType"
  id="VUID-VkImageSwapchainCreateInfoKHR-sType-sType"></a>
  VUID-VkImageSwapchainCreateInfoKHR-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_IMAGE_SWAPCHAIN_CREATE_INFO_KHR`

- <a href="#VUID-VkImageSwapchainCreateInfoKHR-swapchain-parameter"
  id="VUID-VkImageSwapchainCreateInfoKHR-swapchain-parameter"></a>
  VUID-VkImageSwapchainCreateInfoKHR-swapchain-parameter  
  If `swapchain` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkImageSwapchainCreateInfoKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
