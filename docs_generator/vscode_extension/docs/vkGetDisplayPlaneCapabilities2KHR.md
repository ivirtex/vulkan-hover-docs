# vkGetDisplayPlaneCapabilities2KHR(3) Manual Page

## Name

vkGetDisplayPlaneCapabilities2KHR - Query capabilities of a mode and plane combination



## [](#_c_specification)C Specification

To query the capabilities of a given mode and plane combination, call:

```c++
// Provided by VK_KHR_get_display_properties2
VkResult vkGetDisplayPlaneCapabilities2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkDisplayPlaneInfo2KHR*               pDisplayPlaneInfo,
    VkDisplayPlaneCapabilities2KHR*             pCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device associated with `pDisplayPlaneInfo`.
- `pDisplayPlaneInfo` is a pointer to a [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneInfo2KHR.html) structure describing the plane and mode.
- `pCapabilities` is a pointer to a [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilities2KHR.html) structure in which the capabilities are returned.

## [](#_description)Description

`vkGetDisplayPlaneCapabilities2KHR` behaves similarly to [vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetDisplayPlaneCapabilitiesKHR.html), with the ability to specify extended inputs via chained input structures, and to return extended information via chained output structures.

Valid Usage (Implicit)

- [](#VUID-vkGetDisplayPlaneCapabilities2KHR-physicalDevice-parameter)VUID-vkGetDisplayPlaneCapabilities2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetDisplayPlaneCapabilities2KHR-pDisplayPlaneInfo-parameter)VUID-vkGetDisplayPlaneCapabilities2KHR-pDisplayPlaneInfo-parameter  
  `pDisplayPlaneInfo` **must** be a valid pointer to a valid [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneInfo2KHR.html) structure
- [](#VUID-vkGetDisplayPlaneCapabilities2KHR-pCapabilities-parameter)VUID-vkGetDisplayPlaneCapabilities2KHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilities2KHR.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## [](#_see_also)See Also

[VK\_KHR\_get\_display\_properties2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_get_display_properties2.html), [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilities2KHR.html), [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneInfo2KHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDisplayPlaneCapabilities2KHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0