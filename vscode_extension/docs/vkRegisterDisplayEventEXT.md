# vkRegisterDisplayEventEXT(3) Manual Page

## Name

vkRegisterDisplayEventEXT - Signal a fence when a display event occurs



## [](#_c_specification)C Specification

To create a fence that will be signaled when an event occurs on a [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) object, call:

```c++
// Provided by VK_EXT_display_control
VkResult vkRegisterDisplayEventEXT(
    VkDevice                                    device,
    VkDisplayKHR                                display,
    const VkDisplayEventInfoEXT*                pDisplayEventInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkFence*                                    pFence);
```

## [](#_parameters)Parameters

- `device` is a logical device associated with `display`
- `display` is the display on which the event **may** occur.
- `pDisplayEventInfo` is a pointer to a [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventInfoEXT.html) structure describing the event of interest to the application.
- `pAllocator` controls host memory allocation as described in the [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation) chapter.
- `pFence` is a pointer to a handle in which the resulting fence object is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkRegisterDisplayEventEXT-device-parameter)VUID-vkRegisterDisplayEventEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkRegisterDisplayEventEXT-display-parameter)VUID-vkRegisterDisplayEventEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkRegisterDisplayEventEXT-pDisplayEventInfo-parameter)VUID-vkRegisterDisplayEventEXT-pDisplayEventInfo-parameter  
  `pDisplayEventInfo` **must** be a valid pointer to a valid [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventInfoEXT.html) structure
- [](#VUID-vkRegisterDisplayEventEXT-pAllocator-parameter)VUID-vkRegisterDisplayEventEXT-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkRegisterDisplayEventEXT-pFence-parameter)VUID-vkRegisterDisplayEventEXT-pFence-parameter  
  `pFence` **must** be a valid pointer to a [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html) handle
- [](#VUID-vkRegisterDisplayEventEXT-commonparent)VUID-vkRegisterDisplayEventEXT-commonparent  
  Both of `device`, and `display` **must** have been created, allocated, or retrieved from the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_EXT\_display\_control](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_display_control.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDisplayEventInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayEventInfoEXT.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkFence](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFence.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkRegisterDisplayEventEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0