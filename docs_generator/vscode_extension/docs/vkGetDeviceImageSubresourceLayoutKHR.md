# vkGetDeviceImageSubresourceLayout(3) Manual Page

## Name

vkGetDeviceImageSubresourceLayout - Retrieve information about an image subresource without an image object



## [](#_c_specification)C Specification

To query the memory layout of an image subresource, without an image object, call:

```c++
// Provided by VK_VERSION_1_4
void vkGetDeviceImageSubresourceLayout(
    VkDevice                                    device,
    const VkDeviceImageSubresourceInfo*         pInfo,
    VkSubresourceLayout2*                       pLayout);
```

or the equivalent command

```c++
// Provided by VK_KHR_maintenance5
void vkGetDeviceImageSubresourceLayoutKHR(
    VkDevice                                    device,
    const VkDeviceImageSubresourceInfo*         pInfo,
    VkSubresourceLayout2*                       pLayout);
```

## [](#_parameters)Parameters

- `device` is the logical device that owns the image.
- `pInfo` is a pointer to a [VkDeviceImageSubresourceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfo.html) structure containing parameters required for the subresource layout query.
- `pLayout` is a pointer to a [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure in which the layout is returned.

## [](#_description)Description

`vkGetDeviceImageSubresourceLayout` behaves similarly to [vkGetImageSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetImageSubresourceLayout2.html), but uses a [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html) structure to specify the image rather than a [VkImage](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImage.html) object.

Valid Usage (Implicit)

- [](#VUID-vkGetDeviceImageSubresourceLayout-device-parameter)VUID-vkGetDeviceImageSubresourceLayout-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkGetDeviceImageSubresourceLayout-pInfo-parameter)VUID-vkGetDeviceImageSubresourceLayout-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkDeviceImageSubresourceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfo.html) structure
- [](#VUID-vkGetDeviceImageSubresourceLayout-pLayout-parameter)VUID-vkGetDeviceImageSubresourceLayout-pLayout-parameter  
  `pLayout` **must** be a valid pointer to a [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html) structure

## [](#_see_also)See Also

[VK\_KHR\_maintenance5](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance5.html), [VK\_VERSION\_1\_4](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_4.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkDeviceImageSubresourceInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceImageSubresourceInfo.html), [VkSubresourceLayout2](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSubresourceLayout2.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetDeviceImageSubresourceLayout)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0