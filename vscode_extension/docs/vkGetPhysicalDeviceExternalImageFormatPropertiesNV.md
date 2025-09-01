# vkGetPhysicalDeviceExternalImageFormatPropertiesNV(3) Manual Page

## Name

vkGetPhysicalDeviceExternalImageFormatPropertiesNV - Determine image capabilities compatible with external memory handle types



## [](#_c_specification)C Specification

To determine the image capabilities compatible with an external memory handle type, call:

```c++
// Provided by VK_NV_external_memory_capabilities
VkResult vkGetPhysicalDeviceExternalImageFormatPropertiesNV(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkImageType                                 type,
    VkImageTiling                               tiling,
    VkImageUsageFlags                           usage,
    VkImageCreateFlags                          flags,
    VkExternalMemoryHandleTypeFlagsNV           externalHandleType,
    VkExternalImageFormatPropertiesNV*          pExternalImageFormatProperties);
```

## [](#_parameters)Parameters

- `physicalDevice` is the physical device from which to query the image capabilities
- `format` is the image format, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`format`.
- `type` is the image type, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`imageType`.
- `tiling` is the image tiling, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`tiling`.
- `usage` is the intended usage of the image, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`usage`.
- `flags` is a bitmask describing additional parameters of the image, corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateInfo.html)::`flags`.
- `externalHandleType` is either one of the bits from [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html), or 0.
- `pExternalImageFormatProperties` is a pointer to a [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html) structure in which capabilities are returned.

## [](#_description)Description

If `externalHandleType` is 0, `pExternalImageFormatProperties->imageFormatProperties` will return the same values as a call to [vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetPhysicalDeviceImageFormatProperties.html), and the other members of `pExternalImageFormatProperties` will all be 0. Otherwise, they are filled in as described for [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html).

Valid Usage

- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-07721)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-07721  
  `externalHandleType` **must** not have more than one bit set

Valid Usage (Implicit)

- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-physicalDevice-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html) handle
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-format-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html) value
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-type-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html) value
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-tiling-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html) value
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-parameter  
  `usage` **must** be a valid combination of [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlagBits.html) values
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-requiredbitmask)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-requiredbitmask  
  `usage` **must** not be `0`
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-flags-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-flags-parameter  
  `flags` **must** be a valid combination of [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlagBits.html) values
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-parameter  
  `externalHandleType` **must** be a valid combination of [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html) values
- [](#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-pExternalImageFormatProperties-parameter)VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-pExternalImageFormatProperties-parameter  
  `pExternalImageFormatProperties` **must** be a valid pointer to a [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html) structure

Return Codes

On success, this command returns

- `VK_SUCCESS`

On failure, this command returns

- `VK_ERROR_FORMAT_NOT_SUPPORTED`
- `VK_ERROR_OUT_OF_DEVICE_MEMORY`
- `VK_ERROR_OUT_OF_HOST_MEMORY`
- `VK_ERROR_UNKNOWN`
- `VK_ERROR_VALIDATION_FAILED`

## [](#_see_also)See Also

[VK\_NV\_external\_memory\_capabilities](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_external_memory_capabilities.html), [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalImageFormatPropertiesNV.html), [VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalMemoryHandleTypeFlagsNV.html), [VkFormat](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFormat.html), [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageCreateFlags.html), [VkImageTiling](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageType.html), [VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkImageUsageFlags.html), [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDevice.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkGetPhysicalDeviceExternalImageFormatPropertiesNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0