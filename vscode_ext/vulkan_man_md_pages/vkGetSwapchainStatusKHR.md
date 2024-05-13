# vkGetSwapchainStatusKHR(3) Manual Page

## Name

vkGetSwapchainStatusKHR - Get a swapchain's status



## <a href="#_c_specification" class="anchor"></a>C Specification

In order to query a swapchain’s status when rendering to a shared
presentable image, call:

``` c
// Provided by VK_KHR_shared_presentable_image
VkResult vkGetSwapchainStatusKHR(
    VkDevice                                    device,
    VkSwapchainKHR                              swapchain);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `swapchain` is the swapchain to query.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkGetSwapchainStatusKHR-device-parameter"
  id="VUID-vkGetSwapchainStatusKHR-device-parameter"></a>
  VUID-vkGetSwapchainStatusKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetSwapchainStatusKHR-swapchain-parameter"
  id="VUID-vkGetSwapchainStatusKHR-swapchain-parameter"></a>
  VUID-vkGetSwapchainStatusKHR-swapchain-parameter  
  `swapchain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkGetSwapchainStatusKHR-swapchain-parent"
  id="VUID-vkGetSwapchainStatusKHR-swapchain-parent"></a>
  VUID-vkGetSwapchainStatusKHR-swapchain-parent  
  `swapchain` **must** have been created, allocated, or retrieved from
  `device`

Host Synchronization

- Host access to `swapchain` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_SUBOPTIMAL_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

- `VK_ERROR_SURFACE_LOST_KHR`

- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_shared_presentable_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_shared_presentable_image.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetSwapchainStatusKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
