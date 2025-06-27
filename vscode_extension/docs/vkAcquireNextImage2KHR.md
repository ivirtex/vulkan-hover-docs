# vkAcquireNextImage2KHR(3) Manual Page

## Name

vkAcquireNextImage2KHR - Retrieve the index of the next available
presentable image



## <a href="#_c_specification" class="anchor"></a>C Specification

To acquire an available presentable image to use, and retrieve the index
of that image, call:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
VkResult vkAcquireNextImage2KHR(
    VkDevice                                    device,
    const VkAcquireNextImageInfoKHR*            pAcquireInfo,
    uint32_t*                                   pImageIndex);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapchain`.

- `pAcquireInfo` is a pointer to a
  [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html) structure
  containing parameters of the acquire.

- `pImageIndex` is a pointer to a `uint32_t` that is set to the index of
  the next image to use.

## <a href="#_description" class="anchor"></a>Description

If the `swapchain` has been created with the
`VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_EXT` flag, the image
whose index is returned in `pImageIndex` will be fully backed by memory
before this call returns to the application.

Valid Usage

- <a href="#VUID-vkAcquireNextImage2KHR-surface-07784"
  id="VUID-vkAcquireNextImage2KHR-surface-07784"></a>
  VUID-vkAcquireNextImage2KHR-surface-07784  
  If <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#swapchain-acquire-forward-progress"
  target="_blank" rel="noopener">forward progress</a> cannot be
  guaranteed for the `surface` used to create `swapchain`, the `timeout`
  member of `pAcquireInfo` **must** not be `UINT64_MAX`

Valid Usage (Implicit)

- <a href="#VUID-vkAcquireNextImage2KHR-device-parameter"
  id="VUID-vkAcquireNextImage2KHR-device-parameter"></a>
  VUID-vkAcquireNextImage2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAcquireNextImage2KHR-pAcquireInfo-parameter"
  id="VUID-vkAcquireNextImage2KHR-pAcquireInfo-parameter"></a>
  VUID-vkAcquireNextImage2KHR-pAcquireInfo-parameter  
  `pAcquireInfo` **must** be a valid pointer to a valid
  [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html) structure

- <a href="#VUID-vkAcquireNextImage2KHR-pImageIndex-parameter"
  id="VUID-vkAcquireNextImage2KHR-pImageIndex-parameter"></a>
  VUID-vkAcquireNextImage2KHR-pImageIndex-parameter  
  `pImageIndex` **must** be a valid pointer to a `uint32_t` value

Return Codes

On success, this command returns  
- `VK_SUCCESS`

- `VK_TIMEOUT`

- `VK_NOT_READY`

- `VK_SUBOPTIMAL_KHR`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_DEVICE_LOST`

- `VK_ERROR_OUT_OF_DATE_KHR`

- `VK_ERROR_SURFACE_LOST_KHR`

- `VK_ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html),
[VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAcquireNextImageInfoKHR.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAcquireNextImage2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
