# vkAcquireNextImage2KHR(3) Manual Page

## Name

vkAcquireNextImage2KHR - Retrieve the index of the next available presentable image



## [](#_c_specification)C Specification

To acquire an available presentable image to use, and retrieve the index of that image, call:

```c++
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_swapchain
VkResult vkAcquireNextImage2KHR(
    VkDevice                                    device,
    const VkAcquireNextImageInfoKHR*            pAcquireInfo,
    uint32_t*                                   pImageIndex);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapchain`.
- `pAcquireInfo` is a pointer to a [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html) structure containing parameters of the acquire.
- `pImageIndex` is a pointer to a `uint32_t` value specifying the index of the next image to use.

## [](#_description)Description

If the `swapchain` has been created with the `VK_SWAPCHAIN_CREATE_DEFERRED_MEMORY_ALLOCATION_BIT_KHR` flag, the image whose index is returned in `pImageIndex` will be fully backed by memory before this call returns to the application.

Valid Usage

- [](#VUID-vkAcquireNextImage2KHR-surface-07784)VUID-vkAcquireNextImage2KHR-surface-07784  
  If [forward progress](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#swapchain-acquire-forward-progress) cannot be guaranteed for the `surface` used to create `swapchain`, the `timeout` member of `pAcquireInfo` **must** not be `UINT64_MAX`

Valid Usage (Implicit)

- [](#VUID-vkAcquireNextImage2KHR-device-parameter)VUID-vkAcquireNextImage2KHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkAcquireNextImage2KHR-pAcquireInfo-parameter)VUID-vkAcquireNextImage2KHR-pAcquireInfo-parameter  
  `pAcquireInfo` **must** be a valid pointer to a valid [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html) structure
- [](#VUID-vkAcquireNextImage2KHR-pImageIndex-parameter)VUID-vkAcquireNextImage2KHR-pImageIndex-parameter  
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

## [](#_see_also)See Also

[VK\_KHR\_device\_group](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_device_group.html), [VK\_KHR\_swapchain](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_swapchain.html), [VK\_VERSION\_1\_1](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_1.html), [VkAcquireNextImageInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAcquireNextImageInfoKHR.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkAcquireNextImage2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0