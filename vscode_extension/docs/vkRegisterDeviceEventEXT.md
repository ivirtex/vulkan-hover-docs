# vkRegisterDeviceEventEXT(3) Manual Page

## Name

vkRegisterDeviceEventEXT - Signal a fence when a device event occurs



## [](#_c_specification)C Specification

To create a fence that will be signaled when an event occurs on a device, call:

```c++
// Provided by VK_EXT_display_control
VkResult vkRegisterDeviceEventEXT(
    VkDevice                                    device,
    const VkDeviceEventInfoEXT*                 pDeviceEventInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFence*                                    pFence);
```

## [](#_parameters)Parameters

- `device` is a logical device on which the event **may** occur.
- `pDeviceEventInfo` is a pointer to a [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceEventInfoEXT.html) structure describing the event of interest to the application.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pFence` is a pointer to a handle in which the resulting fence object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkRegisterDeviceEventEXT-device-parameter)VUID-vkRegisterDeviceEventEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkRegisterDeviceEventEXT-pDeviceEventInfo-parameter)VUID-vkRegisterDeviceEventEXT-pDeviceEventInfo-parameter  
  `pDeviceEventInfo` **must** be a valid pointer to a valid [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceEventInfoEXT.html) structure
- [](#VUID-vkRegisterDeviceEventEXT-pAllocator-parameter)VUID-vkRegisterDeviceEventEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkRegisterDeviceEventEXT-pFence-parameter)VUID-vkRegisterDeviceEventEXT-pFence-parameter  
  `pFence` **must** be a valid pointer to a [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceEventInfoEXT.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkRegisterDeviceEventEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0