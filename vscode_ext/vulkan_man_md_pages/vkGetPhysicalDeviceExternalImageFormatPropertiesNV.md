# vkGetPhysicalDeviceExternalImageFormatPropertiesNV(3) Manual Page

## Name

vkGetPhysicalDeviceExternalImageFormatPropertiesNV - Determine image
capabilities compatible with external memory handle types



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine the image capabilities compatible with an external memory
handle type, call:

``` c
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

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the image
  capabilities

- `format` is the image format, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`.

- `type` is the image type, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`imageType`.

- `tiling` is the image tiling, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`tiling`.

- `usage` is the intended usage of the image, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`.

- `flags` is a bitmask describing additional parameters of the image,
  corresponding to [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`.

- `externalHandleType` is either one of the bits from
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html),
  or 0.

- `pExternalImageFormatProperties` is a pointer to a
  [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)
  structure in which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

If `externalHandleType` is 0,
`pExternalImageFormatProperties->imageFormatProperties` will return the
same values as a call to
[vkGetPhysicalDeviceImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties.html),
and the other members of `pExternalImageFormatProperties` will all be 0.
Otherwise, they are filled in as described for
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html).

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-07721"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-07721"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-07721  
  `externalHandleType` **must** not have more than one bit set

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-format-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-format-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-type-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-type-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-tiling-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-tiling-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-requiredbitmask"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-requiredbitmask"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-flags-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-flags-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) values

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-externalHandleType-parameter  
  `externalHandleType` **must** be a valid combination of
  [VkExternalMemoryHandleTypeFlagBitsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagBitsNV.html)
  values

- <a
  href="#VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-pExternalImageFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-pExternalImageFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceExternalImageFormatPropertiesNV-pExternalImageFormatProperties-parameter  
  `pExternalImageFormatProperties` **must** be a valid pointer to a
  [VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html)
  structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_FORMAT_NOT_SUPPORTED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_external_memory_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_external_memory_capabilities.html),
[VkExternalImageFormatPropertiesNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalImageFormatPropertiesNV.html),
[VkExternalMemoryHandleTypeFlagsNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalMemoryHandleTypeFlagsNV.html),
[VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceExternalImageFormatPropertiesNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
