# vkGetDisplayPlaneCapabilities2KHR(3) Manual Page

## Name

vkGetDisplayPlaneCapabilities2KHR - Query capabilities of a mode and
plane combination



## <a href="#_c_specification" class="anchor"></a>C Specification

To query the capabilities of a given mode and plane combination, call:

``` c
// Provided by VK_KHR_get_display_properties2
VkResult vkGetDisplayPlaneCapabilities2KHR(
    VkPhysicalDevice                            physicalDevice,
    const VkDisplayPlaneInfo2KHR*               pDisplayPlaneInfo,
    VkDisplayPlaneCapabilities2KHR*             pCapabilities);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device associated with
  `pDisplayPlaneInfo`.

- `pDisplayPlaneInfo` is a pointer to a
  [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneInfo2KHR.html) structure
  describing the plane and mode.

- `pCapabilities` is a pointer to a
  [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html)
  structure in which the capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetDisplayPlaneCapabilities2KHR` behaves similarly to
[vkGetDisplayPlaneCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDisplayPlaneCapabilitiesKHR.html),
with the ability to specify extended inputs via chained input
structures, and to return extended information via chained output
structures.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetDisplayPlaneCapabilities2KHR-physicalDevice-parameter"
  id="VUID-vkGetDisplayPlaneCapabilities2KHR-physicalDevice-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilities2KHR-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetDisplayPlaneCapabilities2KHR-pDisplayPlaneInfo-parameter"
  id="VUID-vkGetDisplayPlaneCapabilities2KHR-pDisplayPlaneInfo-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilities2KHR-pDisplayPlaneInfo-parameter  
  `pDisplayPlaneInfo` **must** be a valid pointer to a valid
  [VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneInfo2KHR.html) structure

- <a
  href="#VUID-vkGetDisplayPlaneCapabilities2KHR-pCapabilities-parameter"
  id="VUID-vkGetDisplayPlaneCapabilities2KHR-pCapabilities-parameter"></a>
  VUID-vkGetDisplayPlaneCapabilities2KHR-pCapabilities-parameter  
  `pCapabilities` **must** be a valid pointer to a
  [VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_get_display_properties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_get_display_properties2.html),
[VkDisplayPlaneCapabilities2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneCapabilities2KHR.html),
[VkDisplayPlaneInfo2KHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPlaneInfo2KHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDisplayPlaneCapabilities2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
