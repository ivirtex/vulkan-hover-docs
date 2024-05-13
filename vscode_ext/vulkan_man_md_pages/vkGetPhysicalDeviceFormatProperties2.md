# vkGetPhysicalDeviceFormatProperties2(3) Manual Page

## Name

vkGetPhysicalDeviceFormatProperties2 - Lists physical device's format
capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query supported format features which are properties of the physical
device, call:

``` c
// Provided by VK_VERSION_1_1
void vkGetPhysicalDeviceFormatProperties2(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkFormatProperties2*                        pFormatProperties);
```

or the equivalent command

``` c
// Provided by VK_KHR_get_physical_device_properties2
void vkGetPhysicalDeviceFormatProperties2KHR(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkFormatProperties2*                        pFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the format
  properties.

- `format` is the format whose properties are queried.

- `pFormatProperties` is a pointer to a
  [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html) structure in which
  physical device properties for `format` are returned.

## <a href="#_description" class="anchor"></a>Description

`vkGetPhysicalDeviceFormatProperties2` behaves similarly to
[vkGetPhysicalDeviceFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceFormatProperties.html),
with the ability to return extended information in a `pNext` chain of
output structures.

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceFormatProperties2-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties2-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties2-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a href="#VUID-vkGetPhysicalDeviceFormatProperties2-format-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties2-format-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties2-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceFormatProperties2-pFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceFormatProperties2-pFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceFormatProperties2-pFormatProperties-parameter  
  `pFormatProperties` **must** be a valid pointer to a
  [VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html) structure

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormatProperties2.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceFormatProperties2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
