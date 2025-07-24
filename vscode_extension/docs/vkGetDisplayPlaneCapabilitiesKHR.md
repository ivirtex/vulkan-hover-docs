# vkGetDisplayPlaneCapabilitiesKHR(3) Manual Page

## Name

vkGetDisplayPlaneCapabilitiesKHR - Query capabilities of a mode and plane combination



## [](#_c_specification)C Specification

Applications that wish to present directly to a display **must** select which layer, or “plane” of the display they wish to target, and a mode to use with the display. Each display supports at least one plane. The capabilities of a given mode and plane combination are determined by calling:

```c++
// Provided by VK_KHR_display
VkResult vkGetDisplayPlaneCapabilitiesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayModeKHR                            mode,
    uint32_t                                    planeIndex,
    VkDisplayPlaneCapabilitiesKHR*              pCapabilities);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device associated with the display specified by `mode`
- `mode` is the display mode the application intends to program when using the specified plane. Note this parameter also implicitly specifies a display.
- `planeIndex` is the plane which the application intends to use with the display, and is less than the number of display planes supported by the device.
- `pCapabilities` is a pointer to a [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilitiesKHR.html) structure in which the capabilities are returned.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkGetDisplayPlaneCapabilitiesKHR-physicalDevice-parameter)VUID-vkGetDisplayPlaneCapabilitiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parameter)VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parameter  
  `mode` **must** be a valid [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html) handle
- [](#VUID-vkGetDisplayPlaneCapabilitiesKHR-pCapabilities-parameter)VUID-vkGetDisplayPlaneCapabilitiesKHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilitiesKHR.html) structure
- [](#VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parent)VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parent  
  `mode` **must** have been created, allocated, or retrieved from `physicalDevice`

Host Synchronization

- Host access to `mode` **must** be externally synchronized

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_KHR\_display](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_display.html), [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayModeKHR.html), [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayPlaneCapabilitiesKHR.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDisplayPlaneCapabilitiesKHR)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0