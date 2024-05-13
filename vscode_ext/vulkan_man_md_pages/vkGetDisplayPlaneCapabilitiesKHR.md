# vkGetDisplayPlaneCapabilitiesKHR(3) Manual Page

## Name

vkGetDisplayPlaneCapabilitiesKHR - Query capabilities of a mode and
plane combination



## <a href="#_c_specification" class="anchor"></a>C Specification

Applications that wish to present directly to a display **must** select
which layer, or “plane” of the display they wish to target, and a mode
to use with the display. Each display supports at least one plane. The
capabilities of a given mode and plane combination are determined by
calling:

``` c
// Provided by VK_KHR_display
VkResult vkGetDisplayPlaneCapabilitiesKHR(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayModeKHR                            mode,
    uint32_t                                    planeIndex,
    VkDisplayPlaneCapabilitiesKHR*              pCapabilities);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device associated with the display
  specified by `mode`

- `mode` is the display mode the application intends to program when
  using the specified plane. Note this parameter also implicitly
  specifies a display.

- `planeIndex` is the plane which the application intends to use with
  the display, and is less than the number of display planes supported
  by the device.

- `pCapabilities` is a pointer to a
  [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilitiesKHR.html)
  structure in which the capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDisplayPlaneCapabilitiesKHR-physicalDevice-parameter"
  id="VUID-vkGetDisplayPlaneCapabilitiesKHR-physicalDevice-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilitiesKHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parameter"
  id="VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parameter  
  `mode` **must** be a valid [VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html)
  handle

- <a href="#VUID-vkGetDisplayPlaneCapabilitiesKHR-pCapabilities-parameter"
  id="VUID-vkGetDisplayPlaneCapabilitiesKHR-pCapabilities-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilitiesKHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a
  [VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilitiesKHR.html)
  structure

- <a href="#VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parent"
  id="VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parent"></a>
  VUID-vkGetDisplayPlaneCapabilitiesKHR-mode-parent  
  `mode` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Host Synchronization

- Host access to `mode` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_display.html),
[VkDisplayModeKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayModeKHR.html),
[VkDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilitiesKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDisplayPlaneCapabilitiesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
