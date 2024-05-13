# vkGetPhysicalDeviceImageFormatProperties(3) Manual Page

## Name

vkGetPhysicalDeviceImageFormatProperties - Lists physical device's image
format capabilities



## <a href="#_c_specification" class="anchor"></a>C Specification

To query additional capabilities specific to image types, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkGetPhysicalDeviceImageFormatProperties(
    VkPhysicalDevice                            physicalDevice,
    VkFormat                                    format,
    VkImageType                                 type,
    VkImageTiling                               tiling,
    VkImageUsageFlags                           usage,
    VkImageCreateFlags                          flags,
    VkImageFormatProperties*                    pImageFormatProperties);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device from which to query the image
  capabilities.

- `format` is a [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value specifying the image
  format, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`format`.

- `type` is a [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value specifying the image
  type, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`imageType`.

- `tiling` is a [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value specifying the
  image tiling, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`tiling`.

- `usage` is a bitmask of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) specifying the
  intended usage of the image, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`usage`.

- `flags` is a bitmask of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) specifying
  additional parameters of the image, corresponding to
  [VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)::`flags`.

- `pImageFormatProperties` is a pointer to a
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html) structure in
  which capabilities are returned.

## <a href="#_description" class="anchor"></a>Description

The `format`, `type`, `tiling`, `usage`, and `flags` parameters
correspond to parameters that would be consumed by
[vkCreateImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateImage.html) (as members of
[VkImageCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateInfo.html)).

If `format` is not a supported image format, or if the combination of
`format`, `type`, `tiling`, `usage`, and `flags` is not supported for
images, then `vkGetPhysicalDeviceImageFormatProperties` returns
`VK_ERROR_FORMAT_NOT_SUPPORTED`.

The limitations on an image format that are reported by
`vkGetPhysicalDeviceImageFormatProperties` have the following property:
if `usage1` and `usage2` of type
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html) are such that the bits set
in `usage1` are a subset of the bits set in `usage2`, and `flags1` and
`flags2` of type [VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html) are such
that the bits set in `flags1` are a subset of the bits set in `flags2`,
then the limitations for `usage1` and `flags1` **must** be no more
strict than the limitations for `usage2` and `flags2`, for all values of
`format`, `type`, and `tiling`.

If [`VK_EXT_host_image_copy`](VK_EXT_host_image_copy.html) is supported,
`usage` includes `VK_IMAGE_USAGE_SAMPLED_BIT`, and `flags` does not
include either of `VK_IMAGE_CREATE_SPARSE_BINDING_BIT`,
`VK_IMAGE_CREATE_SPARSE_RESIDENCY_BIT`, or
`VK_IMAGE_CREATE_SPARSE_ALIASED_BIT`, then the result of calls to
`vkGetPhysicalDeviceImageFormatProperties` with identical parameters
except for the inclusion of `VK_IMAGE_USAGE_HOST_TRANSFER_BIT_EXT` in
`usage` **must** be identical.

Valid Usage

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-02248"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-02248"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-02248  
  `tiling` **must** not be `VK_IMAGE_TILING_DRM_FORMAT_MODIFIER_EXT`.
  (Use
  [vkGetPhysicalDeviceImageFormatProperties2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceImageFormatProperties2.html)
  instead)

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties-format-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-format-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-format-parameter  
  `format` **must** be a valid [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html) value

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties-type-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-type-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-type-parameter  
  `type` **must** be a valid [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html) value

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-tiling-parameter  
  `tiling` **must** be a valid [VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html) value

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties-usage-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-usage-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-usage-parameter  
  `usage` **must** be a valid combination of
  [VkImageUsageFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlagBits.html) values

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties-usage-requiredbitmask"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-usage-requiredbitmask"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-usage-requiredbitmask  
  `usage` **must** not be `0`

- <a href="#VUID-vkGetPhysicalDeviceImageFormatProperties-flags-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-flags-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-flags-parameter  
  `flags` **must** be a valid combination of
  [VkImageCreateFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlagBits.html) values

- <a
  href="#VUID-vkGetPhysicalDeviceImageFormatProperties-pImageFormatProperties-parameter"
  id="VUID-vkGetPhysicalDeviceImageFormatProperties-pImageFormatProperties-parameter"></a>
  VUID-vkGetPhysicalDeviceImageFormatProperties-pImageFormatProperties-parameter  
  `pImageFormatProperties` **must** be a valid pointer to a
  [VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html) structure

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_FORMAT_NOT_SUPPORTED`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkFormat](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFormat.html),
[VkImageCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageCreateFlags.html),
[VkImageFormatProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageFormatProperties.html),
[VkImageTiling](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageTiling.html), [VkImageType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageType.html),
[VkImageUsageFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageUsageFlags.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceImageFormatProperties"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
