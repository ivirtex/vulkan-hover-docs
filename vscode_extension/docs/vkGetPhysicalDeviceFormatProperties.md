# vkGetPhysicalDeviceFormatProperties(3) Manual Page

## Name

vkGetPhysicalDeviceFormatProperties - Lists physical device's format
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query supported format features which are properties of the physical
device, call:

``` c
// Provided by VK_VERSION_1_0
void vkGetPhysicalDeviceFormatProperties(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkFormatProperties*                         pFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the format
  properties.

- `format` is the format whose properties are queried.

- `pFormatProperties` is a pointer to a
  [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html) structure in which
  physical device properties for `format` are returned.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceFormatProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceFormatProperties-format-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties-format-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceFormatProperties-pFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties-pFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties-pFormatProperties-parameter  
  `pFormatProperties` **must** be a valid pointer to a
  [VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
