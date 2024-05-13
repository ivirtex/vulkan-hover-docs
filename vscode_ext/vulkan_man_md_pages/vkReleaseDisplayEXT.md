# vkReleaseDisplayEXT(3) Manual Page

## Name

vkReleaseDisplayEXT - Release access to an acquired VkDisplayKHR



## <a href="#_c_specification" class="anchor"></a>C Specification

To release a previously acquired display, call:

``` c
// Provided by VK_EXT_direct_mode_display
VkResult vkReleaseDisplayEXT(
    VkPhysicalDevice                            physicalDevice,
    VkDisplayKHR                                display);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` The physical device the display is on.

- `display` The display to release control of.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkReleaseDisplayEXT-physicalDevice-parameter"
  id="VUID-vkReleaseDisplayEXT-physicalDevice-parameter"></a>
  VUID-vkReleaseDisplayEXT-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkReleaseDisplayEXT-display-parameter"
  id="VUID-vkReleaseDisplayEXT-display-parameter"></a>
  VUID-vkReleaseDisplayEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkReleaseDisplayEXT-display-parent"
  id="VUID-vkReleaseDisplayEXT-display-parent"></a>
  VUID-vkReleaseDisplayEXT-display-parent  
  `display` **must** have been created, allocated, or retrieved from
  `physicalDevice`

Return Codes

On success, this command returns

- `VK_SUCCESS`

This command does not return any failure codes

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_direct_mode_display](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_direct_mode_display.html),
[VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkReleaseDisplayEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
