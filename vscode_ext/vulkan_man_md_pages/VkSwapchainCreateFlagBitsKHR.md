# VkSwapchainCreateFlagBitsKHR(3) Manual Page

## Name

VkSwapchainCreateFlagBitsKHR - Bitmask controlling swapchain creation



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html)::`flags`,
specifying parameters of swapchain creation, are:

``` c
// Provided by VK_KHR_swapchain
typedef enum VkSwapchainCreateFlagBitsKHR {
  // Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
    VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR = 0x00000001,
  // Provided by VK_VERSION_1_1 with VK_KHR_swapchain
    VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR = 0x00000002,
  // Provided by VK_KHR_swapchain_mutable_format
    VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR = 0x00000004,
  // Provided by VK_EXT_swapchain_maintenance1
    VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT = 0x00000008,
} VkSwapchainCreateFlagBitsKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_SWAPCHAIN_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT_KHR` specifies
  that images created from the swapchain (i.e. with the `swapchain`
  member of
  [VkImageSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageSwapchainCreateInfoKHR.html)
  set to this swapchain’s handle) **must** use
  `VK_IMAGE_CREATE_SPLIT_INSTANCE_BIND_REGIONS_BIT`.

- `VK_SWAPCHAIN_CREATE_PROTECTED_BIT_KHR` specifies that images created
  from the swapchain are protected images.

- `VK_SWAPCHAIN_CREATE_MUTABLE_FORMAT_BIT_KHR` specifies that the images
  of the swapchain **can** be used to create a `VkImageView` with a
  different format than what the swapchain was created with. The list of
  allowed image view formats is specified by adding a
  [VkImageFormatListCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatListCreateInfo.html)
  structure to the `pNext` chain of
  [VkSwapchainCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateInfoKHR.html). In
  addition, this flag also specifies that the swapchain **can** be
  created with usage flags that are not supported for the format the
  swapchain is created with but are supported for at least one of the
  allowed image view formats.

- `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` specifies
  that the implementation **may** defer allocation of memory associated
  with each swapchain image until its index is to be returned from
  [vkAcquireNextImageKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImageKHR.html) or
  [vkAcquireNextImage2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkAcquireNextImage2KHR.html) for the first
  time.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VkSwapchainCreateFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainCreateFlagsKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkSwapchainCreateFlagBitsKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
