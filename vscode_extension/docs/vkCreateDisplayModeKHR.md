# vkCreateDisplayModeKHR(3) Manual Page

## Name

vkCreateDisplayModeKHR - Create a display mode



## [](#_c_specification)C Specification

Additional modes **may** also be created by calling:

```c++
// Provided by VK_KHR_display
VkResult vkCreateDisplayModeKHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display,
    const VkDisplayModeCreateInfoKHR*           pCreateInfo,
    const VkAllocationCallbacks*                pAllocator,
    VkDisplayModeKHR*                           pMode);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device associated with `display`.
- `display` is the display to create an additional mode for.
- `pCreateInfo` is a pointer to a [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateInfoKHR.html) structure describing the new mode to create.
- `pAllocator` is the allocator used for host memory allocated for the display mode object when there is no more specific allocator available (see [Memory Allocation](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#memory-allocation)).
- `pMode` is a pointer to a [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html) handle in which the mode created is returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkCreateDisplayModeKHR-physicalDevice-parameter)VUID-vkCreateDisplayModeKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkCreateDisplayModeKHR-display-parameter)VUID-vkCreateDisplayModeKHR-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html) handle
- [](#VUID-vkCreateDisplayModeKHR-pCreateInfo-parameter)VUID-vkCreateDisplayModeKHR-pCreateInfo-parameter  
  `pCreateInfo` **must** be a valid pointer to a valid [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateInfoKHR.html) structure
- [](#VUID-vkCreateDisplayModeKHR-pAllocator-parameter)VUID-vkCreateDisplayModeKHR-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html) structure
- [](#VUID-vkCreateDisplayModeKHR-pMode-parameter)VUID-vkCreateDisplayModeKHR-pMode-parameter  
  `pMode` **must** be a valid pointer to a [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html) handle
- [](#VUID-vkCreateDisplayModeKHR-display-parent)VUID-vkCreateDisplayModeKHR-display-parent  
  `display` **must** have been created, allocated, or retrieved from `physicalDevice`

Host Synchronization

- Host access to `display` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_INITIALIZATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAllocationCallbacks.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayKHR.html), [VkDisplayModeCreateInfoKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeCreateInfoKHR.html), [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCreateDisplayModeKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0